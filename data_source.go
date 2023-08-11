package neos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/errors"
)

func (c *NeosClient) DataSourceDelete(ctx context.Context, id string) error {

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s", c.coreUri, id)
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

func (c *NeosClient) DataSourcePost(ctx context.Context, dspr DataSourcePostRequest) (DataSourcePostResponse, error) {

	tflog.Info(ctx, fmt.Sprintf("DataSourcePost request [%s] [%s] [%s] ", dspr.Entity.Label, dspr.Entity.Name, dspr.Entity.Description))

	var rtn DataSourcePostResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)

	//	tflog.Info(ctx, fmt.Sprintf("£##> Client Post json [%s] ", unquotedString))

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source", c.coreUri)
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

func (c *NeosClient) DataSourcePut(ctx context.Context, id string, dspr DataSourcePutRequest) (DataSourcePutResponse, error) {
	var rtn DataSourcePutResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s", c.coreUri, id)

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

func (c *NeosClient) DataSourcePutInfo(ctx context.Context, id string, dspr DataSourcePutRequestEntityInfo) (DataSourcePutInfoResponse, error) {
	var rtn DataSourcePutInfoResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source/%s/info", c.coreUri, id)
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

func (c *NeosClient) DataSourceGet() (DataSourceList, error) {

	var rtn DataSourceList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_source", c.coreUri)
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
