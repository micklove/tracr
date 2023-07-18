package tracr

import "github.com/gofrs/uuid"

// TODO - Add default http error handler func

// Add default http error handler func to the struct below

// CorrelationIDOptions
// Struct to be passed into tracr methods to allow retrieval of the preferred correlation id
// http header name, e.g. x-correlation-id, x-tracer-id, x-request-id, etc... and the actual correlation id
// Note, a struct, with funcs is used here, to allow callers to use whatever method / strategy is preferred to
// retrieve the header name, env variables, config file, db, ssm, etc... and the correlation id that will be generated
// if no correlation ID was provided in the request.
type CorrelationIDOptions struct {
	CorrelationIDHttpHeaderFn CorrelationIDHeaderFn
	CorrelationIDGeneratorFn  CorrelationIDGenerator
}

// DefaultCorrelationIDHeaderName  If no default correlation id header strategy is provided
// in CorrelationIDOptions, tracr uses the one below
var DefaultCorrelationIDHeaderName = func() (string, error) { return "x-correlation-id", nil }

// If no correlation id strategy is provided in CorrelationIDOptions, tracr uses the gofr
// uuid lib, to generate a uuid.
var correlationIDGenerator = func() (string, error) { return uuid.Must(uuid.NewV4()).String(), nil }

type CorrelationIDGenerator func() (correlationID string, err error)

// CorrelationIDHeaderFn - allows users of the lib to use their own correlation id http header
// name strategy e.g. x-correlation-id, trace-id, request-id, etc..
// Allows clients to load the http header name from the ENV, config file, etc...
type CorrelationIDHeaderFn func() (correlationIDHeaderName string, err error)

func (o *CorrelationIDOptions) GetCorrelationIDHttpHeaderName() (string, error) {
	if o.CorrelationIDHttpHeaderFn != nil {
		return o.CorrelationIDHttpHeaderFn()
	}
	return DefaultCorrelationIDHeaderName()
}

// getCorrelationID - returns a correlation ID
// Allows users of the lib to use their own correlation id generation strategy
// e.g. Database, different library, etc...
func (o *CorrelationIDOptions) getCorrelationID() (string, error) {
	if o.CorrelationIDGeneratorFn != nil {
		return o.CorrelationIDGeneratorFn()
	}
	return correlationIDGenerator()
}
