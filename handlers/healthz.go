package handlers

import (
	"net/http"

	"github.com/derrynEdwards/rss-confluence-feeder/internal/helpers"
)

func HandlerHealthz(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}

	helpers.RespondWithJSON(w, http.StatusOK, response{
		Status: "OK",
	})
}
