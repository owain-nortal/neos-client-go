package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type AccountClient struct {
	hubUri  string
	http    *NeosHttp
	Account string
}

func NewAccountClient(hubUri string, http *NeosHttp, account string) *AccountClient {
	return &AccountClient{
		hubUri:  hubUri,
		http:    http,
		Account: account,
	}
}

func (c *AccountClient) Get(filter string) (AccountList, error) {
	var rtn AccountList

	requestURL := fmt.Sprintf("%s/api/hub/iam/account%s", c.hubUri, filterQuery(filter, "account"))
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *AccountClient) accountIsNotRootOrEmpty(account string) bool {
	return account != "" && account != "root"
}

func (c *AccountClient) Post(ctx context.Context, dspr AccountPostRequest) (Account, error) {
	var rtn Account

	if c.accountIsNotRootOrEmpty(dspr.Name) {
		c.http.AddHeader("x-account-override", dspr.Name)
	}
	c.http.AddHeader("x-account", dspr.Name)

	requestURL := fmt.Sprintf("%s/api/hub/iam/account", c.hubUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *AccountClient) Put(ctx context.Context, id string, dspr AccountPutRequest) (Account, error) {
	var rtn Account

	if c.accountIsNotRootOrEmpty(id) {
		c.http.AddHeader("x-account-override", id)
	}
	c.http.AddHeader("x-account", id)

	requestURL := fmt.Sprintf("%s/api/hub/iam/account/%s", c.hubUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *AccountClient) Delete(ctx context.Context, id string) error {
	if c.accountIsNotRootOrEmpty(id) {
		c.http.AddHeader("x-account-override", id)
	}
	c.http.AddHeader("x-account", id)
	
	requestURL := fmt.Sprintf("%s/api/hub/iam/account/%s", c.hubUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}
