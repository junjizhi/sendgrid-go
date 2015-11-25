package sendgrid

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	mux      *http.ServeMux
	server   *httptest.Server
	client   *SGHTTPClient
	fakeUser = "username"
	fakePwd  = "password"
)

func loadTestData(filename string, t *testing.T) io.ReadCloser {
	data, err := os.Open("fixtures/" + filename)
	if err != nil {
		t.Fatal("Failed to open fixture data file")
	}
	return data
}

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	baseURI = server.URL
	client = NewSendGridHTTPClient(fakeUser, fakePwd)
}

func teardown() {
	server.Close()
}

// Checks that the HTTP Request's method matches the given method.
func checkMethod(t *testing.T, r *http.Request, method string) {
	if method != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, method)
	}
}

// Checks that the HTTP Request's URL path ends with suffix, ignoring any trailing slashes.
func checkURLSuffix(t *testing.T, r *http.Request, suffix string) {
	if !strings.HasSuffix(strings.TrimSuffix(r.URL.Path, "/"), suffix) {
		t.Fatalf("URL path = %s, expected suffix = %s", r.URL.Path, suffix)
	}
}

// Executes fn, expecting it to return an error
func testBadURL(t *testing.T, fn func() error) {
	origURL := baseURI
	baseURI = "https://%api.sendgrid.com/v3"
	if err := fn(); err == nil {
		t.Fatal("expected HTTP Request URL error")
	}
	baseURI = origURL
}
