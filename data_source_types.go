package neos

import (
	"time"
)

type DataSourceState struct {
	State   string `json:"state"`
	Healthy bool   `json:"healthy"`
}

type DataSource struct {
	Identifier  string          `json:"identifier"`
	Urn         string          `json:"urn"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Label       string          `json:"label"`
	CreatedAt   time.Time       `json:"created_at"`
	Owner       string          `json:"owner"`
	State       DataSourceState `json:"state"`
}

type DataSourceList struct {
	Entities []DataSource `json:"entities"`
}

type DataSourcePostRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type DataSourcePostRequestEntityInfo struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type DataSourcePostRequest struct {
	Entity     DataSourcePostRequestEntity     `json:"entity"`
	EntityInfo DataSourcePostRequestEntityInfo `json:"entity_info"`
}

type DataSourcePostResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
}

type DataSourcePutRequest struct {
	Entity DataSourcePutRequestEntity `json:"entity"`
}

type DataSourcePutRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type DataSourcePutRequestEntityInfo struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type DataSourcePutResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
}

type DataSourcePutInfoResponse struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

