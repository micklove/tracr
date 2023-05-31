package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

// create a chi router and register the handlers
func main() {
	// create a chi router
	r := chi.NewRouter()

	// create a handler named hello
	hello := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}

	// register the handlers
	r.Get("/hello", hello)

	// TODO

	// start the server
	log.Fatal(http.ListenAndServe(":3000", r))
}
