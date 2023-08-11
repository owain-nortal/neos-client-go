package neos

import (
	"time"
)

// type DataUnitState struct {
// 	State   string `json:"state"`
// 	Healthy bool   `json:"healthy"`
// }

// type DataUnit struct {
// 	Identifier  string        `json:"identifier"`
// 	Urn         string        `json:"urn"`
// 	Name        string        `json:"name"`
// 	Description string        `json:"description"`
// 	Label       string        `json:"label"`
// 	CreatedAt   time.Time     `json:"created_at"`
// 	Owner       string        `json:"owner"`
// 	State       DataUnitState `json:"state"`
// }

// type DataUnitList struct {
// 	Entities []DataUnit `json:"entities"`
// }

// type DataUnitPostRequestEntity struct {
// 	Name        string `json:"name"`
// 	Label       string `json:"label"`
// 	Description string `json:"description"`
// }

// type DataUnitPostRequestEntityInfo struct {
// 	Owner      string   `json:"owner"`
// 	ContactIds []string `json:"contact_ids"`
// 	Links      []string `json:"links"`
// }

// type DataUnitPostRequest struct {
// 	Entity     DataUnitPostRequestEntity     `json:"entity"`
// 	EntityInfo DataUnitPostRequestEntityInfo `json:"entity_info"`
// }

// type DataUnitPostResponse struct {
// 	Identifier  string    `json:"identifier"`
// 	Urn         string    `json:"urn"`
// 	Name        string    `json:"name"`
// 	Description string    `json:"description"`
// 	Label       string    `json:"label"`
// 	CreatedAt   time.Time `json:"created_at"`
// }

// type DataUnitPutRequest struct {
// 	Entity DataUnitPutRequestEntity `json:"entity"`
// }

// type DataUnitPutRequestEntity struct {
// 	Name        string `json:"name"`
// 	Label       string `json:"label"`
// 	Description string `json:"description"`
// }

// type DataUnitPutRequestEntityInfo struct {
// 	Owner      string   `json:"owner"`
// 	ContactIds []string `json:"contact_ids"`
// 	Links      []string `json:"links"`
// }

// type DataUnitPutResponse struct {
// 	Identifier  string    `json:"identifier"`
// 	Urn         string    `json:"urn"`
// 	Name        string    `json:"name"`
// 	Description string    `json:"description"`
// 	Label       string    `json:"label"`
// 	CreatedAt   time.Time `json:"created_at"`
// }

// type DataUnitPutInfoResponse struct {
// 	Owner      string   `json:"owner"`
// 	ContactIds []string `json:"contact_ids"`
// 	Links      []string `json:"links"`
// }

type LinkState struct {
	Code    string `json:"code"`
	Reason  string `json:"reason"`
	Healthy bool   `json:"healthy"`
}

type LinkParent struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	IsSystem    bool      `json:"is_system"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
	State       LinkState `json:"state"`
	Owner       string    `json:"owner"`
	EntityType  string    `json:"entity_type"`
	OutputType  string    `json:"output_type"`
}

type LinkChild struct {
	Identifier  string    `json:"identifier"`
	Urn         string    `json:"urn"`
	Name        string    `json:"name"`
	IsSystem    bool      `json:"is_system"`
	Description string    `json:"description"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
	State       LinkState `json:"state"`
	Owner       string    `json:"owner"`
	EntityType  string    `json:"entity_type"`
	OutputType  string    `json:"output_type"`
}

type LinksGetResponse struct {
	Links []struct {
		Parent LinkParent `json:"parent"`
		Child  LinkChild  `json:"child"`
	} `json:"links"`
}

type LinkPostResponse struct {
	Parent struct {
		Identifier  string    `json:"identifier"`
		Urn         string    `json:"urn"`
		Name        string    `json:"name"`
		IsSystem    bool      `json:"is_system"`
		Description string    `json:"description"`
		Label       string    `json:"label"`
		CreatedAt   time.Time `json:"created_at"`
		State       struct {
			Code    string `json:"code"`
			Reason  string `json:"reason"`
			Healthy bool   `json:"healthy"`
		} `json:"state"`
		Owner      string `json:"owner"`
		EntityType string `json:"entity_type"`
		OutputType string `json:"output_type"`
	} `json:"parent"`
	Child struct {
		Identifier  string    `json:"identifier"`
		Urn         string    `json:"urn"`
		Name        string    `json:"name"`
		IsSystem    bool      `json:"is_system"`
		Description string    `json:"description"`
		Label       string    `json:"label"`
		CreatedAt   time.Time `json:"created_at"`
		State       struct {
			Code    string `json:"code"`
			Reason  string `json:"reason"`
			Healthy bool   `json:"healthy"`
		} `json:"state"`
		Owner      string `json:"owner"`
		EntityType string `json:"entity_type"`
		OutputType string `json:"output_type"`
	} `json:"child"`
}

type LinkDataSourceDataUnitPostRequest struct {
	ParentIdentifier string `json:"parent_identifier"`
	ChildIdentifier  string `json:"child_identifier"`
}
