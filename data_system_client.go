package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type DataSystemClient struct {
	coreUri string
	http    *NeosHttp
}

func NewDataSystemClient(coreUri string, http *NeosHttp) *DataSystemClient {
	return &DataSystemClient{
		coreUri: coreUri,
		http:    http,
	}
}

func (c *DataSystemClient) Delete(ctx context.Context, id string) error {
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_system/%s", c.coreUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *DataSystemClient) Post(ctx context.Context, dspr DataSystemPostRequest) (DataSystemPostResponse, error) {
	var rtn DataSystemPostResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_system", c.coreUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataSystemClient) Put(ctx context.Context, id string, dspr DataSystemPutRequest) (DataSystemPutResponse, error) {
	var rtn DataSystemPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_system/%s", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataSystemClient) PutInfo(ctx context.Context, id string, dspr DataSystemPutRequestEntityInfo) (DataSystemPutInfoResponse, error) {
	var rtn DataSystemPutInfoResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_system/%s/info", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataSystemClient) Get() (DataSystemList, error) {
	var rtn DataSystemList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_system", c.coreUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}
