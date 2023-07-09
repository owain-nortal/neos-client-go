package neos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedJson)
	}))

	defer svr.Close()

	request := DataSystemPostRequest{
		Entity: DataSystemPostRequestEntity{
			Name: expected.Name,
		},
		EntityInfo: DataSystemPostRequestEntityInfo{
			Owner: "Mr Owner",
		},
	}

	c := NewDataSystemClientV2(svr.URL)
	res, err := c.Post(request)
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
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedJson)
	}))

	defer svr.Close()

	request := DataSystemPutRequest{
		Entity: DataSystemPutRequestEntity{
			Name: expected.Name,
		},
	}

	c := NewDataSystemClientV2(svr.URL)
	res, err := c.Put("123", request)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if res.Name != request.Entity.Name {
		t.Errorf("expected entiry name to be %s got %s", request.Entity.Name, res.Name)
	}
}

func TestDataSystemV2PutFailed(t *testing.T) {

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer svr.Close()
	request := DataSystemPutRequest{}
	c := NewDataSystemClientV2(svr.URL)
	_, err := c.Put("321ads", request)
	if err == nil {
		t.Errorf("expected err to be set not nil")
	}
}

func TestDataSystemV2PostFailed(t *testing.T) {

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer svr.Close()
	request := DataSystemPostRequest{}
	c := NewDataSystemClientV2(svr.URL)
	_, err := c.Post(request)
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
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedJson)
	}))

	defer svr.Close()
	c := NewDataSystemClientV2(svr.URL)
	res, err := c.Get(false)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if res.Entities[0].Identifier != expected.Entities[0].Identifier {
		t.Errorf("expected res to be %s got %s", expected.Entities[0].Identifier, res.Entities[0].Identifier)
	}
}

func TestDataSystemV2DeleteOK(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer svr.Close()
	c := NewDataSystemClientV2(svr.URL)
	err := c.Delete("abc123")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
}

func TestDataSystemV2DeleteFailed(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))

	defer svr.Close()
	c := NewDataSystemClientV2(svr.URL)
	err := c.Delete("abc123")
	if err == nil {
		t.Errorf("expected err to not be nil ")
	}
}

func TestDataSystemV2GetFailed(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))

	defer svr.Close()
	c := NewDataSystemClientV2(svr.URL)
	_, err := c.Get(false)
	if err == nil {
		t.Errorf("expected err to be set got nil")
	}
}
