package neos

import (
	"fmt"
	"net/url"
	"time"
)

func MaintainAccessToken(iamUrl string, username string, password string) {
	nextLoginTime := time.Now()

	for {
		if time.Now().After(nextLoginTime) {
			nextLoginTime = LoginToGetToken(iamUrl, username, password)
			time.Sleep(time.Second)
		}
	}

}

func LoginToGetToken(iamUrl string, username string, password string) time.Time {
	var nextLoginTime time.Time
	iam := NewIAMClient(iamUrl, username, password)
	loginResult, err := iam.Login()
	if err != nil {
		fmt.Println(err)
	}
	AccessToken = loginResult.AccessToken
	expires, err := loginResult.TokenExpires()
	if err != nil {
		fmt.Println(err)
	} else {
		nextLoginTime = time.Now().Add(expires)
	}

	return nextLoginTime
}

func GetAccessToken() string {
	return AccessToken
}

var AccessToken = ""

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

// func maintainAccessToken(iamUrl string, username string, password string) {
// 	nextLoginTime := time.Now()
// 	iam := NewIAMClient(iamUrl, username, password)
// 	for {
// 		if time.Now().After(nextLoginTime) {
// 			loginResult, err := iam.Login()
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			AccessToken = loginResult.AccessToken
// 			expires, err := loginResult.TokenExpires()
// 			if err != nil {
// 				fmt.Println(err)
// 			} else {
// 				nextLoginTime = time.Now().Add(expires)
// 			}
// 			time.Sleep(time.Second)
// 		}
// 	}

// }

// func GetAccessToken() string {
// 	return AccessToken
// }

// var AccessToken = ""

// func main() {
// 	iamUrl := "https://sandbox.city3os.com/api/iam"
// 	url := "https://op-02.neosdata.net"
// 	username := "owain.perry"
// 	password := "**Marley22"
// 	go maintainAccessToken(iamUrl, username, password)
// 	sum := 0
// 	for i := 1; i < 60; i++ {
// 		ds := NewDataSystemClientV2(url)
// 		results, err := ds.Get(true)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		for _, v := range results.Entities {
// 			fmt.Println(v.Identifier)
// 		}
// 		time.Sleep(time.Second)
// 	}
// 	fmt.Println(sum)

// }

type NeosClient struct {
	hubHost                    string
	coreHost                   string
	scheme                     string
	coreUri                    string
	AccountClient              AccountClient
	DataProductClient          DataProductClient
	DataSourceClient           DataSourceClient
	DataSourceConnectionClient DataSourceConnectionClient
	DataSourceSecretClient     DataSourceSecretClient
	DataSystemClient           DataSystemClient
	DataProductSchemaClient    DataProductSchemaClient
	DataUnitClient             DataUnitClient
	GroupClient                GroupClient
	IAMClient                  IAMClient
	LinksClient                LinksClient
	OutputClient               OutputClient
	PolicyClient               PolicyClient
	RegistryCoreClient         RegistryCoreClient
	SecretClient               SecretClient
	UserClient                 UserClient
}

func NewNeosClient(hubHost, coreHost string, scheme string, account string, partition string) (NeosClient, error) {
	var rtn NeosClient

	coreUri, err := resolveUri(coreHost, scheme)
	if err != nil {
		return rtn, err
	}

	hubUri, err := resolveUri(hubHost, scheme)
	if err != nil {
		return rtn, err
	}

	httpClient := NewNeosHttp(account, partition)

	rtn = NeosClient{
		hubHost:                    hubHost,
		coreHost:                   coreHost,
		scheme:                     scheme,
		coreUri:                    coreUri,
		AccountClient:              *NewAccountClient(hubUri, httpClient, account),
		DataProductClient:          *NewDataProductClient(coreUri, httpClient, account),
		DataSourceClient:           *NewDataSourceClient(coreUri, httpClient, account),
		DataSourceConnectionClient: *NewDataSourceConnectionClient(coreUri, httpClient, account),
		DataSourceSecretClient:     *NewDataSourceSecretClient(coreUri, httpClient, account),
		DataSystemClient:           *NewDataSystemClient(coreUri, httpClient, account),
		DataProductSchemaClient:    *NewDataProductSchemaClient(coreUri, httpClient, account),
		DataUnitClient:             *NewDataUnitClient(coreUri, httpClient, account),
		GroupClient:                *NewGroupClient(hubUri, httpClient, account),
		LinksClient:                *NewLinksClient(coreUri, httpClient, account),
		OutputClient:               *NewOutputClient(coreUri, httpClient, account),
		PolicyClient:               *NewPolicyClient(hubUri, httpClient, account),
		RegistryCoreClient:         *NewRegistryCoreClient(hubUri, httpClient, account),
		SecretClient:               *NewSecretClient(coreUri, httpClient, account),
		UserClient:                 *NewUserClient(hubUri, httpClient, account),
	}
	return rtn, nil
}

func resolveUri(host string, scheme string) (string, error) {
	coreHostUri, err := url.Parse(host)
	if err != nil {
		return "", err
	}

	coreUri := fmt.Sprintf("%s://%s", scheme, host)
	if coreHostUri.Scheme != "" {
		coreUri = host
	}
	return coreUri, nil
}
