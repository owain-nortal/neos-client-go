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
	acc := ""
	if account != "" {
		acc = fmt.Sprintf("&account=%s", account)
	}
	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user?user_nrn=%s%s", c.hubUri, nrn, acc)
	body, err := c.http.Get(requestURL, http.StatusOK)
	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}
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
	acc := ""
	if account != "" {
		acc = fmt.Sprintf("&account=%s", account)
	}
	res := ""
	if resouce != "" {
		res = fmt.Sprintf("&resouce=%s", resouce)
	}
	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/users?page=%d&page_size=%d%s%s", c.hubUri, page, pageSize, acc, res)
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
	if account != "" {
		query = fmt.Sprintf("%s&account=%s", query, account)
	}
	if query != "" {
		query = fmt.Sprintf("?%s", query)
	}
	query = strings.Replace(query, "?&", "?", -1)

	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user%s", c.hubUri, query)

	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}

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

	acc := ""
	if c.accountIsNotRootOrEmpty(account) {
		acc = fmt.Sprintf("&account=%s", account)
		c.http.AddHeader("x-account-override", account)
	}

	requestURL := fmt.Sprintf("%s/api/hub/iam/policy/user?user_nrn=%s%s", c.hubUri, id, acc)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}
