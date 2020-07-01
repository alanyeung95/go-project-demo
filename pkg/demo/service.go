package demo

import (
	"context"
	"os"

	"github.com/getsentry/sentry-go"
)

// Service interface
type Service interface {
	DemoError(ctx context.Context) error
}

type service struct {
}

// NewService start the new service
func NewService() (Service, error) {
	return &service{}, nil
}

func (s *service) DemoError(ctx context.Context) error {
	_, err := os.Open("not_existing_file.go")
	if err != nil {
		sentry.CaptureMessage(err.Error())
		return err
	}
	return nil
}
