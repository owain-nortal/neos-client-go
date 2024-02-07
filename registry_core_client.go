package neos

import (
	"context"
	"fmt"
	"net/http"
	"github.com/pkg/errors"
)

type RegistryCoreClient struct {
	registryUri string
	http        *NeosHttp
}

func NewRegistryCoreClient(registryUri string, http *NeosHttp) *RegistryCoreClient {
	return &RegistryCoreClient{
		registryUri: registryUri,
		http:        http,
	}
}

func (c *RegistryCoreClient) Delete(ctx context.Context, rcdr RegistryCoreDeleteRequest) error {
	requestURL := fmt.Sprintf("%s/api/registry/core", c.registryUri)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *RegistryCoreClient) Post(ctx context.Context, dspr RegistryCorePostRequest) (RegistryCorePostResponse, error) {
	var rtn RegistryCorePostResponse
	requestURL := fmt.Sprintf("%s/api/registry/core", c.registryUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *RegistryCoreClient) RegistryCoreGet() (RegistryCoreList, error) {
	var rtn RegistryCoreList
	requestURL := fmt.Sprintf("%s/api/registry/core", c.registryUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}
