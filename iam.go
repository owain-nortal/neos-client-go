package neos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"	
	"time"
)

type IAMClient struct {
	url      string
	username string
	password string
}

type LoginResponse struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	SessionState     string `json:"session_state"`
}

func (l *LoginResponse) TokenExpires() (time.Duration, error) {
	//expiresTotal , err := strconv.Atoi(l.ExpiresIn)
	// refresh in half the time
	var err error
	expires := l.ExpiresIn / 2
	rtn := time.Duration(expires) * time.Second
	return rtn, err
}

type LoginRequest struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

func NewIAMClient(url string, username string, password string) IAMClient {
	return IAMClient{url, username, password}
}

func (c IAMClient) Login() (LoginResponse, error) {
	ar := LoginResponse{}
	url := fmt.Sprintf("https://%s%s", c.url, "/login")
	loginJson := fmt.Sprintf("{\"user\":\"%s\",\"password\":\"%s\"}", c.username, c.password)
	var jsonStr = []byte(loginJson)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return ar, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ar, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return ar, fmt.Errorf("%s", body)
	}

	derr := json.Unmarshal(body, &ar)
	if derr != nil {
		return ar, derr
	}

	return ar, nil
}
