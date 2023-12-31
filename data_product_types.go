package neos

import (
	"time"
)

type DataProductState struct {
	State   string `json:"state"`
	Healthy bool   `json:"healthy"`
}

type DataProduct struct {
	Identifier  string           `json:"identifier"`
	Urn         string           `json:"urn"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Label       string           `json:"label"`
	CreatedAt   time.Time        `json:"created_at"`
	Owner       string           `json:"owner"`
	State       DataProductState `json:"state"`
}

type DataProductList struct {
	Entities []DataProduct `json:"entities"`
}

type DataProductPostRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type DataProductPostRequestEntityInfo struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type DataProductPostRequest struct {
	Entity     DataProductPostRequestEntity     `json:"entity"`
	EntityInfo DataProductPostRequestEntityInfo `json:"entity_info"`
}

type DataProductPostResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
}

type DataProductPutRequest struct {
	Entity DataProductPutRequestEntity `json:"entity"`
}

type DataProductPutRequestEntity struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type DataProductPutRequestEntityInfo struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type DataProductPutResponse struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
}

type DataProductPutInfoResponse struct {
	Owner      string   `json:"owner"`
	ContactIds []string `json:"contact_ids"`
	Links      []string `json:"links"`
}

type DataProductSchema struct {
	Fields []DataProductSchemaField `json:"fields"`
}

type DataProductSchemaField struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Primary     bool                      `json:"primary"`
	Optional    bool                      `json:"optional"`
	DataType    DataProductSchemaDataType `json:"data_type"`
	//Tags        []string                  `json:"tags"`
}

type DataProductSchemaDataType struct {
	Meta       map[string]string `json:"meta"`
	ColumnType string            `json:"column_type"`
}

type DataProductSchemaPutRequest struct {
	Details DataProductSchemaDetailsPutRequest `json:"details"`
}

type DataProductSchemaDetailsPutRequest struct {
	ProductType string                             `json:"product_type"`
	Fields      []DataProductSchemaFieldPutRequest `json:"fields"`
}

type DataProductSchemaFieldPutRequest struct {
	Description string                              `json:"description"`
	Name        string                              `json:"name"`
	Primary     bool                                `json:"primary"`
	Optional    bool                                `json:"optional"`
	DataType    DataProductSchemaDataTypePutRequest `json:"data_type"`
	//Type        string                              `json:"type"`
	//Tags        []string                            `json:"tags"`
}

type DataProductSchemaDataTypePutRequest struct {
	Meta       map[string]string `json:"meta"`
	ColumnType string            `json:"column_type"`
}

type DataProductSchemaPutResponse struct {
	Fields []struct {
		//Type        string `json:"type"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Primary     bool   `json:"primary"`
		Optional    bool   `json:"optional"`
		DataType    struct {
			Meta       map[string]string `json:"meta"`
			ColumnType string            `json:"column_type"`
		} `json:"data_type"`
		//Tags []string `json:"tags"`
	} `json:"fields"`
}

type DataProductBuilderPutResponse struct {}
