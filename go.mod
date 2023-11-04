module github.com/derrynEdwards/rss-confluence-feeder

go 1.21.1

replace github.com/derrynEdwards/rss-confluence-feeder/handlers v0.0.0 => ./handlers
replace github.com/derrynEdwards/rss-confluence-feeder/internal/helpers v0.0.0 => ./internal/helpers
replace github.com/derrynEdwards/rss-confluence-feeder/internal/confluence v0.0.0 => ./internal/confluence

require (
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/go-chi/httplog v0.3.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.29.1 // indirect
	golang.org/x/sys v0.10.0 // indirect
)
