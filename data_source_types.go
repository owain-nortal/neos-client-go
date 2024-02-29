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


type DataSourceGetResponse struct {
	// Connection struct {
	// 	ConnectionType string `json:"connection_type"`
	// 	Database       string `json:"database"`
	// 	Engine         string `json:"engine"`
	// 	Host           string `json:"host"`
	// 	Password       struct {
	// 		EnvKey string `json:"env_key"`
	// 	} `json:"password"`
	// 	Port   int    `json:"port"`
	// 	Schema string `json:"schema"`
	// 	User   struct {
	// 		EnvKey string `json:"env_key"`
	// 	} `json:"user"`
	// } `json:"connection"`
	Entity struct {
		CreatedAt   time.Time `json:"created_at"`
		Description string    `json:"description"`
		Identifier  string    `json:"identifier"`
		IsSystem    bool      `json:"is_system"`
		Label       string    `json:"label"`
		Name        string    `json:"name"`
		Owner       string    `json:"owner"`
		State       struct {
			Code    string `json:"code"`
			Healthy bool   `json:"healthy"`
			Reason  string `json:"reason"`
		} `json:"state"`
		Urn string `json:"urn"`
	} `json:"entity"`
	EntityInfo struct {
		ContactIds []string `json:"contact_ids"`
		Links      []string `json:"links"`
		Owner      string   `json:"owner"`
	} `json:"entity_info"`
	Links struct {
		Children []struct {
			CreatedAt   time.Time `json:"created_at"`
			Description string    `json:"description"`
			EntityType  string    `json:"entity_type"`
			Identifier  string    `json:"identifier"`
			IsSystem    bool      `json:"is_system"`
			Label       string    `json:"label"`
			Name        string    `json:"name"`
			Owner       string    `json:"owner"`
			State       struct {
				Code    string `json:"code"`
				Healthy bool   `json:"healthy"`
				Reason  string `json:"reason"`
			} `json:"state"`
			Urn string `json:"urn"`
		} `json:"children"`
		Parents []struct {
			CreatedAt   time.Time `json:"created_at"`
			Description string    `json:"description"`
			EntityType  string    `json:"entity_type"`
			Identifier  string    `json:"identifier"`
			IsSystem    bool      `json:"is_system"`
			Label       string    `json:"label"`
			Name        string    `json:"name"`
			Owner       string    `json:"owner"`
			State       struct {
				Code    string `json:"code"`
				Healthy bool   `json:"healthy"`
				Reason  string `json:"reason"`
			} `json:"state"`
			Urn string `json:"urn"`
		} `json:"parents"`
	} `json:"links"`
	SecretIdentifier string `json:"secret_identifier"`
	SparkIdentifier  string `json:"spark_identifier"`
}