package redirect

import (
	"fmt"
	"net/http"
	"strings"
)

// NakedDomainToWWW performs a redirect to https://www.mydomain.com if request
// is for naked mydomain.com.
// "localhost" is ignored to avoid problems when testing locally.
func NakedDomainToWWW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(strings.ToLower(r.Host), "localhost") {
			if !strings.HasPrefix(strings.ToLower(r.Host), "www") {
				http.Redirect(w, r, fmt.Sprintf("https://www.%s%s", r.Host, r.URL.Path), http.StatusPermanentRedirect)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
