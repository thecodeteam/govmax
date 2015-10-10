package apiv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"encoding/base64"
	"bytes"
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

func (smis *SMIS) query(httpType, objectPath string, body, resp interface{}) error {

	///////////////////////////////////////////
	//      Setup http/JSON header request   //
	///////////////////////////////////////////

	if smis.insecure {
		URL := strings.Join([]string{"http://", smis.host, ":", smis.port, objectPath}, "")
	}
	else {
		URL := strings.Join([]string{"https://", smis.host, ":", smis.port, objectPath}, "")
	}

	// Parse out body interface into JSON
	bodyBytes, _ := json.Marshal(body)

	// Create http resquest & add auth
	req, err := http.NewRequest(httpType, URL, bytes.NewBuffer(bodyBytes))
	req.SetBasicAuth(smis.username, smis.password)

	// Create Authorization token which is then added to header of request. Auth token is in the format of username:password then encoded in base64.
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(smis.username + ":" + smis.password))

	// Add header specific items
	req.Header.Add("Authorization", "Basic " + encodedAuth)
	req.Header.Add("Accept","application/json")
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Connection","Keep-Alive")
	req.Header.Add("Host",smis.host + ":" + smis.port)

	// Perform request
	httpResp, err := smis.client.Do(req)
	if err != nill {
		return err
	}

	// Cleanup Response
	defer httpResp.Body.Close()

	// Deal with errors
	switch {
	case httpResp.StatusCode == 400:
		return errors.New("JSON Build Error")
	case httpResp.StatusCode == 401:
		return errors.New("JSON Auth Error")
	case httpResp.StatusCode == 404:
		return errors.New("Object Could not be found")
	case httpResp == nil:
		return nil
	// Decode JSON of response into our interface defined for the specific request sent
	case httpResp.StatusCode == 200 || httpResp.StatusCode == 201:
		err = json.NewDecoder(httpResp.Body).Decode(resp)
		return err
	default:
		return errors.New("JSON Build Error")
	}
}
