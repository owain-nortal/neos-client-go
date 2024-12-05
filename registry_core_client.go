package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type RegistryCoreClient struct {
	registryUri string
	http        *NeosHttp
	Account     string
}

func NewRegistryCoreClient(registryUri string, http *NeosHttp, account string) *RegistryCoreClient {
	return &RegistryCoreClient{
		registryUri: registryUri,
		http:        http,
		Account:     account,
	}
}

func (c *RegistryCoreClient) Delete(ctx context.Context, id string, account string) error {
	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}
	c.http.AddHeader("x-account", account)
	requestURL := fmt.Sprintf("%s/api/hub/registry/core/%s", c.registryUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *RegistryCoreClient) accountIsNotRootOrEmpty(account string) bool {
	return account != "" && account != "root"
}

func (c *RegistryCoreClient) Post(ctx context.Context, dspr RegistryCorePostRequest, account string) (RegistryCorePostResponse, error) {
	var rtn RegistryCorePostResponse
	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}
	c.http.AddHeader("x-account", account)

	requestURL := fmt.Sprintf("%s/api/hub/registry/core", c.registryUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *RegistryCoreClient) Get(account string) (RegistryCoreList, error) {
	var rtn RegistryCoreList
	if c.accountIsNotRootOrEmpty(account) {
		c.http.AddHeader("x-account-override", account)
	}
	c.http.AddHeader("x-account", account)
	requestURL := fmt.Sprintf("%s/api/hub/registry/core", c.registryUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}
