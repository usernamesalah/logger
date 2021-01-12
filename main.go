package main

import (
	"log"
	"net/http"
	"time"

	"github.com/usernamesalah/logger/internal/logger"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Reading the configuration from environment variables ...")
	cfg, err := ReadConfig()
	if err != nil {
		panic(err)
	}

	logConfig := logger.Configuration{
		EnableConsole:     true,    // next, get from configuration
		ConsoleJSONFormat: true,    // next, get from configuration
		ConsoleLevel:      "debug", // next, get from configuration
	}

	if err := logger.NewLogger(logConfig, logger.InstanceZapLogger); err != nil {
		log.Fatalf("Could not instantiate log %v", err)
	}

	log.Println("Initializing the web server ...")
	e := echo.New()
	e.GET("/ping", ping)
	// Start server
	s := &http.Server{
		Addr:         "0.0.0.0:" + cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	e.Logger.Fatal(e.StartServer(s))
}

// ping write pong to http.ResponseWriter.
func ping(c echo.Context) error {
	logger.WithFields(logger.Fields{"component": "main", "action": "ping"}).
		Infof("testing server connection.")
	return c.String(http.StatusOK, "pong")
}
