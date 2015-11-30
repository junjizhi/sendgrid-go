package sendgrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// SGClient will contain the credentials and default values
type SGHTTPClient struct {
	apiUser           string
	apiPwd            string
	baseURI           string
	NumRequests       int
	JSONWriter        io.Writer
	client            http.Client
	Bounce            *BounceService
	ASMGroup          *ASMGroupService
	ASMSuppression    *ASMSuppressionService
	GlobalSuppression *GlobalSuppressionService
	APIKey            *APIKeyService
}

var baseURI = "https://api.sendgrid.com/v3"

// NewSendGridHTTPClient will return a new SG HTTP Client. Used for username and password
func NewSendGridHTTPClient(apiUser, apiPwd string) *SGHTTPClient {

	c := &SGHTTPClient{}
	c.apiUser = apiUser
	c.apiPwd = apiPwd
	c.Bounce = &BounceService{c}
	c.ASMGroup = &ASMGroupService{c}
	c.ASMSuppression = &ASMSuppressionService{c}
	c.GlobalSuppression = &GlobalSuppressionService{c}
	c.APIKey = &APIKeyService{c}
	return c
}

// NewSendGridHTTPClientWithAPIKey will return a new SG HTTP Client. Used for api key
func NewSendGridHTTPClientWithAPIKey(apiKey string) *SGHTTPClient {

	c := &SGHTTPClient{}
	c.apiPwd = apiKey
	c.Bounce = &BounceService{c}
	c.ASMGroup = &ASMGroupService{c}
	c.ASMSuppression = &ASMSuppressionService{c}
	c.GlobalSuppression = &GlobalSuppressionService{c}
	c.APIKey = &APIKeyService{c}
	return c
}

func (sg *SGHTTPClient) NewRequest(method string, endpoint string, data interface{}) (req *http.Request, err error) {
	var u *url.URL
	u, err = url.Parse(baseURI)
	if err != nil {
		return
	}
	u.Path += endpoint

	var dataVals url.Values
	if data != nil {
		dataVals, err = query.Values(data)
		if err != nil {
			return
		}
	} else {
		dataVals = url.Values{}
	}

	switch method {
	case "GET":
		fallthrough
	case "DELETE":
		u.RawQuery = dataVals.Encode()
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return
		}
		if sg.apiUser == "" {
			req.Header.Set("Authorization", "Bearer "+sg.apiPwd)
		} else {
			req.SetBasicAuth(sg.apiUser, sg.apiPwd)
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
	case "POST":
		fallthrough
	case "PUT":
		q := url.Values{}
		u.RawQuery = q.Encode()

		payload := dataVals.Encode()
		body := bytes.NewBufferString(payload)
		req, err = http.NewRequest(method, u.String(), body)
		if err != nil {
			return
		}
		if sg.apiUser == "" {
			req.Header.Set("Authorization", "Bearer "+sg.apiPwd)
		} else {
			req.SetBasicAuth(sg.apiUser, sg.apiPwd)
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
	default:
		err = fmt.Errorf("Unknown HTTP method: %s", method)
	}

	return
}

// Do performs the given http.Request and optionally
// decodes the JSON response into the given data struct.
func (sg *SGHTTPClient) Do(req *http.Request, data interface{}) error {
	fmt.Println(req.Method, req.URL)
	resp, err := sg.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("HTTP Error %d", resp.StatusCode)
	}

	sg.NumRequests++

	if data != nil {
		var body io.Reader
		// if the client has a JSONWriter, also dump JSON responses
		if sg.JSONWriter != nil {
			body = io.TeeReader(resp.Body, sg.JSONWriter)
		} else {
			body = resp.Body
		}

		err = json.NewDecoder(body).Decode(data)

	}

	return err
}
