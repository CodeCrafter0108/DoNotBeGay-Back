package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eriscrypt/donotbegay/gateway/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(app.recoverPanic)

	// Define routes
	r.Route("/v1", func(r chi.Router) {

		docsUrl := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsUrl)))

		r.Get("/health", app.healthHandler)

		r.Route("/static", func(r chi.Router) {
			r.Put("/tokens/save/{randomBytes}", app.tokenSaveHandler)
			r.Get("/tokens/{chainId}/{tournamentId}/{filename}", app.tokenGetHandler)
			r.Post("/tokens/create/{chainId}/{userAddress}/{randomBytes}/{tournamentId}", app.tokenPreSaveHandler)
		})
	})

	// Serve static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return r
}

func (app *application) run(mux http.Handler) error {
	// Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.BasePath = "/v1"

	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	app.logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)

	return srv.ListenAndServe()
}
