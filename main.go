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

