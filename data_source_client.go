package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type DataSourceClient struct {
	coreUri string
	http    *NeosHttp
	Account string
}

func NewDataSourceClient(coreUri string, http *NeosHttp,account string) *DataSourceClient {
	return &DataSourceClient{
		coreUri: coreUri,
		http:    http,
		Account: account,
	}
}

func (c *DataSourceClient) Delete(ctx context.Context, id string) error {
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s", c.coreUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *DataSourceClient) Post(ctx context.Context, dspr DataSourcePostRequest) (DataSourcePostResponse, error) {
	var rtn DataSourcePostResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source", c.coreUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataSourceClient) Put(ctx context.Context, id string, dspr DataSourcePutRequest) (DataSourcePutResponse, error) {
	var rtn DataSourcePutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataSourceClient) PutInfo(ctx context.Context, id string, dspr DataSourcePutRequestEntityInfo) (DataSourcePutInfoResponse, error) {
	var rtn DataSourcePutInfoResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s/info", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataSourceClient) Get() (DataSourceList, error) {
	var rtn DataSourceList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source", c.coreUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataSourceClient) GetById(id string) (DataSourceGetResponse, error) {
	var rtn DataSourceGetResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s", c.coreUri, id)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}
