package neos

import (
	"fmt"
	"io"
	"net/http"
)

func boolToString(input bool) string {
	if input {
		return "true"
	}
	return "false"
}

func createHttpRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", GetAccessToken()))
	return req, err
}
