# redirect-to-www

Go Middleware for redirecting a naked domain to "www" and "https".

For example, a request to `mydomain.com` will redirect to `https://www.mydomain.com`.

(`localhost` is ignored to avoid redirects when testing sites locally)

## Docs and Sample Usage

[![GoDoc](https://godoc.org/github.com/briangershon/redirect-to-www?status.svg)](https://godoc.org/github.com/briangershon/redirect-to-www)

## Run package tests

    go test
