package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var TcUrl = func() url.URL {
	base := os.Getenv("TEST_BASE_URL")
	if base == "" {
		base = "http://localhost:8082"
	}
	u, _ := url.Parse(base)
	return *u
}()

const (
	testLogin    = "test_user"
	testPassword = "test_password_123"
	testUsername = "Test User"
)

var (
	authToken string
	authOnce  sync.Once
)

func GetAuthToken() string {
	authOnce.Do(func() {
		base := TcUrl.String()

		regBody, _ := json.Marshal(map[string]string{
			"username": testUsername,
			"login":    testLogin,
			"password": testPassword,
		})
		http.Post(base+"/user/sign-up", "application/json", bytes.NewReader(regBody)) //nolint:errcheck

		loginBody, _ := json.Marshal(map[string]string{
			"login":    testLogin,
			"password": testPassword,
		})
		resp, err := http.Post(base+"/user/sign-in", "application/json", bytes.NewReader(loginBody))
		if err != nil {
			panic(fmt.Sprintf("auth: failed to sign in: %v", err))
		}
		defer resp.Body.Close()

		var result struct {
			Token string `json:"TokenJWT"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			panic(fmt.Sprintf("auth: failed to decode sign-in response: %v", err))
		}
		if result.Token == "" {
			panic("auth: sign-in returned empty token")
		}
		authToken = result.Token
	})
	return authToken
}

type authTransport struct{}

func (a *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.Header.Set("Authorization", "Bearer "+GetAuthToken())
	return http.DefaultTransport.RoundTrip(clone)
}

func NewExpect(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  TcUrl.String(),
		Reporter: httpexpect.NewAssertReporter(t),
		Client:   &http.Client{Transport: &authTransport{}},
	})
}
