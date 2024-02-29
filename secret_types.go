package neos

import (
)

type SecretState struct {
	State   string `json:"state"`
	Healthy bool   `json:"healthy"`
}

type Secret struct {
	Identifier string   `json:"identifier"`
	Urn        string   `json:"urn"`
	Name       string   `json:"name"`
	IsSystem   bool     `json:"is_system"`
	Keys       []string `json:"keys"`
}

type SecretList struct {
	Secrets []Secret `json:"secrets"`
}

type SecretPostRequest struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

type SecretPostResponse struct {
	Identifier string   `json:"identifier"`
	Urn        string   `json:"urn"`
	Name       string   `json:"name"`
	IsSystem   bool     `json:"is_system"`
	Keys       []string `json:"keys"`
}

type SecretPutRequest struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

type SecretPutResponse struct {
	Identifier string   `json:"identifier"`
	Urn        string   `json:"urn"`
	Name       string   `json:"name"`
	IsSystem   bool     `json:"is_system"`
	Keys       []string `json:"keys"`
}

