package neos

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type PolicyClient struct {
	hubUri  string
	http    *NeosHttp
	Account string
}

func NewPolicyClient(hubUri string, http *NeosHttp, account string) *PolicyClient {
	return &PolicyClient{
		hubUri:  hubUri,
		http:    http,
		Account: account,
	}
}

// func (c *PolicyClient) List(filter string) (GroupList, error) {
// 	var rtn GroupList
// 	requestURL := fmt.Sprintf("%s/api/hub/iam/group%s", c.hubUri, filterQuery(filter, "account"))
// 	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
// 	return rtn, err
// }

func (c *PolicyClient) Get(nrn string, account string) (Policy, error) {
	var rtn Policy
	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user?user_nrn=%s", c.hubUri, nrn)
	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}
	c.http.AddHeader("x-account", account)

	body, err := c.http.Get(requestURL, http.StatusOK)
	if err != nil {
		return rtn, err
	}

	json, err := c.NormalizeJson(string(body))
	if err != nil {
		return rtn, errors.Wrap(err, "error normalizing json")
	}

	rtn.Policy = json
	return rtn, err
}

func (c *PolicyClient) List(resouce string, account string) (UserPolicyList, error) {
	page := 1
	pageSize := 50
	var rtn UserPolicyList

	res := ""
	if resouce != "" {
		res = fmt.Sprintf("&resouce=%s", resouce)
	}

	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}
	c.http.AddHeader("x-account", account)

	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/users?page=%d&page_size=%d%s", c.hubUri, page, pageSize, res)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *PolicyClient) Post(ctx context.Context, dspr PolicyPostRequest, account string) (PolicyPostResponse, error) {
	rtn := PolicyPostResponse{}
	json := ""

	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user%s", c.hubUri, filterQuery(account, "account"))

	json, err := c.NormalizeJson(dspr.Policy)
	if err != nil {
		return rtn, errors.Wrap(err, "error normalizing json")
	}

	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}
	c.http.AddHeader("x-account", account)

	res, err := c.http.Post(requestURL, json, http.StatusOK)
	x := string(res)
	fmt.Println(x)
	rtn.Policy = json
	return rtn, err
}

func (c *PolicyClient) Put(ctx context.Context, id string, dspr PolicyPutRequest, account string) (PolicyPutResponse, error) {

	var r PolicyPutResponse
	query := ""
	if id != "" {
		query = fmt.Sprintf("%s&user_nrn=%s", query, id)
	}

	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}
	c.http.AddHeader("x-account", account)

	if query != "" {
		query = fmt.Sprintf("?%s", query)
	}
	query = strings.Replace(query, "?&", "?", -1)

	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user%s", c.hubUri, query)

	json, err := c.NormalizeJson(dspr.Policy)
	if err != nil {
		return r, errors.Wrap(err, "error normalizing json")
	}
	res, err := c.http.Put(requestURL, json, http.StatusOK)
	r = PolicyPutResponse{Policy: string(res)}
	return r, err
}

func (c *PolicyClient) NormalizeJson(j string) (string, error) {
	up, err := c.ConvertJsonToUserPolicy(j)
	if err != nil {
		return "", err
	}
	return c.ConvertUserPolicyToJson(up)
}

func (c *PolicyClient) ConvertJsonToUserPolicy(j string) (UserPolicy, error) {
	doc := []byte(j)
	rtn := UserPolicy{}
	err := json.Unmarshal(doc, &rtn)
	return rtn, err
}

func (c *PolicyClient) ConvertUserPolicyToJson(userPolicy UserPolicy) (string, error) {
	b, err := json.Marshal(userPolicy)
	return string(b), err
}

func (c *PolicyClient) accountIsNotRootOrEmpty(account string) bool {
	return account != "" && account != "root"
}

func (c *PolicyClient) Delete(ctx context.Context, id string, account string) error {

	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}

	c.http.AddHeader("x-account", account)

	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user?user_nrn=%s", c.hubUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}
