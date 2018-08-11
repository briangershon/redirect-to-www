package main

import (
	"fmt"
	"net/http"

	"github.com/briangershon/redirect-to-www"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	http.Handle("/", redirect.NakedDomainToWWW(r, []string{"localhost", "mysite.appspot.com"}))

	fmt.Println("Listening on :5555")
	fmt.Println(`
Try "curl -v localhost:5555" and no redirect happens.

Update your /etc/hosts file and add:

127.0.0.1 tryme.com
127.0.0.1 mysite.appspot.com

Try "curl -v mysite.appspot.com:5555" and no redirect happens.
Try "curl -v tryme.com:5555" and NOW you should see the direct to "https://www.tryme.com:5555"
`)

	_ = http.ListenAndServe(":5555", http.DefaultServeMux)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from redirect-to-www example.\n")
}
