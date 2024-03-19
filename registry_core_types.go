package neos

import ()

type RegistryCore struct {
	ID        string `json:"id"`
	Host      string `json:"host"`
	Urn       string `json:"urn"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	AccessKey string `json:"access_key"`
	Account   string `json:"account"`
}

type RegistryCoreList struct {
	Cores []RegistryCore `json:"cores"`
}

type RegistryCorePostRequest struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
}

type RegistryCoreDeleteRequest struct {
	Id string `json:"id"`
}

type RegistryCoreKeyPairPostResponse struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	IsSysyem        bool   `json:"is_system"`
}

type RegistryCorePostResponse struct {
	Identifier string                          `json:"identifier"`
	Urn        string                          `json:"urn"`
	KeyPair    RegistryCoreKeyPairPostResponse `json:"keypair"`
}
