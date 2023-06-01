package tracr

import (
	"context"
	"net/http"
)

// AddCorrelationIDToRequest - For http clients.
// Adds a CorrelationID, from the given context, as a request Header to the request.
// If the context does not contain an existing CorrelationID, a new one will be generated using the
// correlationIDGeneratorFn provided.
// Note, If context or request are null, new ones will be returned
func AddCorrelationIDToRequest(ctx context.Context, req *http.Request, correlationIDHeaderFn CorrelationIDHeaderFn, correlationIDGeneratorFn CorrelationIDGenerator) (*http.Request, error) {
	correlationID, _ := GetCID(ctx)
	if ctx == nil {
		ctx = context.TODO() // TODO - is this the best way to handle this?
	}

	// Copy the current context, containing the correlation ID (will generator one, if required)
	ctx, err := ContextWithCID(ctx, correlationID, correlationIDGeneratorFn)
	if err != nil {
		return nil, err
	}

	// Correlation ID, existing or newly created, will now be available on the context
	correlationID, _ = GetCID(ctx)

	if req == nil {
		req = &http.Request{}
	}

	req = req.WithContext(ctx)
	correlationIDHeaderName, err := correlationIDHeaderFn()
	if err != nil {
		return nil, err
	}

	if req.Header == nil {
		req.Header = http.Header{}
	}
	req.Header.Set(correlationIDHeaderName, correlationID)
	return req, nil
}
