package neos

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	//"github.com/owain-nortal/neos-client-go"
)

func createServers() {

}

func TestDataSystemV2PostOK(t *testing.T) {

	expected := DataSystemPostResponse{}
	expected.Identifier = "xyz321"
	expected.Name = "something"
	b, err := json.Marshal(expected)
	if err != nil {
		fmt.Println(err)
		return
	}
	expectedJson := string(b)
	iamsvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	registrysvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	coresvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedJson)
	}))

	defer iamsvr.Close()
	defer registrysvr.Close()
	defer coresvr.Close()

	request := DataSystemPostRequest{
		Entity: DataSystemPostRequestEntity{
			Name: expected.Name,
		},
		EntityInfo: DataSystemPostRequestEntityInfo{
			Owner: "Mr Owner",
		},
	}

	c := NewNeosClient(iamsvr.URL, registrysvr.URL, coresvr.URL)
	res, err := c.DataSystemPost(context.Background(), request)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if res.Name != request.Entity.Name {
		t.Errorf("expected entiry name to be %s got %s", request.Entity.Name, res.Name)
	}
}

func TestDataSystemV2PutOK(t *testing.T) {

	expected := DataSystemPutResponse{}
	expected.Identifier = "xyz321"
	expected.Name = "something"
	b, err := json.Marshal(expected)
	if err != nil {
		fmt.Println(err)
		return
	}
	expectedJson := string(b)
	coresvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedJson)
	}))
	iamsvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	registrysvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	defer iamsvr.Close()
	defer registrysvr.Close()
	defer coresvr.Close()

	request := DataSystemPutRequest{
		Entity: DataSystemPutRequestEntity{
			Name: expected.Name,
		},
	}

	c := NewNeosClient(iamsvr.URL, registrysvr.URL, coresvr.URL)
	res, err := c.DataSystemPut(context.Background(), "123", request)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if res.Name != request.Entity.Name {
		t.Errorf("expected entiry name to be %s got %s", request.Entity.Name, res.Name)
	}
}

func TestDataSystemV2PutFailed(t *testing.T) {

	coresvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	iamsvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	registrysvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer iamsvr.Close()
	defer registrysvr.Close()
	defer coresvr.Close()
	request := DataSystemPutRequest{}
	c := NewNeosClient(iamsvr.URL, registrysvr.URL, coresvr.URL)
	_, err := c.DataSystemPut(context.Background(), "321ads", request)
	if err == nil {
		t.Errorf("expected err to be set not nil")
	}
}

func TestDataSystemV2PostFailed(t *testing.T) {

	coresvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	iamsvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	registrysvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer iamsvr.Close()
	defer registrysvr.Close()
	defer coresvr.Close()
	request := DataSystemPostRequest{}
	c := NewNeosClient(iamsvr.URL, registrysvr.URL, coresvr.URL)
	_, err := c.DataSystemPost(context.Background(), request)
	if err == nil {
		t.Errorf("expected err to be set not nil")
	}
}

func TestDataSystemV2GetOK(t *testing.T) {

	expected := DataSystemList{}
	ds := []DataSystem{
		{
			Identifier: "abc123",
		},
	}

	expected.Entities = ds

	b, err := json.Marshal(expected)
	if err != nil {
		fmt.Println(err)
		return
	}
	expectedJson := string(b)
	coresvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedJson)
	}))
	iamsvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	registrysvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer iamsvr.Close()
	defer registrysvr.Close()
	defer coresvr.Close()
	c := NewNeosClient(iamsvr.URL, registrysvr.URL, coresvr.URL)
	res, err := c.DataSystemGet()
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if res.Entities[0].Identifier != expected.Entities[0].Identifier {
		t.Errorf("expected res to be %s got %s", expected.Entities[0].Identifier, res.Entities[0].Identifier)
	}
}

func TestDataSystemV2DeleteOK(t *testing.T) {
	coresvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	iamsvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	registrysvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer iamsvr.Close()
	defer registrysvr.Close()
	defer coresvr.Close()
	c := NewNeosClient(iamsvr.URL, registrysvr.URL, coresvr.URL)
	err := c.DataSystemDelete(context.Background(), "abc123")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
}

func TestDataSystemV2DeleteFailed(t *testing.T) {
	coresvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	iamsvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	registrysvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer iamsvr.Close()
	defer registrysvr.Close()
	defer coresvr.Close()
	c := NewNeosClient(iamsvr.URL, registrysvr.URL, coresvr.URL)
	err := c.DataSystemDelete(context.Background(), "abc123")
	if err == nil {
		t.Errorf("expected err to not be nil ")
	}
}

func TestDataSystemV2GetFailed(t *testing.T) {
	coresvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	iamsvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	registrysvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer iamsvr.Close()
	defer registrysvr.Close()
	defer coresvr.Close()
	c := NewNeosClient(iamsvr.URL, registrysvr.URL, coresvr.URL)
	_, err := c.DataSystemGet()
	if err == nil {
		t.Errorf("expected err to be set got nil")
	}
}

func TestDataCreate(t *testing.T) {
	iamClient := NewIAMClient("https://sandbox.city3os.com/api/iam", "owain.perry", "**Marley22")
	loginReq, err := iamClient.Login()
	if err != nil {
		t.Errorf("expected err to be set got nil")
	}

	accessToken := loginReq.AccessToken

	AccessToken = accessToken
	coreClient := NewNeosClient("https://op-02.neosdata.net")

	dspr := DataSystemPostRequest{
		Entity: DataSystemPostRequestEntity{
			Name:        "neos-test1",
			Label:       "ABD",
			Description: "Some description",
		},
		EntityInfo: DataSystemPostRequestEntityInfo{
			Owner:      "some owner 123",
			ContactIds: []string{"abc321"},
			Links:      []string{"link 3"},
		},
	}

	_, err = coreClient.DataSystemPost(context.Background(), dspr)
	if err != nil {
		t.Errorf("expected err to be set got nil")
	}

}

func TestEscapce(t *testing.T) {
	orig := "\"a\"a\"a\"a\"a\"a\"a"
	new := strings.Replace(orig, "\"", "", -1)
	if new != "aaaaaaa" {
		t.Errorf("replace didnt work :(")
	}
}
