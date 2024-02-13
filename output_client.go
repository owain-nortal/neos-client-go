package neos

import (
	"context"
	"fmt"
	"net/http"
	"github.com/pkg/errors"
)

type OutputClient struct {
	coreUri string
	http    *NeosHttp
}

func NewOutputClient(coreUri string, http *NeosHttp) *OutputClient {
	return &OutputClient{
		coreUri: coreUri,
		http:    http,
	}
}

func (c *OutputClient) Delete(ctx context.Context, id string) error {
	requestURL := fmt.Sprintf("%s/api/gateway/v2/output/%s", c.coreUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *OutputClient) Post(ctx context.Context, dspr OutputPostRequest) (OutputPostResponse, error) {
	var rtn OutputPostResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/output", c.coreUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *OutputClient) Put(ctx context.Context, id string, dspr OutputPutRequest) (OutputPutResponse, error) {
	var rtn OutputPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/output/%s", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *OutputClient) PutInfo(ctx context.Context, id string, dspr OutputPutRequestEntityInfo) (OutputPutInfoResponse, error) {
	var rtn OutputPutInfoResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/output/%s/info", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *OutputClient) Get() (OutputList, error) {
	var rtn OutputList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/output", c.coreUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}
