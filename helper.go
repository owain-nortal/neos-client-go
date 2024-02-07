package neos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

func boolToString(input bool) string {
	if input {
		return "true"
	}
	return "false"
}

func createHttpRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", GetAccessToken()))
	return req, err
}

type NeosHttp struct {
}

func NewNeosHttp() *NeosHttp {
	return &NeosHttp{}
}

func (n *NeosHttp) CreateHttpRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", GetAccessToken()))
	return req, err
}

func (n *NeosHttp) Delete(requestURL string, expectedCode int) error {
	req, err := n.CreateHttpRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		return errors.Wrap(err, " could not create delete request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, " error making http delete request ")
	}

	if res.StatusCode != expectedCode {
		return fmt.Errorf("delete unexpected response code %d", res.StatusCode)
	}
	return nil
}

func (n *NeosHttp) Post(requestURL string, unquotedString string, expectedCode int) ([]byte, error) {
	req, err := n.CreateHttpRequest(http.MethodPost, requestURL, bytes.NewBuffer([]byte(unquotedString)))

	if err != nil {
		return nil, errors.Wrap(err, "could not create post request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error making http post request ")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, " could not read post response body")
	}

	byteBody := []byte(resBody)
	if res.StatusCode != expectedCode {
		return nil, fmt.Errorf(" post unexpected response code %d %s", res.StatusCode, byteBody)
	}
	return byteBody, nil
}

func (n *NeosHttp) PostUnmarshal(requestURL string, input any, expectedCode int, output any) error {
	b, err := json.Marshal(input)
	if err != nil {
		return errors.Wrap(err, " could not marshal request")
	}

	byteBody, err := n.Post(requestURL, unquoteString(b), expectedCode)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteBody, &output)
	if err != nil {
		return errors.Wrap(err, " could not unmashal body")
	}
	return nil
}

func (n *NeosHttp) Put(requestURL string, unquotedString string, expectedCode int) ([]byte, error) {
	req, err := n.CreateHttpRequest(http.MethodPost, requestURL, bytes.NewBuffer([]byte(unquotedString)))

	if err != nil {
		return nil, errors.Wrap(err, "could not create put request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error making http put request ")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, " could not read put response body")
	}

	byteBody := []byte(resBody)
	if res.StatusCode != expectedCode {
		return nil, fmt.Errorf(" put unexpected response code %d %s", res.StatusCode, byteBody)
	}
	return byteBody, nil
}

func (n *NeosHttp) PutUnmarshal(requestURL string, input any, expectedCode int, output any) error {

	b, err := json.Marshal(input)
	if err != nil {
		return errors.Wrap(err, " could not marshal request")
	}

	byteBody, err := n.Put(requestURL, unquoteString(b), expectedCode)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteBody, &output)
	if err != nil {
		return errors.Wrap(err, " could not unmashal body")
	}

	return nil
}

func (n *NeosHttp) Get(requestURL string, expectedCode int) ([]byte, error) {
	req, err := n.CreateHttpRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, " could not create get request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, " error making http get request ")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, " could not read get response body")
	}

	if res.StatusCode != expectedCode {
		return nil, fmt.Errorf("get unexpected response code %d", res.StatusCode)
	}
	return resBody, nil
}

func (n *NeosHttp) GetUnmarshal(requestURL string, expectedCode int, output any) error {
	byteBody, err := n.Get(requestURL, expectedCode)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteBody, &output)
	if err != nil {
		return errors.Wrap(err, " could not unmashal body")
	}
	return nil
}

func unquoteString(b []byte) string {
	unquotedString := strings.Replace(string(b), "\\\"", "", -1)
	return unquotedString
}
