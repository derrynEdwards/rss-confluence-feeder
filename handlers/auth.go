package handlers

import (
	"net/http"

	"github.com/derrynEdwards/rss-confluence-feeder/internal/helpers"
)

type authedHandler func(http.ResponseWriter, *http.Request)

func (cfg *Config) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		paramValue := queryParams.Get("token")

		if paramValue != "" {
			if paramValue != cfg.Token {
				helpers.RespondWithError(w, http.StatusUnauthorized, "Invalid token!")
				return
			} else {
				handler(w, r)
			}
		} else {
			helpers.RespondWithError(w, http.StatusBadRequest, "Missing 'token' query param!")
			return
		}
	}
}
