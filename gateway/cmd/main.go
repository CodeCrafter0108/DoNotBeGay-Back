package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/eriscrypt/donotbegay/common/cache"
	"github.com/eriscrypt/donotbegay/common/database"
	"github.com/eriscrypt/donotbegay/common/env"
	"github.com/go-redis/redis/v8"
	"github.com/lpernett/godotenv"
	"go.uber.org/zap"
)

const version = "1.0.0"

type application struct {
	config config
	logger *zap.SugaredLogger
	cache  *redis.Client
}

type config struct {
	addr  string
	env   string
	db    dbConfig
	redis redisConfig
}

type redisConfig struct {
	addr     string
	password string
	db       int
	enabled  bool
}

type dbConfig struct {
	uri                string
	maxOpenConnections int
	maxIdleConnections int
	maxIdleTime        string
}

//	@title			DoNotBeGay API
//	@description	This is a Gateway API for DoNotBeGay service
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath	/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	host := env.GetString("HOST", "localhost")
	dbURI := os.Getenv("DB_URI")
	dbURI = strings.Replace(dbURI, "localhost", host, 1)

	cfg := config{
		addr: fmt.Sprintf(":%d", env.GetInt("PORT", 4000)),
		db: dbConfig{
			uri:                dbURI,
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNECTIONS", 30),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNECTIONS", 30),
			maxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		redis: redisConfig{
			addr:     env.GetString("REDIS_ADDR", "localhost:6379"),
			password: env.GetString("REDIS_PASSWORD", ""),
			db:       env.GetInt("REDIS_DB", 0),
			enabled:  env.GetBool("REDIS_ENABLED", false),
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Database
	db, err := database.New(cfg.db.uri, cfg.db.maxOpenConnections, cfg.db.maxIdleConnections, cfg.db.maxIdleTime)
	if err != nil {
		logger.Fatal(err)
	}

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			logger.Fatal(err)
		}
	}(db)
	logger.Info("database connection has been established")

	// Cache
	var rCache *redis.Client
	if cfg.redis.enabled {
		rCache = cache.NewRedisClient(cfg.redis.addr, cfg.redis.password, cfg.redis.db)
		logger.Info("redis cache connection established")
	}

	app := &application{
		config: cfg,
		logger: logger,
		cache:  rCache,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
