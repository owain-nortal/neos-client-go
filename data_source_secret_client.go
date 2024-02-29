package neos

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type DataSourceSecretClient struct {
	coreUri string
	http    *NeosHttp
	Account	string
}

func NewDataSourceSecretClient(coreUri string, http *NeosHttp, account string) *DataSourceSecretClient {
	return &DataSourceSecretClient{
		coreUri: coreUri,
		http:    http,
		Account: account,
	}
}

func (c *DataSourceSecretClient) Post(ctx context.Context, id string, data map[string]string) (string, error) {
	rtn := ""

	blob, err := json.Marshal(data)
	if err != nil {
		return rtn, errors.Wrap(err, "failed to marshal data")
	}
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s/secret", c.coreUri, id)
	res, err := c.http.PostRaw(requestURL, string(blob), http.StatusOK)

	if err != nil {
		return rtn, errors.Wrap(err, "error doing http Post")

	}
	return string(res), err
}

