package neos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/errors"
)

func (c *NeosClient) OutputDelete(ctx context.Context, id string) error {

	requestURL := fmt.Sprintf("%s/api/gateway/v2/output/%s", c.coreUri, id)
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

func (c *NeosClient) OutputPost(ctx context.Context, dspr OutputPostRequest) (OutputPostResponse, error) {

	tflog.Info(ctx, fmt.Sprintf("£##> Client Post request [%s] [%s] [%s] ", dspr.Entity.Label, dspr.Entity.Name, dspr.Entity.Description))

	var rtn OutputPostResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)

	os.WriteFile("/tmp/output_post_json", []byte(b), 0644)
	os.WriteFile("/tmp/output_post-json_unq", []byte(unquotedString), 0644)

	tflog.Info(ctx, fmt.Sprintf("£##> Client Post json [%s] ", unquotedString))

	requestURL := fmt.Sprintf("%s/api/gateway/v2/output", c.coreUri)
	req, err := createHttpRequest(http.MethodPost, requestURL, bytes.NewBuffer([]byte(unquotedString)))

	//	tflog.Info(ctx, fmt.Sprintf("Method %s", req.Method))
	if err != nil {
		return rtn, errors.Wrap(err, " could not create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return rtn, errors.Wrap(err, " error making http request ")
	}

	resBody, err := ioutil.ReadAll(res.Body)
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

func (c *NeosClient) OutputPut(ctx context.Context, id string, dspr OutputPutRequest) (OutputPutResponse, error) {
	var rtn OutputPutResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	requestURL := fmt.Sprintf("%s/api/gateway/v2/output/%s", c.coreUri, id)

	// os.WriteFile("/tmp/put-id", []byte(id), 0644)
	// os.WriteFile("/tmp/put-json", []byte(b), 0644)

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)

	req, err := createHttpRequest(http.MethodPut, requestURL, bytes.NewBuffer([]byte(unquotedString)))
	if err != nil {
		return rtn, errors.Wrap(err, " could not create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return rtn, errors.Wrap(err, " error making http request ")
	}

	resBody, err := ioutil.ReadAll(res.Body)
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

func (c *NeosClient) OutputPutInfo(ctx context.Context, id string, dspr OutputPutRequestEntityInfo) (OutputPutInfoResponse, error) {
	var rtn OutputPutInfoResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	requestURL := fmt.Sprintf("%s/api/gateway/v2/output/%s/info", c.coreUri, id)

	// os.WriteFile("/tmp/put-id", []byte(id), 0644)
	// os.WriteFile("/tmp/put-json", []byte(b), 0644)

	unquotedString := strings.Replace(string(b), "\\\"", "", -1)

	req, err := createHttpRequest(http.MethodPut, requestURL, bytes.NewBuffer([]byte(unquotedString)))
	if err != nil {
		return rtn, errors.Wrap(err, " could not create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return rtn, errors.Wrap(err, " error making http request ")
	}

	resBody, err := ioutil.ReadAll(res.Body)
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

func (c *NeosClient) OutputGet() (OutputList, error) {

	var rtn OutputList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/output", c.coreUri)
	req, err := createHttpRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return rtn, errors.Wrap(err, " could not create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return rtn, errors.Wrap(err, " error making http request ")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return rtn, errors.Wrap(err, " could not read response body")
	}

	byteBody := []byte(resBody)

	os.WriteFile("/tmp/get-body", byteBody, 0644)

	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d", res.StatusCode)
	}
	err = json.Unmarshal([]byte(resBody), &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, nil
}
