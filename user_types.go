package neos

import ()

type User struct {
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Identifier string `json:"identifier"`
	Urn        string `json:"urn"`
	IsSystem   bool `json:"is_system"`
	Enabled    bool   `json:"enabled"`
	Account    string `json:"account"`
}

type UserPostRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserList struct {
	Users []User `json:"users"`
}
