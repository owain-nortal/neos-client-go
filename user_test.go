package neos

import (
	"context"
	"testing"
)

func TestUser(t *testing.T) {
	LoginToGetToken("owain10.neosdata.cloud/api/hub/iam", "neosadmin", "ZWZjYWY4MDll")
	ac := NewUserClient("https://owain10.neosdata.cloud", NewNeosHttp("root", "KSA"))
	list, err := ac.List("", "", "")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	for _, v := range list.Users {
		t.Log(v.Identifier)
	}

	user := UserPostRequest{
		Username:  "user3",
		FirstName: "first1",
		LastName:  "last1",
		Email:     "test3@example.com",
	}

	acresp, err := ac.Post(context.Background(), user)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}


	err = ac.Delete(context.Background(), acresp.Identifier, "")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
}
