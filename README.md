# redirect-to-www

[![GoDoc](https://godoc.org/github.com/briangershon/redirect-to-www?status.svg)](https://godoc.org/github.com/briangershon/redirect-to-www)

Package redirect contains http middleware for redirecting a naked domain to "www" and "https".

For example, a request to `http://mydomain.com` will redirect to `https://www.mydomain.com`.

You can also provide a list of hosts to ignore such as `localhost` and a PaaS host like `mysite.appspot.com`.

Example usage:

    go run example/example.go

## Run package tests

    go test ./...
