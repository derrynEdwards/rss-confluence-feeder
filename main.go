package main

import (
	"log"
	"net/http"
	"os"

	"github.com/derrynEdwards/rss-confluence-feeder/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	api_token := os.Getenv("CONFLUENCE_API_TOKEN")
	confluence_url := os.Getenv("CONFLUENCE_API_URL")
	auth_token := os.Getenv("AUTH_TOKEN")
	space_id := os.Getenv("CONFLUENCE_SPACE_ID")
	api_user := os.Getenv("CONFLUENCE_API_USER")

	logger := httplog.NewLogger("access-log", httplog.Options{
		JSON:    true,
		Concise: true,
	})

	if port == "" {
		port = "8080"
	}
	if api_token == "" {
		log.Fatal("No API Token Found.")
		os.Exit(1)
	}
	if confluence_url == "" {
		log.Fatal("No Confluence URL Found.")
		os.Exit(1)
	}

	cfg := handlers.Config{
		ApiToken:      api_token,
		ConfluenceURL: confluence_url,
		User:          api_user,
		Token:         auth_token,
		SpaceID:       space_id,
	}

	router := chi.NewRouter()
	router.Use(httplog.RequestLogger(logger))
	corsMux := middleWareLog(middlewareCors(router))

	router.Get("/healthz", handlers.HandlerHealthz)
	router.Post("/webhook", cfg.MiddlewareAuth(cfg.HandlerWebhook))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
