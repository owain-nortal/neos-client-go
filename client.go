package neos

import (
	"fmt"
	"time"
)

func MaintainAccessToken(iamUrl string, username string, password string) {
	nextLoginTime := time.Now()
	iam := NewIAMClient(iamUrl, username, password)
	for {
		if time.Now().After(nextLoginTime) {
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
			time.Sleep(time.Second)
		}
	}

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
	iamHost      string
	registryHost string
	coreHost     string
	scheme       string
	coreUri      string
	iamUri       string
	registryUri  string
}

func NewNeosClient(iamHost string, registryHost string, coreHost string, scheme string) NeosClient {
	return NeosClient{
		iamHost:      iamHost,
		registryHost: registryHost,
		coreHost:     coreHost,
		scheme:       scheme,
		coreUri:      fmt.Sprintf("%s://%s", scheme, coreHost),
		iamUri:       fmt.Sprintf("%s://%s", scheme, iamHost),
		registryUri:  fmt.Sprintf("%s://%s", scheme, registryHost),
	}
}
