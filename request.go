package tracr

import (
	"context"
	"net/http"
)

// AddCorrelationIDToRequest - For http clients.
// Add a CorrelationID, as a request Header to the given request. The correlation id is upserted from the context.
// nb: If context or request are null, new ones will be returned
func AddCorrelationIDToRequest(ctx context.Context, req *http.Request, fn CorrelationIDOptions) *http.Request {
	correlationID, _ := GetCID(ctx)
	if ctx == nil {
		ctx = context.TODO()
	}
	if correlationID == "" {
		correlationID = correlationIDGenerator()
	}
	// Copy the current context, add the correlation ID
	ctx = ContextWithCID(ctx, correlationID)
	if req == nil {
		req = &http.Request{}
	}
	req = req.WithContext(ctx)
	correlationIDHeaderName := DefaultCorrelationIDHeaderName
	if fn != nil {
		correlationIDHeaderName = fn()
	}

	if req.Header == nil {
		req.Header = http.Header{}
	}
	req.Header.Set(correlationIDHeaderName, correlationID)
	return req
}
