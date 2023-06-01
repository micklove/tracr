package main

import (
	"fmt"
	"github.com/go-chi/chi"
	mid "github.com/micklove/tracr/internal/middleware"
	"github.com/micklove/tracr/internal/tracr"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// Create the correlation ID options, in this case, we will simply hard code in the funcs that will
	// retrieve the correlation ID http header name and the correlation ID value. This allows the caller to
	// use whatever strategy they want, to generate the correlation ID e.g. DB Sequence, UUID, GUID, etc... and / or
	// to use their preferred strategy for naming the correlation id http header (e.g. from env var, config , ssm, etc..)
	cidHttpHeaderName := "my-trace-header"
	correlationIDOptions := tracr.CorrelationIDOptions{
		CorrelationIDHttpHeaderFn: func() (string, error) { return cidHttpHeaderName, nil },
		CorrelationIDGeneratorFn:  func() (string, error) { return "b63a65cc-20fa-4b17-97ad-b796bdb6d338", nil },
	}

	// register the correlation ID middleware
	r.Use(mid.MiddlewareCorrelationID(correlationIDOptions, nil))

	// create a handler named hello
	hello := func(w http.ResponseWriter, r *http.Request) {
		// get my-trace-header from the context
		cid, err := tracr.GetCID(r.Context())
		if err != nil {
			log.Printf("GetCID() = returned error %v", err)
		}

		// get the my-trace-header from the request
		headerCid := r.Header.Get(cidHttpHeaderName)
		log.Printf("Header %s = %s", cidHttpHeaderName, headerCid)
		log.Printf("Context CID = %s", cid)

		// echo the correlation id header and value in the response
		w.Header().Set(cidHttpHeaderName, cid)
		w.Write([]byte(fmt.Sprintf("Hello World!\ncorrelation id header in  = %s\ncorrelation id header out = %s\n", headerCid, cid)))
	}

	// register the handlers
	r.Get("/hello", hello)

	// TODO

	// start the server
	log.Fatal(http.ListenAndServe(":3000", r))
}
