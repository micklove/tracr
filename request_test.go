package tracr

import (
	"context"
	"github.com/gofrs/uuid"
	"net/http"
	"reflect"
	"testing"
)

func TestAddCorrelationIDToRequest_null_request_creates_request_with_correlation_id(t *testing.T) {
	wantCorrelationID := "abcdef-ghijkl-1234"
	cIDGeneratorFn := func() (string, error) { return wantCorrelationID, nil } // override the gotCorrelationID generator for testing blank values
	expectedCorrelationIDHeader := "my-request-id"
	overrideHeaderFunc := func() (string, error) { return expectedCorrelationIDHeader, nil }

	//SUT
	req, _ := AddCorrelationIDToRequest(nil, nil, CorrelationIDOptions{
		CorrelationIDGeneratorFn:  cIDGeneratorFn,
		CorrelationIDHttpHeaderFn: overrideHeaderFunc,
	})

	gotCorrelationID := req.Header.Get(expectedCorrelationIDHeader)
	if gotCorrelationID != wantCorrelationID {
		t.Errorf("want [%s], got[%s]", wantCorrelationID, gotCorrelationID)
	}
}

func TestAddCorrelationIDToRequest_null_context_creates_request_with_correlation_id(t *testing.T) {
	expectedUrl := "/blah"
	req, _ := http.NewRequest(http.MethodGet, expectedUrl, nil)

	wantCorrelationID := "abcdef-ghijkl-12345678"
	cIDGenerator := func() (string, error) { return wantCorrelationID, nil } // override the gotCorrelationID generator for testing blank values
	expectedCorrelationIDHeader := "my-request-id"
	overrideHeaderFunc := func() (string, error) { return expectedCorrelationIDHeader, nil }

	//SUT
	gotReq, _ := AddCorrelationIDToRequest(nil, req, CorrelationIDOptions{
		CorrelationIDGeneratorFn:  cIDGenerator,
		CorrelationIDHttpHeaderFn: overrideHeaderFunc,
	})

	gotCorrelationID := gotReq.Header.Get(expectedCorrelationIDHeader)
	if gotCorrelationID != wantCorrelationID {
		t.Errorf("want [%s], got[%s]", wantCorrelationID, gotCorrelationID)
	}

	//Ensure we haven't lost any of the original request properties
	if gotReq.URL.String() != expectedUrl {
		t.Errorf("req url, want [%s], got[%s]", expectedUrl, gotReq.URL.String())
	}

	wantCtx := context.WithValue(context.TODO(), contextKeyCorrelationID, wantCorrelationID)
	if !reflect.DeepEqual(gotReq.Context(), wantCtx) {
		t.Errorf("AddCorrelationIDToRequest() got = %v, wantContext %v", gotReq.Context(), wantCtx)
	}

}

func TestAddCorrelationIDToRequest_null_header_name_func_uses_default(t *testing.T) {
	expectedUrl := "/blah"
	req, _ := http.NewRequest(http.MethodGet, expectedUrl, nil)

	wantCorrelationID := "abcdef-ghijkl-12345678"
	cIDGenerator := func() (string, error) { return wantCorrelationID, nil } // override the gotCorrelationID generator for testing blank values

	//SUT
	gotReq, _ := AddCorrelationIDToRequest(nil, req, CorrelationIDOptions{
		CorrelationIDGeneratorFn:  cIDGenerator,
		CorrelationIDHttpHeaderFn: nil,
	})

	gotCorrelationID := gotReq.Header.Get("x-correlation-id")
	if gotCorrelationID != wantCorrelationID {
		t.Errorf("want [%s], got[%s]", wantCorrelationID, gotCorrelationID)
	}

	//Ensure we haven't lost any of the original request properties
	if gotReq.URL.String() != expectedUrl {
		t.Errorf("req url, want [%s], got[%s]", expectedUrl, gotReq.URL.String())
	}

	wantCtx := context.WithValue(context.TODO(), contextKeyCorrelationID, wantCorrelationID)
	if !reflect.DeepEqual(gotReq.Context(), wantCtx) {
		t.Errorf("AddCorrelationIDToRequest() got = %v, wantContext %v", gotReq.Context(), wantCtx)
	}

}
func TestAddCorrelationIDToRequest_null_nil_correlation_id_generator_func_uses_default(t *testing.T) {
	expectedUrl := "/blah"
	req, _ := http.NewRequestWithContext(context.TODO(), http.MethodGet, expectedUrl, nil)

	correlationIDHttpHeaderFn := func() (string, error) { return "x-myrequest-id", nil } // override the gotCorrelationID generator for testing blank values

	//SUT
	gotReq, _ := AddCorrelationIDToRequest(nil, req, CorrelationIDOptions{
		CorrelationIDGeneratorFn:  nil,
		CorrelationIDHttpHeaderFn: correlationIDHttpHeaderFn,
	})

	// get the randomly generated cid, from the request header
	wantCorrelationID := gotReq.Header.Get("x-myrequest-id")
	gotCorrelationID, _ := GetCID(gotReq.Context())

	gotCIDAsUUID, _ := uuid.FromString(gotCorrelationID)

	if gotCIDAsUUID == uuid.Nil {
		t.Errorf("wanted a non-nil , random, uuid")
	}

	// does the context match the http header?
	if wantCorrelationID != gotCorrelationID {
		t.Errorf("want [%s], got[%s]", wantCorrelationID, gotCorrelationID)
	}

	//Ensure we haven't lost any of the original request properties
	if gotReq.URL.String() != expectedUrl {
		t.Errorf("req url, want [%s], got[%s]", expectedUrl, gotReq.URL.String())
	}
}
