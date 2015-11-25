package sendgrid

import "net/http"

type APIKeyService struct {
	sg *SGHTTPClient
}

type APIKey struct {
	Name   string   `json:"name"`
	ID     string   `json:"api_key_id"`
	Scopes []string `json:"scopes"`
}

// type APIKeyList struct {
// 	Result struct {
// 		Name string `json:"name"`
// 		ID   string `json:"api_key_id"`
// 	} `json:"result"`
// }

// func (ak *APIKeyService) List() (al []APIKeyList, err error) {
// 	// GET: /api_keys
// 	var req *http.Request
// 	req, err = ak.sg.NewRequest("GET", "/api_keys", nil)
// 	if err != nil {
// 		return
// 	}

// 	err = ak.sg.Do(req, &al)
// 	return al, err
// }

func (ak *APIKeyService) Get(APIKey string) (k APIKey, err error) {
	// GET: /api_keys/:api_key_id
	var req *http.Request
	req, err = ak.sg.NewRequest("GET", "/api_keys/"+APIKey, nil)
	if err != nil {
		return
	}

	err = ak.sg.Do(req, &k)

	return k, err
}

func (ak *APIKeyService) Delete(apiKeyID string) error {
	// DELETE: /api_keys/:api_key_id
	var req *http.Request
	req, err := ak.sg.NewRequest("DELETE", "/api_keys/"+apiKeyID, nil)
	if err != nil {
		return err
	}

	return ak.sg.Do(req, nil)
}
