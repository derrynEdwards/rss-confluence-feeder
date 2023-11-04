package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/derrynEdwards/rss-confluence-feeder/internal/confluence"
	"github.com/derrynEdwards/rss-confluence-feeder/internal/helpers"
	"log"
	"net/http"
)

type PostBody struct {
	Representation string
	Value          string
}

func (cfg *Config) HandlerWebhook(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string
	}

	decoder := json.NewDecoder(r.Body)
	params := helpers.Feed{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Print("Unable to decode parameters!")
	} else {
		post := helpers.BlogRequest{}
		item := params.Items[0]

		post.SpaceID = cfg.SpaceID
		post.Title = item.Title
		post.Status = "current"
		post.Body = helpers.BlogBody{
			Representation: "wiki",
			Value:          fmt.Sprintf("[%v]", item.Canonical[0].Href),
		}

		posted, err := confluence.PostBlog(cfg.ConfluenceURL, cfg.User, cfg.ApiToken, post)

		if err != nil {
			log.Print("Error creating blog!")
		}

		log.Print(posted)
	}

	helpers.RespondWithJSON(w, http.StatusOK, response{
		Status: "OK",
	})
}
