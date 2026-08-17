package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func Message(title string, msg string, priority int) {
	server := os.Getenv("GOTIFY_SERVER")

	resp, err := http.PostForm(server+"/message", url.Values{
		"title":    {title},
		"message":  {msg},
		"priority": {fmt.Sprintf("%d", priority)},
	})

	if err != nil {
		fmt.Println("Failed to send message")
	}

	if resp.StatusCode != http.StatusOK {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to read response")
		}

		fmt.Printf("Error: '%s' - '%s'\n", resp.Status, respBody)
	}
}
