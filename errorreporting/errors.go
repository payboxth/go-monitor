package errorreporting

import (
	"context"
	"log"

	ger "cloud.google.com/go/errorreporting"
)

var client *ger.Client

// NewClient use to create a error reporting client
func NewClient(ctx context.Context, projectID, serviceName string,) error {
	var err error
	client, err = ger.NewClient(ctx, projectID, ger.Config{
		ServiceName: serviceName,
		OnError: func(err error) {
			log.Printf("Could not log error: %v", err)
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// Close use to close a error reporting client
func Close() error {
	return client.Close()
}

// LogError use to send error
func LogError(err error) {
	client.Report(ger.Entry{Error: err})
}
