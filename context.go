package tracr

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
)

var DefaultCorrelationIDHeaderName = "x-correlation-id"

// CorrelationIDOptions
// Func to be passed into tracr methods to allow retrieval of the preferred correlation id
// header name, e.g. x-correlation-id, x-tracer-id, x-request-id, etc...
// Note, a func is used here, to allow callers to use whatever method is preferred to
// retrieve the header name, env variables, config file, db, ssm, etc...
type CorrelationIDOptions func() (correlationIDHeaderName string)

// func string here, uses the gofr lib, to generate a uuid (allows us to override during testing)
var correlationIDGenerator = uuid.Must(uuid.NewV4()).String

type contextKey string

var (
	ErrContextIncompatibleType = errors.New("value from context did not match expected type")
	ErrContextValueNotFound    = errors.New("context value not found")
)

// nb: This is the string used to add the correlation id to the go context
// It is NOT the name of the request header (e.g. x-correlation-id), see README, the correlation id
// header name used can be set dynamically
const contextKeyCorrelationID contextKey = "correlation_id"

// ContextWithCID adds the correlation id to the given context.
// Context is immutable so a new context is returned.
func ContextWithCID(ctx context.Context, correlationID string) context.Context {
	if correlationID == "" {
		correlationID = correlationIDGenerator()
	}
	return context.WithValue(ctx, contextKeyCorrelationID, correlationID)
}

// GetCID retrieve existing correlation id from the context.
func GetCID(ctx context.Context) (string, error) {
	correlationID, err := FromContext[string](contextKeyCorrelationID, ctx)
	return correlationID, err
}

// FromContext - Get value from context
func FromContext[T any](key contextKey, ctx context.Context) (T, error) {
	var ret T
	if ctx == nil {
		return ret, errors.New("FromContext() context was null")
	}

	tmp := ctx.Value(key)

	if tmp == nil {
		return ret, ErrContextValueNotFound
	}

	ret, ok := tmp.(T)
	if !ok {
		return ret, ErrContextIncompatibleType
	}
	return ret, nil
}
