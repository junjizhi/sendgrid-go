package sendgrid

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestAPIKeyGet(t *testing.T) {
	setup()
	defer teardown()

	data := loadTestData("api_key.get.json", t)
	defer data.Close()

	const apiKey = "123445678"
	mux.HandleFunc("/api_keys/", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		checkURLSuffix(t, r, apiKey)
		io.Copy(w, data)
	})

	a, err := client.APIKey.Get(apiKey)
	if err != nil {
		t.Fatal(err)
	}

	if a.ID != apiKey {
		t.Fatalf("APIKey ID = %v, want %v", a.ID, apiKey)
	}

	testBadURL(t, func() error {
		_, err := client.APIKey.Get(apiKey)
		return err
	})
}

// func TestAPIKeysList(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	data := loadTestData("api_keys.list.json", t)
// 	defer data.Close()

// 	mux.HandleFunc("/api_keys", func(w http.ResponseWriter, r *http.Request) {
// 		checkMethod(t, r, "GET")
// 		io.Copy(w, data)
// 	})

// 	a, err := client.APIKey.List()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Println(a)
// 	// if len(a.Result) < 2 {
// 	// 	t.Fatalf("Should include 2 or more asm groups")
// 	// }

// 	testBadURL(t, func() error {
// 		_, err := client.APIKey.List()
// 		return err
// 	})
// }

func TestAPIKeyDelete(t *testing.T) {
	setup()
	defer teardown()

	const apiKeyID = "12345678"
	const nonExistAPIKey = "9999999"
	mux.HandleFunc("/api_keys/", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "DELETE")
		split := strings.Split(r.URL.Path, "/")
		if split[2] == nonExistAPIKey {
			http.Error(w, "invalid id", http.StatusNotFound)
		}
	})

	if err := client.APIKey.Delete(apiKeyID); err != nil {
		t.Fatal(err)
	}

	if err := client.APIKey.Delete(nonExistAPIKey); err == nil {
		t.Fatal("expected HTTP 404")
	}

	testBadURL(t, func() error {
		return client.APIKey.Delete(apiKeyID)
	})
}
