package router

import (
	"fmt"
	"net/http"

	"github.com/Finn-dot-de/Neural-Gateway/src/llmfuncs"
	"github.com/Finn-dot-de/Neural-Gateway/src/middleware"
	"github.com/go-chi/chi/v5"
)

// NewRouter erstellt einen neuen Router mit allen Routen und Middleware
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware einhängen
	r.Use(middleware.LoggerMiddleware)
	r.Use(middleware.NoCacheMiddleware)

	// API-Routen
	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		// 🦙 NEU: Llama Kommunikation
		r.Get("/ask", func(w http.ResponseWriter, r *http.Request) {
			prompt := r.URL.Query().Get("prompt")
			if prompt == "" {
				http.Error(w, "Fehler: prompt-Parameter fehlt", http.StatusBadRequest)
				return
			}

			answer, err := llmfuncs.ContactLLama(prompt)
			if err != nil {
				http.Error(w, fmt.Sprintf("Fehler bei der Anfrage an Llama: %v", err), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`{"answer": "%s"}`, answer)))
		})
	})

	// Statische Files
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./app")))
	r.Handle("/static/*", fs)

	// Root-Route (z.B. für deine index.html)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./app/index.html")
	})

	return r
}
