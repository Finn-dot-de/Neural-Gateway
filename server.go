package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Finn-dot-de/Neural-Gateway/src/llmfuncs"
	"github.com/Finn-dot-de/Neural-Gateway/src/middleware"
	"github.com/Finn-dot-de/Neural-Gateway/src/utils"
	"github.com/go-chi/chi"
)

var question string

const DefaultAppPort = "8080"

func main() {
	initialize()

	question = "Hallo, Llama!"

	answer, err := llmfuncs.ContactLLama(question)
	if err != nil {
		fmt.Printf("Fehler: %v\n", err)
		return
	} else {
		fmt.Printf("Antwort: %s\n", answer)
	}
}

// initializeRouter initializes the HTTP router and configures routes, middleware, and protected route groups.
func initializeRouter() *chi.Mux {
	r := chi.NewMux()
	r.Use(middleware.LoggerMiddleware)
	r.Use(middleware.NoCacheMiddleware)

	return r
}

func initialize() {
	utils.LoadEnv()
	initializeLoggerOrExit("app.log")

	r := initializeRouter()
	setupRoutes(r)

	startServer(r)
	serveStaticFiles(r)
}

func setupRoutes(r *chi.Mux) {
	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

}

func serveStaticFiles(r *chi.Mux) {
	fs := http.FileServer(http.Dir("./app"))
	r.Handle("/*", fs)
}

func startServer(r *chi.Mux) {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = DefaultAppPort
	}
	log.Println("Der Server läuft auf Port " + appPort + "!")
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}

func initializeLoggerOrExit(logFile string) {
	err := middleware.InitializeLogger(logFile)
	if err != nil {
		log.Fatalf("Fehler beim Initialisieren des Loggings: %v", err)
	}
}
