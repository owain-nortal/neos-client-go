package neos

import (
	"context"
	"fmt"
	"net/http"
	"github.com/pkg/errors"
)

type DataUnitClient struct {
	coreUri string
	http    *NeosHttp
}

func NewDataUnitClient(coreUri string, http *NeosHttp) *DataUnitClient {
	return &DataUnitClient{
		coreUri: coreUri,
		http:    http,
	}
}

func (c *DataUnitClient) Delete(ctx context.Context, id string) error {
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s", c.coreUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *DataUnitClient) Post(ctx context.Context, dspr DataUnitPostRequest) (DataUnitPostResponse, error) {
	var rtn DataUnitPostResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit", c.coreUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataUnitClient) Put(ctx context.Context, id string, dspr DataUnitPutRequest) (DataUnitPutResponse, error) {
	var rtn DataUnitPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataUnitClient) ConfigTablePut(ctx context.Context, id string, config DataUnitConfigurationTablePutRequest) (DataUnitConfigurationTablePutResponse, error) {
	var rtn DataUnitConfigurationTablePutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s/config", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, config, http.StatusOK, &rtn)
	return rtn, err
}
func (c *DataUnitClient) ConfigDataProductPut(ctx context.Context, id string, config DataUnitConfigurationDataProductPutRequest) (DataUnitConfigurationDataProductPutResponse, error) {
	var rtn DataUnitConfigurationDataProductPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s/config", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, config, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataUnitClient) ConfigQueryPut(ctx context.Context, id string, config DataUnitConfigurationQueryPutRequest) (DataUnitConfigurationQueryPutResponse, error) {
	var rtn DataUnitConfigurationQueryPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s/config", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, config, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataUnitClient) ConfigCSVPut(ctx context.Context, id string, config DataUnitConfigurationCSVPutRequest) (DataUnitConfigurationCSVPutResponse, error) {
	var rtn DataUnitConfigurationCSVPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s/config", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, config, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataUnitClient) ConfigParquetPut(ctx context.Context, id string, config DataUnitConfigurationParquetPutRequest) (DataUnitConfigurationParquetPutResponse, error) {
	var rtn DataUnitConfigurationParquetPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s/config", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, config, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataUnitClient) PutInfo(ctx context.Context, id string, dspr DataUnitPutRequestEntityInfo) (DataUnitPutInfoResponse, error) {
	var rtn DataUnitPutInfoResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s/info", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataUnitClient) Get() (DataUnitList, error) {
	var rtn DataUnitList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit", c.coreUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *DataUnitClient) ConfigTableGet(ctx context.Context, id string) (DataUnitConfigurationTableGetResponse, error) {
	var rtn DataUnitConfigurationTableGetResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s/config", c.coreUri, id)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}


