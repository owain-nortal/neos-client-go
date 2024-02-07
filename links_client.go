package neos

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/errors"
	"net/http"
)

type LinksClient struct {
	coreUri string
	http    *NeosHttp
}

func NewLinksClient(coreUri string, http *NeosHttp) *LinksClient {
	return &LinksClient{
		coreUri: coreUri,
		http:    http,
	}
}

func (c *LinksClient) Delete(ctx context.Context, source string, dest string, parentId string, childId string) error {
	requestURL := fmt.Sprintf("%s/api/gateway/v2/link/%s/%s/%s/%s", c.coreUri, source, parentId, dest, childId)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *LinksClient) Post(ctx context.Context, source string, dest string, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	tflog.Info(ctx, fmt.Sprintf("LinkDataSourceDataUnitPost request parent: [%s] child: [%s]", parentIdentifier, childIdentifier))
	requestURL := fmt.Sprintf("%s/api/gateway/v2/link/%s/%s/%s/%s", c.coreUri, source, parentIdentifier, dest, childIdentifier)	
	var rtn LinkPostResponse	
	err := c.http.PostUnmarshal(requestURL, []byte{}, http.StatusOK, &rtn)
	return rtn, err
}

func (c *LinksClient) LinkDataSourceToDataUnit(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.Post(ctx, "data_source", "data_unit", parentIdentifier, childIdentifier)
}

func (c *LinksClient) DeleteLinkDataSourceToDataUnit(ctx context.Context, parentId string, childId string) error {
	return c.Delete(ctx, "data_source", "data_unit", parentId, childId)
}

func (c *LinksClient) LinkDataProductToOutput(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.Post(ctx, "data_product", "output", parentIdentifier, childIdentifier)
}

func (c *LinksClient) DeleteLinkDataProductToOutput(ctx context.Context, parentId string, childId string) error {
	return c.Delete(ctx, "data_product", "output", parentId, childId)
}

func (c *LinksClient) LinkDataProductToDataProduct(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.Post(ctx, "data_product", "data_product", parentIdentifier, childIdentifier)
}

func (c *LinksClient) DeleteLinkDataProductToDataProduct(ctx context.Context, parentId string, childId string) error {
	return c.Delete(ctx, "data_product", "data_product", parentId, childId)
}

func (c *LinksClient) LinkDataUnitToDataProduct(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.Post(ctx, "data_unit", "data_product", parentIdentifier, childIdentifier)
}

func (c *LinksClient) DeleteLinkDataUnitToDataProduct(ctx context.Context, parentId string, childId string) error {
	return c.Delete(ctx, "data_unit", "data_product", parentId, childId)
}

func (c *LinksClient) LinkDataSystemToDataSource(ctx context.Context, parentIdentifier string, childIdentifier string) (LinkPostResponse, error) {
	return c.Post(ctx, "data_system", "data_source", parentIdentifier, childIdentifier)
}

func (c *LinksClient) DeleteLinkDataSystemToDataSource(ctx context.Context, parentId string, childId string) error {
	return c.Delete(ctx, "data_system", "data_source", parentId, childId)
}

func (c *LinksClient) Get() (LinksGetResponse, error) {
	var rtn LinksGetResponse
	requestURL := fmt.Sprintf("%s/api/gateway/v2/link", c.coreUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}
