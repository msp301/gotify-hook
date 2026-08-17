package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Message(title string, msg string, priority int) {
	server := os.Getenv("GOTIFY_SERVER")

	form := url.Values{
		"title":    {title},
		"message":  {msg},
		"priority": {fmt.Sprintf("%d", priority)},
	}
	req, err := http.NewRequest("POST", server+"/message", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Failed to setup message request")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Gotify-Key", "")

	resp, err := http.DefaultClient.Do(req)
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
