package tracr

import (
	"context"
	"errors"
)

// TODO - remove

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
// If not correlation ID exists, a new one is generated, using the provided correlation id generator.
func ContextWithCID(ctx context.Context, correlationID string, cidGenerator CorrelationIDGenerator) (context.Context, error) {

	cid := correlationID
	var err error

	if cid == "" {
		cid, err = cidGenerator()
	}
	return context.WithValue(ctx, contextKeyCorrelationID, cid), err
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
