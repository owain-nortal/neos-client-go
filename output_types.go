package neos

import (
	"time"
)

type OutputState struct {
	State   string `json:"state"`
	Healthy bool   `json:"healthy"`
}

type Output struct {
	Identifier  string      `json:"identifier"`
	Urn         string      `json:"urn"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Label       string      `json:"label"`
	CreatedAt   time.Time   `json:"created_at"`
	Owner       string      `json:"owner"`
	State       OutputState `json:"state"`
	OutputType  string      `json:"output_type"`
}

type OutputList struct {
	Entities []Output `json:"entities"`
}

type OutputPostRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	OutputType  string `json:"output_type"`
}

type OutputPostRequestEntityInfo struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type OutputPostRequest struct {
	Entity     OutputPostRequestEntity     `json:"entity"`
	EntityInfo OutputPostRequestEntityInfo `json:"entity_info"`
}

type OutputPostResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
	OutputType  string    `json:"output_type"`
}

type OutputPutRequest struct {
	Entity OutputPutRequestEntity `json:"entity"`
}

type OutputPutRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	OutputType  string `json:"output_type"`
}

type OutputPutRequestEntityInfo struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type OutputPutResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
	OutputType  string    `json:"output_type"`
}

type OutputPutInfoResponse struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}
