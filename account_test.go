package neos

import (
	"context"
	"testing"
)

func TestAccount(t *testing.T) {
	LoginToGetToken("owain10.neosdata.cloud/api/hub/iam", "neosadmin", "ZWZjYWY4MDll")
	ac := NewAccountClient("https://owain10.neosdata.cloud", NewNeosHttp("root", "KSA"), "root")
	list, err := ac.Get("")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	for _, v := range list.Accounts {
		t.Log(v.Identifier)
	}

	account := AccountPostRequest{
		Name:        "test-a3",
		DisplayName: "dis-test-a1",
		Description: "desc-test-a1",
		Owner:       "own-test-a1",
	}

	acresp, err := ac.Post(context.Background(), account)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	accountPut := AccountPutRequest{
		DisplayName: "up-dis-test-a1",
		Description: "up-desc-test-a1",
		Owner:       "up-own-test-a1",
	}

	acresput, err := ac.Put(context.Background(), acresp.Identifier, accountPut)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = ac.Delete(context.Background(), acresput.Identifier)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
}
