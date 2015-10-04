package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"encoding/base64"
)

type SMIS struct {
	host string
	port string
	insecure bool
	username string
	password string
	httpClient *http.Client
}

func New(host string, port string, insecure bool, username string, password string) (*SMIS, error) {
	if host == "" || port == "" || username == "" || password == "" {
		return nill errors New("Missing host (SMIS Host IP), port (SMIS Host Port), username, or password \n Check Environment Variables..")
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

	return &SMIS{host, port, insecure, username, password, client}, nil
}

func (smis *SMIS) query(httpType string, objectPath string, body io.Reader) error {

	///////////////////////////////////////////
	//      Setup http/JSON header request   //
	///////////////////////////////////////////

	if smis.insecure {
		URL := strings.Join([]string{"http://", smis.host, ":", smis.port, objectPath}, "")
	}
	else {
		URL := strings.Join([]string{"https://", smis.host, ":", smis.port, objectPath}, "")
	}


	req, err := http.NewRequest(httpType, URL, body)
	req.SetBasicAuth(smis.username, smis.password)

	// Create Authorization token which is then added to header of request. Auth token is in the format of username:password then encoded in base64.
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(smis.username + ":" + smis.password))

	// Add header specific items
	req.Header.Add("Authorization", "Basic" + encodedAuth)
	req.Header.Add("Accept","application/json")
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Connection","Keep-Alive")
	req.Header.Add("Host",smis.host + ":" + smis.port)

	// Perform request
	resp, err := smis.client.Do(req)
	if err != nill {
		return err
	}
	resp.Body.Close()

	switch {
	case resp.StatusCode == 400:
		return errors.New("JSON Build Error")

	case resp == nil:
		return nil

}
