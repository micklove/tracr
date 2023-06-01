package tracr

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestAddCorrelationIDToRequest_null_request_creates_request_with_correlation_id(t *testing.T) {
	wantCorrelationID := "abcdef-ghijkl-1234"
	correlationIDGenerator := func() (string, error) { return wantCorrelationID, nil } // override the gotCorrelationID generator for testing blank values
	expectedCorrelationIDHeader := "my-request-id"
	overrideHeaderFunc := func() (string, error) { return expectedCorrelationIDHeader, nil }

	//SUT
	req, _ := AddCorrelationIDToRequest(nil, nil, overrideHeaderFunc, correlationIDGenerator)

	gotCorrelationID := req.Header.Get(expectedCorrelationIDHeader)
	if gotCorrelationID != wantCorrelationID {
		t.Errorf("want [%s], got[%s]", wantCorrelationID, gotCorrelationID)
	}
}

func TestAddCorrelationIDToRequest_null_context_creates_request_with_correlation_id(t *testing.T) {
	expectedUrl := "/blah"
	req, _ := http.NewRequest(http.MethodGet, expectedUrl, nil)

	wantCorrelationID := "abcdef-ghijkl-12345678"
	correlationIDGenerator := func() (string, error) { return wantCorrelationID, nil } // override the gotCorrelationID generator for testing blank values
	expectedCorrelationIDHeader := "my-request-id"
	overrideHeaderFunc := func() (string, error) { return expectedCorrelationIDHeader, nil }

	//SUT
	gotReq, _ := AddCorrelationIDToRequest(nil, req, overrideHeaderFunc, correlationIDGenerator)

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
