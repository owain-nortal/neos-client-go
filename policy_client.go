package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type PolicyClient struct {
	hubUri string
	http   *NeosHttp
}

func NewPolicyClient(hubUri string, http *NeosHttp) *PolicyClient {
	return &PolicyClient{
		hubUri: hubUri,
		http:   http,
	}
}

// func (c *PolicyClient) List(filter string) (GroupList, error) {
// 	var rtn GroupList
// 	requestURL := fmt.Sprintf("%s/api/hub/iam/group%s", c.hubUri, filterQuery(filter, "account"))
// 	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
// 	return rtn, err
// }

func (c *PolicyClient) Get(id string, account string) (Policy, error) {
	var rtn Policy
	acc := ""
	if account != "" {
		acc = fmt.Sprintf("&account=%s", account)
	}
	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user&user_nrn=%s%s", c.hubUri, id, acc)
	body, err := c.http.Get(requestURL, http.StatusOK)
	rtn.Policy = string(body[:])
	return rtn, err
}

func (c *PolicyClient) Post(ctx context.Context, dspr PolicyPostRequest, account string) (PolicyPostResponse, error) {
	var rtn PolicyPostResponse
	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user%s", c.hubUri, filterQuery(account, "account"))
	err := c.http.PostUnmarshal(requestURL, dspr.Policy, http.StatusOK, &rtn)
	return rtn, err
}

func (c *PolicyClient) Put(ctx context.Context, id string, dspr PolicyPutRequest, account string) (PolicyPutResponse, error) {
	var rtn PolicyPutResponse

	query := ""
	if id != "" {
		query = fmt.Sprintf("%s&user_nrn=%s", query, id)
	}
	if account != "" {
		query = fmt.Sprintf("%s&account=%s", query, account)
	}
	if query != "" {
		query = fmt.Sprintf("?%s", query)
	}
	query = strings.Replace(query, "?&", "?", -1)

	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user%s", c.hubUri, query)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *PolicyClient) Delete(ctx context.Context, id string, account string) error {
	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user?user_nrn=%s", c.hubUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}
