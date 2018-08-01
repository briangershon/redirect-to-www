/*
Package redirect contains http middleware for redirecting a naked domain to "www" and "https".

For example, a request to `mydomain.com` will redirect to `https://www.mydomain.com`.

(`localhost` is ignored to avoid redirects when testing sites locally)

Example usage:

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
*/
package redirect
