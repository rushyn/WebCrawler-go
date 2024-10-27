package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)



func getHTML(rawURL string) (string, error) {
	resp, err := http.Get("https://" + rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("http Status Code %d", resp.StatusCode)
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html"){
		return "", errors.New("text/html not found")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}


	return string(b), nil
}