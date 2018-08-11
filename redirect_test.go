package redirect

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedirectNakedToWWW(t *testing.T) {
	tests := []struct {
		description                 string
		host                        string
		excludes                    []string
		expectedLocationHeader      bool
		expectedLocationHeaderValue string
		expectedCode                int
	}{
		{
			description:                 "naked domain should redirect to https://www.mydomain.com",
			host:                        "mydomain.com",
			excludes:                    []string{},
			expectedLocationHeader:      true,
			expectedLocationHeaderValue: "https://www.mydomain.com/abc",
			expectedCode:                308,
		},
		{
			description:                 "subdomains will also redirect to https://www.subdomain.mydomain.com",
			host:                        "subdomain.mydomain.com",
			excludes:                    []string{},
			expectedLocationHeader:      true,
			expectedLocationHeaderValue: "https://www.subdomain.mydomain.com/abc",
			expectedCode:                308,
		},
		{
			description:                 "www domain should not redirect",
			host:                        "www.mydomain.com",
			excludes:                    []string{},
			expectedLocationHeader:      false,
			expectedLocationHeaderValue: "",
			expectedCode:                200,
		},
		{
			description: "localhost domain should not redirect",
			excludes:    []string{"localhost"},
			host:        "localhost",
			expectedLocationHeader:      false,
			expectedLocationHeaderValue: "",
			expectedCode:                200,
		},
		{
			description: "localhost and serverless domain should not redirect",
			excludes:    []string{"localhost", "mysite.appspot.com"},
			host:        "mysite.appspot.com",
			expectedLocationHeader:      false,
			expectedLocationHeaderValue: "",
			expectedCode:                200,
		},
	}

	for _, tc := range tests {
		log.Printf("Testing %q", tc.description)
		rr := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/abc", nil)
		req.Host = tc.host
		if err != nil {
			t.Fatal(err)
		}

		handler := NakedDomainToWWW(GetTestHandler(), tc.excludes)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, tc.expectedCode, rr.Code)
		if tc.expectedLocationHeader {
			assert.Equal(t, tc.expectedLocationHeaderValue, rr.Header().Get("Location"))
		} else {
			if len(rr.Header().Get("Location")) > 0 {
				assert.Fail(t, "Location header should not be present")
			}
		}
	}
}

// GetTestHandler returns a http.HandlerFunc for testing http middleware
func GetTestHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}
