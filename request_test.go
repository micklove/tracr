package tracr

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestAddCorrelationIDToRequest_null_request_creates_request_with_correlation_id(t *testing.T) {
	wantCorrelationID := "abcdef-ghijkl-1234"
	correlationIDGenerator = func() string { return wantCorrelationID } // override the gotCorrelationID generator for testing blank values
	expectedCorrelationIDHeader := "my-request-id"
	overrideHeaderFunc := func() string { return expectedCorrelationIDHeader }

	//SUT
	req := AddCorrelationIDToRequest(nil, nil, overrideHeaderFunc)

	gotCorrelationID := req.Header.Get(expectedCorrelationIDHeader)
	if gotCorrelationID != wantCorrelationID {
		t.Errorf("want [%s], got[%s]", wantCorrelationID, gotCorrelationID)
	}
}

func TestAddCorrelationIDToRequest_null_context_creates_request_with_correlation_id(t *testing.T) {
	expectedUrl := "/blah"
	req, _ := http.NewRequest(http.MethodGet, expectedUrl, nil)

	wantCorrelationID := "abcdef-ghijkl-12345678"
	correlationIDGenerator = func() string { return wantCorrelationID } // override the gotCorrelationID generator for testing blank values
	expectedCorrelationIDHeader := "my-request-id"
	overrideHeaderFunc := func() string { return expectedCorrelationIDHeader }

	//SUT
	gotReq := AddCorrelationIDToRequest(nil, req, overrideHeaderFunc)

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

//func TestAddCorrelationIDToRequest(t *testing.T) {
//	type args struct {
//		ctx context.Context
//		req func() *http.Request
//	}
//	tests := []struct {
//		name        string
//		args        args
//		wantContext context.Context
//		wantRequest *http.Request
//	}{
//		{
//			name: "correlation id header added when none found in context",
//			args: args{
//				ctx: context.WithValue(context.TODO(), contextKeyCorrelationID, 1),
//				req: func() *http.Request {
//					return nil
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := AddCorrelationIDToRequest(tt.args.ctx, tt.args.req, nil)
//			if !reflect.DeepEqual(got, tt.wantContext) {
//				t.Errorf("AddCorrelationIDToRequest() got = %v, wantContext %v", got, tt.wantContext)
//			}
//			if !reflect.DeepEqual(got1, tt.wantRequest()) {
//				t.Errorf("AddCorrelationIDToRequest() got1 = %v, wantContext %v", got1, tt.wantRequest)
//			}
//		})
//	}
//}
