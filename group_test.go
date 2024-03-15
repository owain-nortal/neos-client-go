package neos

import (
	"context"
	"testing"
)

func TestGroup(t *testing.T) {
	LoginToGetToken("owain10.neosdata.cloud/api/hub/iam", "neosadmin", "ZWZjYWY4MDll")
	ac := NewGroupClient("https://owain10.neosdata.cloud", NewNeosHttp("root", "KSA"), "root")
	list, err := ac.List("root")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	for _, v := range list.Groups {
		t.Log(v.Identifier)
	}

	group := GroupPostRequest{
		Name:        "test-a5",
		Description: "desc-test-a1",
	}

	acresp, err := ac.Post(context.Background(), group, "root")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	groupPut := GroupPutRequest{
		Name:        "up-dis-test-a5",
		Description: "up-desc-test-a5",
	}

	acresput, err := ac.Put(context.Background(), acresp.Identifier, groupPut, "root")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = ac.Delete(context.Background(), acresput.Identifier, "root")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
}
