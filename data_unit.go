package neos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/errors"
)

func (c *NeosClient) DataUnitDelete(ctx context.Context, id string) error {

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s", c.coreUri, id)
	req, err := createHttpRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		return errors.Wrap(err, " could not create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, " error making http request ")
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(" unexpected response code %d", res.StatusCode)
	}

	return nil
}

func (c *NeosClient) DataUnitPost(ctx context.Context, dspr DataUnitPostRequest) (DataUnitPostResponse, error) {

	tflog.Info(ctx, fmt.Sprintf("DataUnitPost request [%s] [%s] [%s] ", dspr.Entity.Label, dspr.Entity.Name, dspr.Entity.Description))

	var rtn DataUnitPostResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)

	//	tflog.Info(ctx, fmt.Sprintf("£##> Client Post json [%s] ", unquotedString))

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit", c.coreUri)
	req, err := createHttpRequest(http.MethodPost, requestURL, bytes.NewBuffer([]byte(unquotedString)))

	//	tflog.Info(ctx, fmt.Sprintf("Method %s", req.Method))
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
	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d %s", res.StatusCode, byteBody)
	}

	//	tflog.Info(ctx, fmt.Sprintf("£##> Client Post result [%s] ", string(byteBody)))

	err = json.Unmarshal(byteBody, &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, nil
}

func (c *NeosClient) DataUnitPut(ctx context.Context, id string, dspr DataUnitPutRequest) (DataUnitPutResponse, error) {
	var rtn DataUnitPutResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s", c.coreUri, id)
	unquotedString := strings.Replace(string(b), "\\\"", "", -1)

	req, err := createHttpRequest(http.MethodPut, requestURL, bytes.NewBuffer([]byte(unquotedString)))
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
	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d %s", res.StatusCode, byteBody)
	}
	err = json.Unmarshal(byteBody, &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, nil
}

func (c *NeosClient) DataUnitConfigTablePut(ctx context.Context, id string, config DataUnitConfigurationTablePutRequest) (DataUnitConfigurationTablePutResponse, error) {

	var rtn DataUnitConfigurationTablePutResponse
	b, err := json.Marshal(config)
	if err != nil {
		return rtn, errors.Wrap(err, "could not marshal request")
	}

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)
	byteBody, err := c.dataUnitConfigPutBase(ctx, id, []byte(unquotedString))
	if err != nil {
		return rtn, err
	}
	err = json.Unmarshal(byteBody, &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, err
}
func (c *NeosClient) DataUnitConfigDataProductPut(ctx context.Context, id string, config DataUnitConfigurationDataProductPutRequest) (DataUnitConfigurationDataProductPutResponse, error) {

	var rtn DataUnitConfigurationDataProductPutResponse
	b, err := json.Marshal(config)
	if err != nil {
		return rtn, errors.Wrap(err, "could not marshal request")
	}

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)
	byteBody, err := c.dataUnitConfigPutBase(ctx, id, []byte(unquotedString))
	if err != nil {
		return rtn, err
	}
	err = json.Unmarshal(byteBody, &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, err
}

func (c *NeosClient) DataUnitConfigQueryPut(ctx context.Context, id string, config DataUnitConfigurationQueryPutRequest) (DataUnitConfigurationQueryPutResponse, error) {

	var rtn DataUnitConfigurationQueryPutResponse
	b, err := json.Marshal(config)
	if err != nil {
		return rtn, errors.Wrap(err, "could not marshal request")
	}

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)
	byteBody, err := c.dataUnitConfigPutBase(ctx, id, []byte(unquotedString))
	if err != nil {
		return rtn, err
	}
	err = json.Unmarshal(byteBody, &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, err
}



func (c *NeosClient) DataUnitConfigCSVPut(ctx context.Context, id string, config DataUnitConfigurationCSVPutRequest) (DataUnitConfigurationCSVPutResponse, error) {

	var rtn DataUnitConfigurationCSVPutResponse
	b, err := json.Marshal(config)
	if err != nil {
		return rtn, errors.Wrap(err, "could not marshal request")
	}

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)
	byteBody, err := c.dataUnitConfigPutBase(ctx, id, []byte(unquotedString))
	if err != nil {
		return rtn, err
	}
	err = json.Unmarshal(byteBody, &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, err
}

func (c *NeosClient) DataUnitConfigParquetPut(ctx context.Context, id string, config DataUnitConfigurationParquetPutRequest) (DataUnitConfigurationParquetPutResponse, error) {

	var rtn DataUnitConfigurationParquetPutResponse
	b, err := json.Marshal(config)
	if err != nil {
		return rtn, errors.Wrap(err, "could not marshal request")
	}

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)
	byteBody, err := c.dataUnitConfigPutBase(ctx, id, []byte(unquotedString))
	if err != nil {
		return rtn, err
	}
	err = json.Unmarshal(byteBody, &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, err
}

func (c *NeosClient) dataUnitConfigPutBase(ctx context.Context, id string, payload []byte) ([]byte, error) {
	var rtn []byte

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s/config", c.coreUri, id)

	tflog.Info(ctx, fmt.Sprintf("Payload [%s]", string(payload)))


	req, err := createHttpRequest(http.MethodPut, requestURL, bytes.NewBuffer(payload))
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

	rtn = []byte(resBody)
	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d %s", res.StatusCode, rtn)
	}

	return rtn, nil
}

func (c *NeosClient) DataUnitPutInfo(ctx context.Context, id string, dspr DataUnitPutRequestEntityInfo) (DataUnitPutInfoResponse, error) {
	var rtn DataUnitPutInfoResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit/%s/info", c.coreUri, id)
	unquotedString := strings.Replace(string(b), "\\\"", "", -1)

	req, err := createHttpRequest(http.MethodPut, requestURL, bytes.NewBuffer([]byte(unquotedString)))
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
	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d %s", res.StatusCode, byteBody)
	}
	err = json.Unmarshal(byteBody, &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, nil
}

func (c *NeosClient) DataUnitGet() (DataUnitList, error) {

	var rtn DataUnitList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_unit", c.coreUri)
	req, err := createHttpRequest(http.MethodGet, requestURL, nil)
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

	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d", res.StatusCode)
	}
	err = json.Unmarshal([]byte(resBody), &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, nil
}
