package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type GroupClient struct {
	hubUri  string
	http    *NeosHttp
	Account string
}

func NewGroupClient(hubUri string, http *NeosHttp, account string) *GroupClient {
	return &GroupClient{
		hubUri: hubUri,
		http:   http,
	}
}

func (c *GroupClient) List(account string) (GroupList, error) {
	var rtn GroupList
	acc := ""
	if c.accountIsNotRootOrEmpty(account) {
		acc = fmt.Sprintf("?account=%s", account)
		c.http.AddHeader("x-account-override", account)
	}

	requestURL := fmt.Sprintf("%s/api/hub/iam/group%s", c.hubUri, acc)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) Get(id string, account string) (Group, error) {
	var rtn Group
	acc := ""
	if c.accountIsNotRootOrEmpty(account) {
		acc = fmt.Sprintf("?account=%s", account)
		c.http.AddHeader("x-account-override", account)
	}
	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s%s", c.hubUri, id, acc)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) accountIsNotRootOrEmpty(account string) bool {
	return account != "" && account != "root"
}

func (c *GroupClient) Post(ctx context.Context, dspr GroupPostRequest, account string) (Group, error) {
	var rtn Group

	acc := ""
	if c.accountIsNotRootOrEmpty(account) {
		acc = fmt.Sprintf("?account=%s", account)
		c.http.AddHeader("x-account-override", account)
	}

	requestURL := fmt.Sprintf("%s/api/hub/iam/group%s", c.hubUri, acc)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) Put(ctx context.Context, id string, dspr GroupPutRequest, account string) (Group, error) {
	var rtn Group

	acc := ""
	if c.accountIsNotRootOrEmpty(account) {
		acc = fmt.Sprintf("?account=%s", account)
		c.http.AddHeader("x-account-override", account)
	}

	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s%s", c.hubUri, id, acc)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) Delete(ctx context.Context, id string, account string) error {
	acc := ""
	if c.accountIsNotRootOrEmpty(account) {
		acc = fmt.Sprintf("?account=%s", account)
		c.http.AddHeader("x-account-override", account)
	}

	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s%s", c.hubUri, id, acc)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *GroupClient) PrincipalsPost(ctx context.Context, id string, dspr GroupPrincipalPostRequest, account string) (Group, error) {
	var rtn Group
	acc := ""
	if c.accountIsNotRootOrEmpty(account) {
		acc = fmt.Sprintf("?account=%s", account)
		c.http.AddHeader("x-account-override", account)
	}

	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s/principals%s", c.hubUri, id, acc)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) PrincipalsDelete(ctx context.Context, id string, dspr GroupPrincipalDeleteRequest, account string) (Group, error) {
	var rtn Group
	acc := ""
	if c.accountIsNotRootOrEmpty(account) {
		acc = fmt.Sprintf("?account=%s", account)
		c.http.AddHeader("x-account-override", account)
	}

	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s/principals%s", c.hubUri, id, acc)
	err := c.http.DeleteUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, "error doing http delete")
	}
	return rtn, err
}
