package tracr

import (
	"context"
	"fmt"
)

// Example service... For each func, the correlation id will be present in the context, since it was added to the context by the middleware

type Service struct {
}

// Echo return the correlation id from the given context
func (s Service) Echo(ctx context.Context) (string, error) {
	correlationID, err := GetCID(ctx)
	if err != nil {
		return "", err
	}

	if len(correlationID) == 0 {
		msg := "correlation id is blank"
		return "", fmt.Errorf(msg)
	}

	return correlationID, nil
}

func NewService() Service {
	return Service{}
}
