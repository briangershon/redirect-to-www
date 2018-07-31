# redirect-to-www

Go Middleware for redirecting a naked domain to "www" over "https".

For example, a request to `mydomain.com` will redirect to "https://www.mydomain.com".

(`localhost` is ignored to avoid redirects when testing sites locally)

## Example usage

```
package main

import (
	"net/http"

	redirect "github.com/briangershon/redirect-to-www"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	http.Handle("/", redirect.NakedDomainToWWW(r))
  ...
  ...
}
```

## Run package tests

    go test
