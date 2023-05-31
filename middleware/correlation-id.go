package tracr

import (
	"github.com/micklove/tracr"
	"log"
	"net/http"
)

// MiddlewareCorrelationID - Middleware to retrieve the correlation ID header from an incoming request
// It uses a closure to inject a func that provides the preferred correlation id header key for the request.
// Note that if the correlation ID header is missing, the middleware creates one and adds it to the context.
func MiddlewareCorrelationID(fn tracr.CorrelationIDOptions, logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		correlationIDHeaderName := tracr.DefaultCorrelationIDHeaderName

		if logger == nil {
			// no logger provided, use default
			logger = log.New(log.Writer(), "", 0)
		}

		// Execute the function to use the preferred correlation ID header key, e.g. my-cid, request-id, etc... ; if provided
		if fn != nil {
			correlationIDHeaderName = fn()
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			correlationID := r.Header.Get(correlationIDHeaderName)
			ctx := tracr.ContextWithCID(r.Context(), correlationID)
			w.Header().Set(correlationIDHeaderName, correlationID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
