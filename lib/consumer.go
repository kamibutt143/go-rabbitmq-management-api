// Package lib provides an API client for interacting with RabbitMQ consumer.
package lib

import (
	"net/url"
)

// consumer represents the Consumer API client.
type consumer struct {
	client RabbitMQAPIClient
}

// ConsumerInterface defines the interface for interacting with RabbitMQ channels.
type ConsumerInterface interface {
	ListConsumers() (string, error)
	ListConsumersForAVhost(vhost string) (string, error)
}

// NewConsumer creates a new Consumer API client with the provided configuration.
// It returns a ConsumerInterface.
func NewConsumer(config map[string]interface{}) (ConsumerInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}

	return &consumer{
		client,
	}, nil
}

// ListConsumers A list of all consumers
func (c *consumer) ListConsumers() (string, error) {
	path := "/api/consumers"
	return c.client.Get(path)
}

// ListConsumersForAVhost A list of all consumers in a given virtual host.
func (c *consumer) ListConsumersForAVhost(vhost string) (string, error) {
	if err := validateParam(vhost, "channel"); err != nil {
		return "", err
	}
	path := "/api/consumers/" + url.QueryEscape(vhost)
	return c.client.Get(path)
}
