package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type UserClient struct {
	hubUri  string
	http    *NeosHttp
	Account string
}

func NewUserClient(hubUri string, http *NeosHttp, account string) *UserClient {
	return &UserClient{
		hubUri:  hubUri,
		http:    http,
		Account: account,
	}
}

func (c *UserClient) List(search, system, account string) (UserList, error) {
	var rtn UserList

	if c.accountIsNotRootOrEmpty(account) {		
		c.http.AddHeader("x-account-override", account)
	}

	query := ""
	if search != "" {
		query = fmt.Sprintf("%s&search=%s", query, search)
	}
	if system != "" {
		query = fmt.Sprintf("%s&system=%s", query, system)
	}
	if account != "" {
		query = fmt.Sprintf("%s&account=%s", query, account)
	}
	if query != "" {
		query = fmt.Sprintf("?%s", query)
	}
	query = strings.Replace(query, "?&", "?", -1)
	requestURL := fmt.Sprintf("%s/api/hub/iam/users%s", c.hubUri, query)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *UserClient) accountIsNotRootOrEmpty(account string) bool {
	return account != "" && account != "root"
}

func (c *UserClient) Post(ctx context.Context, dspr UserPostRequest, account string) (User, error) {
	var rtn User

	acc := ""
	if c.accountIsNotRootOrEmpty(account) {
		acc = fmt.Sprintf("?account=%s",account)
		c.http.AddHeader("x-account-override", account)
	}
	
	requestURL := fmt.Sprintf("%s/api/hub/iam/user%s", c.hubUri, acc)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *UserClient) Delete(ctx context.Context, id string, account string) error {
	
	if c.accountIsNotRootOrEmpty(account) {		
		c.http.AddHeader("x-account-override", account)
	}

	requestURL := fmt.Sprintf("%s/api/hub/iam/user/%s%s", c.hubUri, id, filterQuery(account, "account"))
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *UserClient) Purge(ctx context.Context, id string, account string) error {

	if c.accountIsNotRootOrEmpty(account) {		
		c.http.AddHeader("x-account-override", account)
	}

	requestURL := fmt.Sprintf("%s/api/hub/iam/user/%s/purge%s", c.hubUri, id, filterQuery(account, "account"))
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}
