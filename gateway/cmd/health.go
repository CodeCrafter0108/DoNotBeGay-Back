package main

import "net/http"

type Health struct {
	Status  string `json:"status"`
	Env     string `json:"env"`
	Version string `json:"version"`
	IP      string `json:"ip"`
}

//	@Summary		Get Health
//	@Description	Get health status of the service
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	Health
//	@Failure		500	{object}	ErrorResponse
//	@Router			/health [get]
func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	data := &Health{
		Status:  "ok",
		Env:     app.config.env,
		Version: version,
		IP:      r.RemoteAddr,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
