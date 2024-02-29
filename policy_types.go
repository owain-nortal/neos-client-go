package neos

import ()

type PolicyPostRequest struct {
	Policy string `json:"policy"`
}

type PolicyPostResponse struct {
	Policy string `json:"policy"`
}

type PolicyPutRequest struct {
	Policy string `json:"policy"`
}

type PolicyPutResponse struct {
	Policy string `json:"policy"`
}

type PolicyDeleteResponse struct {
}

type Policy struct {
	Policy string `json:"policy"`
}

type UserPolicyList struct {
	UserPolicies []UserPolicy `json:"user_policies"`
}

type UserPolicy struct {
	User     string         `json:"user"`
	Policy   UserPolicyType `json:"policy"`
	IsSystem bool           `json:"is_system"`
}

type UserPolicyType struct {
	Version    string `json:"version"`
	Statements []struct {
		Sid       string   `json:"sid"`
		Principal string   `json:"principal"`
		Action    []string `json:"action"`
		Resource  []string `json:"resource"`
		Condition []string `json:"condition"`
		Effect    string   `json:"effect"`
	} `json:"statements"`
}
