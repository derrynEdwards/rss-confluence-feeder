package confluence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/derrynEdwards/rss-confluence-feeder/internal/helpers"
)

func PostBlog(baseURL, user, apiToken string, post helpers.BlogRequest) (string, error) {
	payload, err := json.Marshal(post)
	if err != nil {
		log.Println("Unable to marshal blog post payload!")
		return "Failed to create blog post", err
	}

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Println("Error creating HTTP request", err)
		return "Failed to create blog post", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}

	req.SetBasicAuth(user, apiToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending HTTP request", err)
		return "Failed to create blog post", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("POST request was successful.")
		return "Blog Created Successfully!", nil
	} else {
		log.Printf("POST request failed with status code: %d\n", resp.StatusCode)
		return "Failed to create blog post", errors.New(fmt.Sprintf("%d", resp.StatusCode))
	}
}
