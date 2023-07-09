package neos

import (
	"time"
)

type DataSystemState struct {
	State   string `json:"state"`
	Healthy bool   `json:"healthy"`
}

type DataSystem struct {
	Identifier  string          `json:"identifier"`
	Urn         string          `json:"urn"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Label       string          `json:"label"`
	CreatedAt   time.Time       `json:"created_at"`
	Owner       string          `json:"owner"`
	State       DataSystemState `json:"state"`
}

type DataSystemList struct {
	Entities []DataSystem `json:"entities"`
}

type DataSystemPostRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type DataSystemPostRequestEntityInfo struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type DataSystemPostRequest struct {
	Entity     DataSystemPostRequestEntity     `json:"entity"`
	EntityInfo DataSystemPostRequestEntityInfo `json:"entity_info"`
}

type DataSystemPostResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
}

type DataSystemPutRequest struct {
	Entity DataSystemPutRequestEntity `json:"entity"`
}

type DataSystemPutRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type DataSystemPutResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
}