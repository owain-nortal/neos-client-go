package neos

import ()

type AccountPostRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
}

type AccountPutRequest struct {
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
}

type Account struct {
	Identifier  string `json:"identifier"`
	Urn         string `json:"urn"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	IsSystem    bool   `json:"is_system"`
}

type AccountList struct {
	Accounts []Account `json:"accounts"`
}
