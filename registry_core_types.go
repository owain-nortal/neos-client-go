package neos

import ()

type RegistryCore struct {
	ID        string `json:"id"`
	Host      string `json:"host"`
	Urn       string `json:"urn"`
	Name      string `json:"name"`
	AccessKey string `json:"access_key"`
}

type RegistryCoreList struct {
	Cores []RegistryCore `json:"cores"`
}

type RegistryCorePostRequest struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
}

type RegistryCoreDeleteRequest struct {
	Urn string `json:"urn"`
}

type RegistryCorePostResponse struct {
	Identifier string `json:"identifier"`
	Urn        string `json:"urn"`
	AccessKey  string `json:"access_key"`
}
