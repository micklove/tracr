package middleware

import (
	"golang.org/x/exp/slog"
	"net/http"
	"time"
)

// TODO - get the correlation id from the response context and log it
// See example at https://go.dev/play/p/ia8aJr1RiQB?v=gotip or https://github.com/sollniss/ctxkey

// LoggingMiddleware logs each HTTP response, method, uri, status, duration and body (optional)
func LoggingMiddleware(next http.HandlerFunc, logBody bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Wrap the original ResponseWriter to capture the status code
		res := &ResponseRecorder{ResponseWriter: w, StatusCode: http.StatusOK}

		// Call the original handler with the new ResponseWriter
		next(res, r)

		duration := time.Since(startTime)
		if logBody {
			slog.InfoContext(r.Context(), "http response body",
				slog.String("method", r.Method),
				slog.String("path", r.RequestURI),
				slog.Duration("duration", duration),
				slog.Int("http_status", res.StatusCode),
				slog.String("body", string(res.Body)),
			)
		} else {
			slog.InfoContext(r.Context(), "http response",
				slog.String("method", r.Method),
				slog.String("path", r.RequestURI),
				slog.Duration("duration", duration),
				slog.Int("http_status", res.StatusCode),
			)
		}
	}
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body
	return rec.ResponseWriter.Write(body)
}
