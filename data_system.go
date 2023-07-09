package neos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

type DataSystemClient struct {
	url string
}

func NewDataSystemClientV2(url string) DataSystemClient {
	return DataSystemClient{url}
}

func boolToString(input bool) string {
	if input {
		return "true"
	}
	return "false"
}

func createHttpRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", GetAccessToken()))
	return req, err
}

func (c DataSystemClient) Delete(id string) error {

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_system/%s", c.url, id)
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

func (c DataSystemClient) Post(dspr DataSystemPostRequest) (DataSystemPostResponse, error) {
	var rtn DataSystemPostResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_system", c.url)
	req, err := createHttpRequest(http.MethodPost, requestURL, bytes.NewBuffer(b))
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

	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d", res.StatusCode)
	}
	err = json.Unmarshal([]byte(resBody), &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, nil
}

func (c DataSystemClient) Put(id string, dspr DataSystemPutRequest) (DataSystemPutResponse, error) {
	var rtn DataSystemPutResponse

	b, err := json.Marshal(dspr)
	if err != nil {
		return rtn, errors.Wrap(err, " could not marshal request")
	}

	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_system/%s", c.url, id)
	req, err := createHttpRequest(http.MethodPut, requestURL, bytes.NewBuffer(b))
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

	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d", res.StatusCode)
	}
	err = json.Unmarshal([]byte(resBody), &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, nil
}

func (c DataSystemClient) Get(hasChild bool) (DataSystemList, error) {

	var rtn DataSystemList
	requestURL := fmt.Sprintf("%s/api/gateway/v2/data_system?has_child=%s", c.url, boolToString(hasChild))
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

	if res.StatusCode != http.StatusOK {
		return rtn, fmt.Errorf(" unexpected response code %d", res.StatusCode)
	}
	err = json.Unmarshal([]byte(resBody), &rtn)
	if err != nil {
		return rtn, errors.Wrap(err, " could not unmashal body")
	}

	return rtn, nil
}
