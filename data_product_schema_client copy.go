package neos

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
)

type DataProductSchemaClient struct {
	coreUri string
	http    *NeosHttp
	Account string
}

func NewDataProductSchemaClient(coreUri string, http *NeosHttp, account string) *DataProductSchemaClient {
	return &DataProductSchemaClient{
		coreUri: coreUri,
		http:    http,
		Account: account,
	}
}

func (c DataProductSchemaClient) Get(id string) (DataProductSchema, error) {
	var rtn DataProductSchema
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product/%s/schema", c.coreUri, id)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c DataProductSchemaClient) Put(ctx context.Context, id string, dspr DataProductSchemaPutRequest) (DataProductSchemaPutResponse, error) {
	tflog.Info(ctx, fmt.Sprintf("DataProductSchemaPut %s", id))
	var rtn DataProductSchemaPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product/%s/schema", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}
