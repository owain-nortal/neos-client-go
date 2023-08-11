package neos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"io"
	"net/http"
	"os"
	//"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/errors"
)

func (c *NeosClient) LinkDelete(ctx context.Context, source string, dest string, parentId string, childId string) error {

	requestURL := fmt.Sprintf("%s/api/gateway/v2/link/%s/%s/%s/%s", c.coreUri, source, parentId, dest, childId)
	tflog.Info(ctx, fmt.Sprintf("linkDelete uri: %s", requestURL))

	req, err := createHttpRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		return errors.Wrap(err, " could not create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, " error making http request ")
	}

	if res.StatusCode == http.StatusNotFound {
		return nil
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(" unexpected response code %d", res.StatusCode)
	}

	return nil
}

func (c *NeosClient) linkPost(ctx context.Context, source string, dest string, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	tflog.Info(ctx, fmt.Sprintf("LinkDataSourceDataUnitPost request parent: [%s] child: [%s]", parentIdentifier, childIdentifier))
	requestURL := fmt.Sprintf("%s/api/gateway/v2/link/%s/%s/%s/%s", c.coreUri, source, parentIdentifier, dest, childIdentifier)

	var rtn LinkPostResponse
	tflog.Info(ctx, fmt.Sprintf("linkPost uri: %s", requestURL))

	req, err := createHttpRequest(http.MethodPost, requestURL, bytes.NewBuffer([]byte{}))

	if err != nil {
		return rtn, errors.Wrap(err, "could not create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return rtn, errors.Wrap(err, "error making http request ")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return rtn, errors.Wrap(err, "could not read response body")
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

func (c *NeosClient) LinkDataSourceDataUnitPost(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.linkPost(ctx, "data_source", "data_unit", parentIdentifier, childIdentifier)
}

func (c *NeosClient) LinkDataSourceDataUnitDelete(ctx context.Context, parentId string, childId string) error {
	return c.LinkDelete(ctx, "data_source", "data_unit", parentId, childId)
}

func (c *NeosClient) LinkDataProductOutputPost(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.linkPost(ctx, "data_product", "output", parentIdentifier, childIdentifier)
}

func (c *NeosClient) LinkDataProductOutputDelete(ctx context.Context, parentId string, childId string) error {
	return c.LinkDelete(ctx, "data_product", "output", parentId, childId)
}

func (c *NeosClient) LinkDataProductDataProductPost(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.linkPost(ctx, "data_product", "data_product", parentIdentifier, childIdentifier)
}

func (c *NeosClient) LinkDataProductDataProductDelete(ctx context.Context, parentId string, childId string) error {
	return c.LinkDelete(ctx, "data_product", "data_product", parentId, childId)
}

func (c *NeosClient) LinkDataUnitDataProductPost(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.linkPost(ctx, "data_unit", "data_product", parentIdentifier, childIdentifier)
}

func (c *NeosClient) LinkDataUnitDataProductDelete(ctx context.Context, parentId string, childId string) error {
	return c.LinkDelete(ctx, "data_unit", "data_product", parentId, childId)
}

func (c *NeosClient) LinkDataSystemDataSourcePost(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.linkPost(ctx, "data_system", "data_source", parentIdentifier, childIdentifier)
}

func (c *NeosClient) LinkDataSystemDataSourceDelete(ctx context.Context, parentId string, childId string) error {
	return c.LinkDelete(ctx, "data_system", "data_source", parentId, childId)
}

func (c *NeosClient) LinksGet() (LinksGetResponse, error) {

	var rtn LinksGetResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/link", c.coreUri)
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
