package signhost

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

var (
	tMux    *http.ServeMux
	tServer *httptest.Server
	tClient *Client
	tConf   *Config
)

func setup() {
	tMux = http.NewServeMux()
	tServer = httptest.NewServer(tMux)
	tConf = NewConfig(true, APITokenEnv, AppKeyEnv)
	tClient, _ = NewClient(nil, tConf)
	u, _ := url.Parse(tServer.URL + "/")
	tClient.BaseURL = u
}

func teardown() {
	tServer.Close()
}

func TestClient_NewAPIRequest(t *testing.T) {
	setEnv()
	setup()
	defer teardown()

	b := []string{"hello", "bye"}
	inURL, outURL := "test", tServer.URL+"/test"
	inBody, outBody := b, `["hello","bye"]`+"\n"
	_ = tClient.WithAuthenticationValue("test_token", "test_app")
	req, _ := tClient.NewAPIRequest("GET", inURL, inBody)

	testHeader(t, req, "Accept", RequestAccept)
	testHeader(t, req, AuthHeader, "APIKey test_token")
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}
}

func TestClient_NewAPIRequest_ErrTrailingSlash(t *testing.T) {
	uri, _ := url.Parse("http://localhost")
	tClient = &Client{
		BaseURL: uri,
	}
	_, err := tClient.NewAPIRequest("GET", "test", nil)

	if err == nil {
		t.Errorf("expected error %v not occurred, got %v", errBadBaseURL, err)
	}
}

func TestClient_NewAPIRequest_HTTPReqNativeError(t *testing.T) {
	setup()
	defer teardown()
	_, err := tClient.NewAPIRequest("\\\\\\", "test", nil)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "invalid method") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_NewAPIRequest_ApiKeyAppToken(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("api_key", "app_token")
	req, _ := tClient.NewAPIRequest("GET", "test", nil)

	testHeader(t, req, AuthHeader, "APIKey api_key")
	testHeader(t, req, ApplicationHeader, "APPKey app_token")
}

func TestClient_WithAuthenticationValue_Error(t *testing.T) {
	setup()
	defer teardown()
	err := tClient.WithAuthenticationValue("", "")

	if err == nil {
		t.Errorf("unexpected error, want %v and got %v", errEmptyAuthKey, err)
	}
}

func TestClient_NewAPIRequest_ErrorBodySerialization(t *testing.T) {
	setup()
	defer teardown()
	b := make(chan int)
	_, err := tClient.NewAPIRequest("GET", "test", b)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "unsupported type") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_NewAPIRequest_NativeURLParseError(t *testing.T) {
	setup()
	defer teardown()
	_, err := tClient.NewAPIRequest("GET", ":", nil)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "parse") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_Do(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, AuthHeader, "APIKey api_key")
		testHeader(t, r, ApplicationHeader, "APPKey app_token")
		w.WriteHeader(http.StatusOK)
	})
	_ = tClient.WithAuthenticationValue("api_key", "app_token")
	req, _ := tClient.NewAPIRequest("GET", "test", nil)
	res, err := tClient.Do(req)

	if err != nil {
		t.Errorf("unexpected error received: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("request failed: %+v", res)
	}
}

func setEnv() {
	_ = os.Setenv(APITokenEnv, "banana_token")
	_ = os.Setenv(AppKeyEnv, "dsfksdjfksjdlfs")
}

func unsetEnv() {
	_ = os.Unsetenv(APITokenEnv)
	_ = os.Unsetenv(AppKeyEnv)
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

