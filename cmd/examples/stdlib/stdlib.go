package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/micklove/tracr"
	mid "github.com/micklove/tracr/middleware"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cid, err := tracr.GetCID(r.Context())
	if err != nil {
		log.Printf("GetCID() = returned error %v", err)
	}
	fmt.Fprint(w, fmt.Sprintf("cid: %s", cid))
}

func main() {

	// Create the correlation ID options, in this case, we will simply hard code in the funcs that will
	// retrieve the correlation ID http header name and the correlation ID value. This allows the caller to
	// use whatever strategy they want, to generate the correlation ID e.g. DB Sequence, UUID, GUID, etc... and / or
	// to use their preferred strategy for naming the correlation id http header (e.g. from env var, config , ssm, etc..)
	// Note, the correlation ID http header name function is optional, if not provided, the default will be used.
	correlationIDOptions := tracr.CorrelationIDOptions{
		CorrelationIDGeneratorFn: func() (string, error) {
			return uuid.Must(uuid.NewV4()).String(), nil
		},
	}

	// example usage of the correlation id AND http-logger middleware
	http.HandleFunc("/", mid.MiddlewareCorrelationID(mid.LoggingMiddleware(indexHandler, false), correlationIDOptions, nil))
	http.HandleFunc("/logger-with-body", mid.MiddlewareCorrelationID(mid.LoggingMiddleware(indexHandler, true), correlationIDOptions, nil))

	port := 8087
	log.Println("Server started on http://localhost:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
