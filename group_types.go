package neos

import ()

type GroupPostRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GroupPutRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type Group struct {
	Identifier  string   `json:"identifier"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	IsSystem    bool     `json:"is_system"`
	Principals  []string `json:"principals"`
}

type GroupList struct {
	Groups []Group `json:"groups"`
}

type GroupPrincipalPostRequest struct {
	Principals  []string `json:"principals"`
}

type GroupPrincipalDeleteRequest struct {
	Principals  []string `json:"principals"`
}
