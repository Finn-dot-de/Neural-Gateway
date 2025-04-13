package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Finn-dot-de/Neural-Gateway/src/router"
	"github.com/Finn-dot-de/Neural-Gateway/src/utils"
	"github.com/Finn-dot-de/Neural-Gateway/src/middleware"
)

const DefaultAppPort = "8080"

func main() {
	utils.LoadEnv()
	initializeLoggerOrExit("app.log")

	r := router.NewRouter()

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = DefaultAppPort
	}

	log.Println("🚀 Server läuft auf Port " + appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}

func initializeLoggerOrExit(logFile string) {
	err := middleware.InitializeLogger(logFile)
	if err != nil {
		log.Fatalf("Fehler beim Initialisieren des Loggings: %v", err)
	}
}
