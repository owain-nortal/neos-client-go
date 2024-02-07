package neos

import (
	"bytes"
	"context"
	//"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"strings"
)

type DataProductClient struct {
	coreUri string
	http    *NeosHttp
}

func NewDataProductClient(coreUri string, http *NeosHttp) *DataProductClient {
	return &DataProductClient{
		coreUri: coreUri,
		http:    http,
	}
}

func (c DataProductClient) Delete(ctx context.Context, id string) error {
	tflog.Info(ctx, fmt.Sprintf("DataProductDelete %s", id))
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product/%s", c.coreUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	return err
}

func (c DataProductClient) Post(ctx context.Context, dspr DataProductPostRequest) (DataProductPostResponse, error) {
	tflog.Info(ctx, fmt.Sprintf("Client Post request [%s] [%s] [%s] ", dspr.Entity.Label, dspr.Entity.Name, dspr.Entity.Description))
	var rtn DataProductPostResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product", c.coreUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err

}

func (c DataProductClient) Put(ctx context.Context, id string, dspr DataProductPutRequest) (DataProductPutResponse, error) {
	var rtn DataProductPutResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product/%s", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err

}

func (c DataProductClient) DataProductPutInfo(ctx context.Context, id string, dspr DataProductPutRequestEntityInfo) (DataProductPutInfoResponse, error) {
	var rtn DataProductPutInfoResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product/%s/info", c.coreUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c DataProductClient) Get() (DataProductList, error) {
	var rtn DataProductList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product", c.coreUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

// func (c DataProductClient) SchemaGet(id string) (DataProductSchema, error) {

// 	var rtn DataProductSchema
// 	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product/%s/schema", c.coreUri, id)
// 	resBody, err := c.http.Get(requestURL, http.StatusOK)
// 	if err != nil {
// 		return rtn, errors.Wrap(err, " get failed ")
// 	}
// 	err = json.Unmarshal([]byte(resBody), &rtn)
// 	if err != nil {
// 		return rtn, errors.Wrap(err, " could not unmashal body")
// 	}

// 	return rtn, nil

// }

// func (c DataProductClient) DataProductSchemaPut(ctx context.Context, id string, dspr DataProductSchemaPutRequest) (DataProductSchemaPutResponse, error) {
// 	tflog.Info(ctx, fmt.Sprintf("DataProductSchemaPut %s", id))

// 	var rtn DataProductSchemaPutResponse

// 	b, err := json.Marshal(dspr)
// 	if err != nil {
// 		return rtn, errors.Wrap(err, " could not marshal request")
// 	}

// 	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product/%s/schema", c.coreUri, id)

// 	unquotedString := strings.Replace(string(b), "\\\"", "", -1)

// 	tflog.Info(ctx, fmt.Sprintf("DataProductSchemaPut requestURL %s", requestURL))
// 	tflog.Info(ctx, fmt.Sprintf("DataProductSchemaPut request body %s", unquotedString))
// 	req, err := createHttpRequest(http.MethodPut, requestURL, bytes.NewBuffer([]byte(unquotedString)))
// 	if err != nil {
// 		return rtn, errors.Wrap(err, " could not create request")
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return rtn, errors.Wrap(err, " error making http request ")
// 	}

// 	resBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return rtn, errors.Wrap(err, " could not read response body")
// 	}

// 	byteBody := []byte(resBody)

// 	tflog.Info(ctx, fmt.Sprintf("DataProductSchemaPut result body %s", string(byteBody)))

// 	if res.StatusCode != http.StatusOK {
// 		return rtn, fmt.Errorf(" unexpected response code %d %s", res.StatusCode, byteBody)
// 	}
// 	err = json.Unmarshal(byteBody, &rtn)
// 	if err != nil {
// 		return rtn, errors.Wrap(err, " could not unmashal body")
// 	}

// 	return rtn, nil
// }

func (c DataProductClient) DataProductBuilderPut(ctx context.Context, id string, json string) (DataProductBuilderPutResponse, error) {
	tflog.Info(ctx, fmt.Sprintf("DataProductBuilderPut %s", id))

	var rtn DataProductBuilderPutResponse

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_product/%s/spark/builder", c.coreUri, id)

	unquotedString := strings.Replace(json, "\\\"", "", -1)
	noNL := strings.Replace(unquotedString, "\\n", "", -1)

	d1 := []byte(noNL)
	_ = os.WriteFile("/tmp/dat1", d1, 0644)
	_ = os.WriteFile("/tmp/dat2", []byte(json), 0644)

	tflog.Info(ctx, fmt.Sprintf("DataProductBuilderPut requestURL %s", requestURL))
	tflog.Info(ctx, fmt.Sprintf("DataProductBuilderPut request body %s", noNL))
	req, err := createHttpRequest(http.MethodPut, requestURL, bytes.NewBuffer([]byte(noNL)))
	if err != nil {
		return rtn, errors.Wrap(err, " could not create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return rtn, errors.Wrap(err, " error making http request ")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return rtn, errors.Wrap(err, " could not read response body")
	}

	byteBody := []byte(resBody)

	tflog.Info(ctx, fmt.Sprintf("DataProductBuilderPut result body %s", string(byteBody)))

	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d %s", res.StatusCode, byteBody)
	}

	return rtn, nil
}
