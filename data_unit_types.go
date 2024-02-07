package neos

import (
	"time"
)

type DataUnitState struct {
	State   string `json:"state"`
	Healthy bool   `json:"healthy"`
}

type DataUnit struct {
	Identifier  string        `json:"identifier"`
	Urn         string        `json:"urn"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Label       string        `json:"label"`
	CreatedAt   time.Time     `json:"created_at"`
	Owner       string        `json:"owner"`
	State       DataUnitState `json:"state"`
}

type DataUnitList struct {
	Entities []DataUnit `json:"entities"`
}

type DataUnitPostRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type DataUnitPostRequestEntityInfo struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type DataUnitPostRequest struct {
	Entity     DataUnitPostRequestEntity     `json:"entity"`
	EntityInfo DataUnitPostRequestEntityInfo `json:"entity_info"`
}

type DataUnitPostResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
}

type DataUnitPutRequest struct {
	Entity DataUnitPutRequestEntity `json:"entity"`
}

type DataUnitPutRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type DataUnitPutRequestEntityInfo struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type DataUnitPutResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
}

type DataUnitPutInfoResponse struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type DataUnitConfigurationCSVPutRequest struct {
	Configuration DataUnitConfigurationCSVConfigPutRequest `json:"configuration"`
}

type DataUnitConfigurationCSVConfigPutRequest struct {
	DateUnitType string `json:"data_unit_type"`
	Delimiter    string `json:"delimiter"`
	Path         string `json:"path"`
	HasHeader    bool   `json:"has_header"`
	QuoteChar    string `json:"quote_char"`
	EscapeChar   string `json:"escape_char"`
}

type DataUnitConfigurationCSVPutResponse struct {
	Configuration DataUnitConfigurationCSVConfigPutResponse `json:"configuration"`
}

type DataUnitConfigurationCSVConfigPutResponse struct {
	DateUnitType string `json:"data_unit_type"`
	Delimiter    string `json:"delimiter"`
	Path         string `json:"path"`
	HasHeader    bool   `json:"has_header"`
	QuoteChar    string `json:"quote_char"`
	EscapeChar   string `json:"escape_char"`
}

type DataUnitConfigurationParquetPutRequest struct {
	Configuration DataUnitConfigurationParquetConfigPutRequest `json:"configuration"`
}

type DataUnitConfigurationParquetConfigPutRequest struct {
	DateUnitType string `json:"data_unit_type"`
}

type DataUnitConfigurationParquetPutResponse struct {
	Configuration DataUnitConfigurationParquetConfigPutResponse `json:"configuration"`
}
type DataUnitConfigurationParquetConfigPutResponse struct {
	DateUnitType string `json:"data_unit_type"`
}

type DataUnitConfigurationTablePutRequest struct {
	Configuration DataUnitConfigurationTableConfigPutRequest `json:"configuration"`
}

type DataUnitConfigurationTableConfigPutRequest struct {
	DateUnitType string `json:"data_unit_type"`
	Table        string `json:"table"`
}

type DataUnitConfigurationTablePutResponse struct {
	Configuration DataUnitConfigurationTableConfigPutResponse `json:"configuration"`
}

type DataUnitConfigurationTableConfigPutResponse struct {
	DateUnitType string `json:"data_unit_type"`
	Table        string `json:"table"`
}

type DataUnitConfigurationQueryPutRequest struct {
	Configuration DataUnitConfigurationQueryConfigPutRequest `json:"configuration"`
}

type DataUnitConfigurationQueryConfigPutRequest struct {
	DateUnitType string `json:"data_unit_type"`
	Query        string `json:"query"`
}

type DataUnitConfigurationQueryPutResponse struct {
	Configuration DataUnitConfigurationQueryConfigPutResponse `json:"configuration"`
}

type DataUnitConfigurationQueryConfigPutResponse struct {
	DateUnitType string `json:"data_unit_type"`
	Query        string `json:"query"`
}

type DataUnitConfigurationDataProductPutRequest struct {
	Configuration DataUnitConfigurationDataProductConfigPutRequest `json:"configuration"`
}

type DataUnitConfigurationDataProductConfigPutRequest struct {
	DateUnitType string `json:"data_unit_type"`
	Engine       string `json:"engine"`
	Table        string `json:"table"`
}

type DataUnitConfigurationDataProductPutResponse struct {
	Configuration DataUnitConfigurationDataProductConfigPutResponse `json:"configuration"`
}

type DataUnitConfigurationDataProductConfigPutResponse struct {
	DateUnitType string `json:"data_unit_type"`
	Engine       string `json:"engine"`
	Table        string `json:"table"`
}







type DataUnitConfigurationCSVGetResponse struct {
	Configuration DataUnitConfigurationCSVConfigGetResponse `json:"configuration"`
}

type DataUnitConfigurationCSVConfigGetResponse struct {
	DateUnitType string `json:"data_unit_type"`
	Delimiter    string `json:"delimiter"`
	Path         string `json:"path"`
	HasHeader    bool   `json:"has_header"`
	QuoteChar    string `json:"quote_char"`
	EscapeChar   string `json:"escape_char"`
}

type DataUnitConfigurationParquetGetResponse struct {
	Configuration DataUnitConfigurationParquetConfigGetResponse `json:"configuration"`
}

type DataUnitConfigurationParquetConfigGetResponse struct {
	DateUnitType string `json:"data_unit_type"`
}

type DataUnitConfigurationTableGetResponse struct {
	Configuration DataUnitConfigurationTableConfigGetResponse `json:"configuration"`
}

type DataUnitConfigurationTableConfigGetResponse struct {
	DateUnitType string `json:"data_unit_type"`
	Table        string `json:"table"`
}

type DataUnitConfigurationQueryGetResponse struct {
	Configuration DataUnitConfigurationQueryConfigGetResponse `json:"configuration"`
}

type DataUnitConfigurationQueryConfigGetResponse struct {
	DateUnitType string `json:"data_unit_type"`
	Query        string `json:"query"`
}

type DataUnitConfigurationDataProductGetResponse struct {
	Configuration DataUnitConfigurationDataProductConfigGetResponse `json:"configuration"`
}

type DataUnitConfigurationDataProductConfigGetResponse struct {
	DateUnitType string `json:"data_unit_type"`
	Engine       string `json:"engine"`
	Table        string `json:"table"`
}


type DataUnitConfigurationDataUnitTypeOnlyConfigGetResponse struct {
	DateUnitType string `json:"data_unit_type"`
}
