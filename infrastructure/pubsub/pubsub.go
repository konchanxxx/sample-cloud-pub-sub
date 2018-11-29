package pubsub

import (
	ps "cloud.google.com/go/pubsub"
	"context"

	"github.com/rexitorg/sample-cloud-pub-sub/infrastructure/environment"
)

type PubSubService struct {
	Client *ps.Client
}

func NewPubSubService(ctx context.Context) (*PubSubService, error) {
	e := environment.NewEnvironments()
	projectID := e.GCP.ProjectID
	c, err := ps.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	s := &PubSubService{
		Client: c,
	}

	return s, nil
}
