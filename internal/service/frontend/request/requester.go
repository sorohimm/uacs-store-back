package request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func NewRequester(version, url string) *Requester {
	return &Requester{
		client: http.DefaultClient,
		url:    url + "/" + version,
	}
}

type Requester struct {
	client *http.Client
	url    string
}

func (o *Requester) Login(request *LoginRequest) (*JWT, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	byteBody := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:2104"+"/v1"+loginPath, byteBody)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadStatusCode
	}

	jwt, err := extractJWT(resp)
	if err != nil {
		return nil, err
	}

	return jwt, nil
}

func (o *Requester) Register(request *RegisterRequest) (*JWT, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	byteBody := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:2104"+"/v1"+registerPath, byteBody)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadStatusCode
	}
	jwt, err := extractJWT(resp)
	if err != nil {
		return nil, err
	}

	return jwt, nil
}

func extractJWT(resp *http.Response) (*JWT, error) {
	t := resp.Header.Get(accessTokenHeaderKey)   // access token
	rt := resp.Header.Get(refreshTokenHeaderKey) // refresh token
	if t == "" || rt == "" {
		return nil, ErrBadResponseHeader
	}
	return &JWT{AccessToken: t, RefreshToken: rt}, nil
}
