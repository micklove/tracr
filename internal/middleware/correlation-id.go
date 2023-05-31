package tracr

import (
	"github.com/micklove/tracr/internal/tracr"
	"log"
	"net/http"
)

// MiddlewareCorrelationID - Middleware to retrieve the correlation ID header from an incoming request
// It uses a closure/decorator to inject a func that provides the preferred correlation id header key for the request.
// Note that if the correlation ID header is missing, the middleware creates one and adds it to the context.
func MiddlewareCorrelationID(option tracr.CorrelationIDOption, logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		if logger == nil {
			// no logger provided, use default
			logger = log.New(log.Writer(), "", 0)
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			correlationIDHeaderName, err := option.CorrelationIDHttpHeaderFn()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			correlationID := r.Header.Get(correlationIDHeaderName)
			ctx, err := tracr.ContextWithCID(r.Context(), correlationID, option.CorrelationIDGeneratorFn)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			w.Header().Set(correlationIDHeaderName, correlationID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}