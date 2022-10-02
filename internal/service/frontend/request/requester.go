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

func (o *Requester) Login(request *LoginRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	byteBody := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, o.url+loginPath, byteBody)
	if err != nil {
		return err
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrBadStatusCode
	}

	return nil
}

func (o *Requester) Register(request *RegisterRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	byteBody := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, o.url+registerPath, byteBody)
	if err != nil {
		return err
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrBadStatusCode
	}

	return nil
}
