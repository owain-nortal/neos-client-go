package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type SecretClient struct {
	coreUri string
	http    *NeosHttp
	Account string
}

func NewSecretClient(coreUri string, http *NeosHttp, account string) *SecretClient {
	return &SecretClient{
		coreUri: coreUri,
		http:    http,
		Account: account,
	}
}

func (c *SecretClient) Delete(ctx context.Context, id string) error {
	requestURL := fmt.Sprintf("%s/api/gateway/v2/secret/%s", c.coreUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *SecretClient) Post(ctx context.Context, spr SecretPostRequest) (SecretPostResponse, error) {
	var rtn SecretPostResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/secret", c.coreUri)
	err := c.http.PostUnmarshal(requestURL, spr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *SecretClient) Put(ctx context.Context, id string, dspr SecretPutRequest) (SecretPutResponse, error) {
	var rtn SecretPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/secret/%s", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *SecretClient) Get() (SecretList, error) {
	var rtn SecretList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/secret", c.coreUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *SecretClient) GetById(id string) (Secret, error) {
	var rtn Secret
	requestURL := fmt.Sprintf("%s/api/gateway/v2/secret/%s", c.coreUri, id)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}
