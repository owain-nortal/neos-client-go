package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type DataSourceConnectionClient struct {
	coreUri string
	http    *NeosHttp
	Account string
}

func NewDataSourceConnectionClient(coreUri string, http *NeosHttp, account string) *DataSourceConnectionClient {
	return &DataSourceConnectionClient{
		coreUri: coreUri,
		http:    http,
		Account: account,
	}
}

func (c *DataSourceConnectionClient) Put(ctx context.Context, id string, doc string) (string, error) {
	rtn := ""
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s/connection", c.coreUri, id)
	res, err := c.http.PutRaw(requestURL, doc, http.StatusOK)

	if err != nil {
		return rtn, errors.Wrap(err, "error doing http put")

	}

	return string(res), err
}

func (c *DataSourceConnectionClient) Get(id string) (string, error) {
	rtn := ""
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s/connection", c.coreUri, id)
	res, err := c.http.Get(requestURL, http.StatusOK)
	if err != nil {
		return rtn, errors.Wrap(err, "error doing http get")

	}

	return string(res), err
}
