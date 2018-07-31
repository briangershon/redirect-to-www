// Package redirect contains http middleware that ensures a naked domain (e.g. http://mydomain.com) is redirected to "www" subdomains aver https (i.e. https://www.mydomain.com)
package redirect

import (
	"fmt"
	"net/http"
	"strings"
)

// NakedDomainToWWW is http middleware that ensures a naked domain (e.g. http://mydomain.com) is redirected to "www" subdomains aver https (i.e. https://www.mydomain.com)
// "localhost" is ignored to avoid problems when testing locally.
func NakedDomainToWWW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := strings.ToLower(r.Host)
		if !strings.HasPrefix(host, "www") && !strings.HasPrefix(host, "localhost") {
			http.Redirect(w, r, fmt.Sprintf("https://www.%s%s", r.Host, r.URL.Path), http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
