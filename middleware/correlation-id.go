package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/micklove/tracr"
	"log"
	"net/http"
)

// MiddlewareCorrelationID - Middleware to retrieve the correlation ID header from an incoming request
// It uses a closure/decorator to inject a func that provides the preferred correlation id header key for the request.
// Note that if the correlation ID header is missing, the middleware creates one and adds it to the context.
func MiddlewareCorrelationID(next http.HandlerFunc, option tracr.CorrelationIDOptions, logger *log.Logger) http.HandlerFunc {
	if logger == nil {
		// no logger provided, use default
		logger = log.New(log.Writer(), "", 0)
	}

	return func(w http.ResponseWriter, r *http.Request) {

		correlationIDHeaderName, err := option.GetCorrelationIDHttpHeaderName()
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

		cidFromCtx, err := tracr.GetCID(ctx)
		w.Header().Set(correlationIDHeaderName, cidFromCtx)
		r = r.WithContext(ctx)
		next(w, r)
	}
}

// MiddlewareCorrelationIDChi - Chi framework version of the MiddlewareCorrelationID
func MiddlewareCorrelationIDChi(option tracr.CorrelationIDOptions, logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		if logger == nil {
			// no logger provided, use default
			logger = log.New(log.Writer(), "", 0)
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			correlationIDHeaderName, err := option.GetCorrelationIDHttpHeaderName()
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

			cidFromCtx, err := tracr.GetCID(ctx)
			w.Header().Set(correlationIDHeaderName, cidFromCtx)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// MiddlewareCorrelationIDGin Gin framework version of the MiddlewareCorrelationIDChi middleware
func MiddlewareCorrelationIDGin(option tracr.CorrelationIDOptions, logger *log.Logger) gin.HandlerFunc {
	if logger == nil {
		// no logger provided, use default
		logger = log.New(log.Writer(), "", 0)
	}

	return func(c *gin.Context) {
		correlationIDHeaderName, err := option.GetCorrelationIDHttpHeaderName()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		correlationID := c.GetHeader(correlationIDHeaderName)
		ctx, err := tracr.ContextWithCID(c.Request.Context(), correlationID, option.CorrelationIDGeneratorFn)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		cidFromCtx, err := tracr.GetCID(ctx)
		c.Header(correlationIDHeaderName, cidFromCtx)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
