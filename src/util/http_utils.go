package util

import "net/http"

// GenerateHTTPClient /** Generates an HTTP client pointer which we can use for requests. **/
func GenerateHTTPClient() *http.Client {
	client := &http.Client{}
	return client
}

/** HTTP mocking **/
