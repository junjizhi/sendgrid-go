package sendgrid

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGlobalSuppressionGet(t *testing.T) {
	setup()
	defer teardown()

	data := loadTestData("global_suppression.get.json", t)
	defer data.Close()

	const email = "test1@example.com"
	mux.HandleFunc("/asm/suppressions/global/", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		checkURLSuffix(t, r, email)
		io.Copy(w, data)
	})

	a, err := client.GlobalSuppression.Get(email)
	if err != nil {
		t.Fatal(err)
	}
	if a.RecipientEmail != email {
		t.Fatalf("Email = %v, want %v", a.RecipientEmail, email)
	}

	testBadURL(t, func() error {
		_, err := client.GlobalSuppression.Get(email)
		return err
	})
}

func TestGlobalSuppressionDelete(t *testing.T) {
	setup()
	defer teardown()

	const email = "test1@example.com"
	const nonExistentEmail = "none@example.com"
	mux.HandleFunc("/asm/suppressions/global/", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "DELETE")
		split := strings.Split(r.URL.Path, "/")
		if split[4] == nonExistentEmail {
			http.Error(w, "invalid email", http.StatusNotFound)
		}
	})

	if err := client.GlobalSuppression.Delete(email); err != nil {
		t.Fatal(err)
	}

	if err := client.GlobalSuppression.Delete(nonExistentEmail); err == nil {
		t.Fatal("expected HTTP 404")
	}

	testBadURL(t, func() error {
		return client.GlobalSuppression.Delete(email)
	})
}
