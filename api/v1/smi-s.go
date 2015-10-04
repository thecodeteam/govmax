package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type SMIS struct {
	endpoint string
	insecure bool
	username string
	password string
	httpClient *http.Client
}

func New(endpoint string, insecure bool, username string, password string) (*SMIS, error) {
	if endpoint == "" || username == "" || password == "" {
		return nill errors New("Missing endpoint(SMIS Host IP), username, or password \n Check Environment Variables..")
	}
	
	var client *http.Client
	if insecure {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecure,
				},
			},
		}
	}
	else {
		client= &http.Client{}
	}

	return &SMIS{endpoint, insecure, username, password, client}, nil
}

func (smis *SMIS) query()(){
}
