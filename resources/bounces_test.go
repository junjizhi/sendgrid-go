package sendgrid

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestBounceGet(t *testing.T) {
	setup()
	defer teardown()

	data := loadTestData("bounce.get.json", t)
	defer data.Close()

	const email = "testemail1@test.com"
	mux.HandleFunc("/suppression/bounces/", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		checkURLSuffix(t, r, email)
		io.Copy(w, data)
	})

	b, err := client.Bounce.Get(email)
	if err != nil {
		t.Fatal(err)
	}
	bResp := b[0]
	if bResp.Email != email {
		t.Fatalf("Bounce email address = %v, want %v", bResp.Email, email)
	}

	testBadURL(t, func() error {
		_, err := client.Bounce.Get(email)
		return err
	})
}

func TestBounceList(t *testing.T) {
	setup()
	defer teardown()

	data := loadTestData("bounce.list.json", t)
	defer data.Close()

	const email = "testemail1@test.com"
	mux.HandleFunc("/suppression/bounces/", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		io.Copy(w, data)
	})

	b, err := client.Bounce.List(&BounceListRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if len(b) < 2 {
		t.Fatalf("Should include 2 or more bounces")
	}

	testBadURL(t, func() error {
		_, err := client.Bounce.List(&BounceListRequest{})
		return err
	})
}

func TestBounceDelete(t *testing.T) {
	setup()
	defer teardown()

	const email = "testemail1@test.com"
	const emailNil = "noemail@none.com"
	mux.HandleFunc("/suppression/bounces/", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "DELETE")
		split := strings.Split(r.URL.Path, "/")
		if split[3] == emailNil {
			http.Error(w, "invalid email address", http.StatusNotFound)
		}
	})

	if err := client.Bounce.Delete(email); err != nil {
		t.Fatal(err)
	}

	if err := client.Bounce.Delete(emailNil); err == nil {
		t.Fatal("expected HTTP 404")
	}

	testBadURL(t, func() error {
		return client.Bounce.Delete(email)
	})
}
