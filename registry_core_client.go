package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type RegistryCoreClient struct {
	registryUri string
	http        *NeosHttp
	Account     string
}

func NewRegistryCoreClient(registryUri string, http *NeosHttp, account string) *RegistryCoreClient {
	return &RegistryCoreClient{
		registryUri: registryUri,
		http:        http,
		Account:     account,
	}
}

func (c *RegistryCoreClient) Delete(ctx context.Context, id string) error {
	requestURL := fmt.Sprintf("%s/api/registry/core/%s", c.registryUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *RegistryCoreClient) Post(ctx context.Context, dspr RegistryCorePostRequest) (RegistryCorePostResponse, error) {
	var rtn RegistryCorePostResponse
	requestURL := fmt.Sprintf("%s/api/hub/registry/core", c.registryUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *RegistryCoreClient) Get() (RegistryCoreList, error) {
	var rtn RegistryCoreList
	requestURL := fmt.Sprintf("%s/api/hub/registry/core", c.registryUri)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}
