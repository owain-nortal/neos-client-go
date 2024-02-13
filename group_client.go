package neos

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type GroupClient struct {
	hubUri string
	http   *NeosHttp
}

func NewGroupClient(hubUri string, http *NeosHttp) *GroupClient {
	return &GroupClient{
		hubUri: hubUri,
		http:   http,
	}
}



func (c *GroupClient) List(filter string) (GroupList, error) {
	var rtn GroupList
	requestURL := fmt.Sprintf("%s/api/hub/iam/group%s", c.hubUri, filterQuery(filter, "account"))
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) Get(id string) (Group, error) {
	var rtn Group
	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s", c.hubUri, id)
	err := c.http.GetUnmarshal(requestURL, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) Post(ctx context.Context, dspr GroupPostRequest) (Group, error) {
	var rtn Group
	requestURL := fmt.Sprintf("%s/api/hub/iam/group", c.hubUri)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) Put(ctx context.Context, id string, dspr GroupPutRequest) (Group, error) {
	var rtn Group
	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s", c.hubUri, id)
	err := c.http.PutUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) Delete(ctx context.Context, id string) error {
	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s", c.hubUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}

func (c *GroupClient) PrincipalsPost(ctx context.Context, id string, dspr GroupPrincipalPostRequest) (Group, error) {
	var rtn Group
	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s/principals", c.hubUri, id)
	err := c.http.PostUnmarshal(requestURL, dspr, http.StatusOK, &rtn)
	return rtn, err
}

func (c *GroupClient) PrincipalsDelete(ctx context.Context, id string) error {
	requestURL := fmt.Sprintf("%s/api/hub/iam/group/%s/principals", c.hubUri, id)
	err := c.http.Delete(requestURL, http.StatusOK)
	if err != nil {
		return errors.Wrap(err, "error doing http delete")
	}
	return err
}
