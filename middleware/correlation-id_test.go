package tracr

import (
	"github.com/gofrs/uuid"
	"github.com/micklove/tracr"
	"github.com/micklove/tracr/internal/tracrtest"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func Test_correlation_id_middleware_uses_given_header(t *testing.T) {
	expectedCorrelationIDHeader := "my-trace-header"
	wantCid := "b63a65cc-20fa-4b17-97ad-b796bdb6d338"
	rec := httptest.NewRecorder()
	header := http.Header{}
	header.Set(expectedCorrelationIDHeader, wantCid)
	req := tracrtest.MockHttpRequest("GET", "/example/", nil, url.Values{
		"blah": {"blahblah"},
	}, header)

	// Set our expectations that the correlation id will be available in the future context
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got, err := tracr.GetCID(r.Context())
		if err != nil {
			t.Errorf("GetCID() = returned error %v", err)
		}
		if got != wantCid {
			t.Errorf("MiddlewareCorrelationID() = %v, want %v", got, wantCid)
		}
		w.Write([]byte("OK"))
	})

	// Execute our middleware
	overrideHeaderFunc := func() string { return expectedCorrelationIDHeader }
	MiddlewareCorrelationID(overrideHeaderFunc, nil)(next).ServeHTTP(rec, req)
}

func Test_correlation_id_middleware_load_header_name_from_env(t *testing.T) {
	envKey := "MY_PREFERRED_CORRELATION_HEADER_NAME"
	wantCorrelationIDHeader := "my-trace-header"
	os.Setenv(envKey, wantCorrelationIDHeader)

	wantCid := uuid.Must(uuid.NewV4()).String()
	rec := httptest.NewRecorder()
	header := http.Header{}
	header.Set(wantCorrelationIDHeader, wantCid)
	req := tracrtest.MockHttpRequest("GET", "/example/", nil, nil, header)

	// Set our expectations that the correlation id will be available in the future context
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got, err := tracr.GetCID(r.Context())
		if err != nil {
			t.Errorf("GetCID() = returned error %v", err)
		}
		if got != wantCid {
			t.Errorf("MiddlewareCorrelationID() = %v, want %v", got, wantCid)
		}
		w.Write([]byte("OK"))
	})

	// Execute our middleware, pull the correlation id header key from the env
	overrideHeaderFunc := func() string {
		return os.Getenv(envKey)
	}
	MiddlewareCorrelationID(overrideHeaderFunc, nil)(next).ServeHTTP(rec, req)
}
