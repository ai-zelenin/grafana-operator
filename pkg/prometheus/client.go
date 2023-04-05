package prometheus

import (
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type Client struct {
	api.Client
	v1.API
}

func NewClient(baseUrl string) (*Client, error) {
	client, err := api.NewClient(api.Config{
		Address: baseUrl,
	})
	if err != nil {
		return nil, err
	}
	v1api := v1.NewAPI(client)
	return &Client{
		Client: client,
		API:    v1api,
	}, nil
}
