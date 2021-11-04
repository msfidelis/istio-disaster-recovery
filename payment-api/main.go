package main

import (
	"github.com/gin-contrib/logger"
	"payment-api/controllers/healthcheck"
	"payment-api/controllers/liveness"
	"payment-api/controllers/readiness"
	"payment-api/controllers/version"
	"payment-api/controllers/payments"
	"payment-api/pkg/memory_cache"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	chaos "github.com/msfidelis/gin-chaos-monkey"

	"github.com/gin-gonic/gin"

	_ "payment-api/docs"

	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// logger
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
	subLog := zerolog.New(os.Stdout).With().Logger()

	// Memory Cache Singleton
	c := memory_cache.GetInstance()

	// Readiness Probe Mock Config - Warmup in Seconds
	probe_time_raw := os.Getenv("READINESS_PROBE_MOCK_TIME_IN_SECONDS")
	if probe_time_raw == "" {
		probe_time_raw = "5" // 5 Seconds after boot to success readiness response ok
	}
	probe_time, err := strconv.ParseUint(probe_time_raw, 10, 64)
	if err != nil {
		fmt.Println("Environment variable READINESS_PROBE_MOCK_TIME_IN_SECONDS conversion error", err)
	}

	// Set time in Memory Cache
	c.Set("readiness.ok", "false", time.Duration(probe_time)*time.Second)

	// New Router
	router := gin.New()

	//Middlewares
	router.Use(gin.Recovery())
	router.Use(chaos.Load())

	router.Use(logger.SetLogger(logger.Config{
		Logger:   &subLog,
		UTC:      true,
		SkipPath: []string{"/skip"},
	}))

	// Healthcheck Router
	router.GET("/healthcheck", healthcheck.Ok)

	// Version Router
	router.GET("/version", version.Get)

	// Liveness and Readiness
	router.GET("/liveness", liveness.Ok)
	router.GET("/readiness", readiness.Ok)

	// Faker 
	router.GET("/payments/:id", payments.Get)

	router.Run()
}
