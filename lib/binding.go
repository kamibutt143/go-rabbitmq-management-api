// Package lib provides an API client for interacting with RabbitMQ bindings.
package lib

import (
	"net/url"
)

// binding represents a RabbitMQ binding client.
type binding struct {
	client RabbitMQAPIClient
}

// BindingInterface defines the interface for interacting with RabbitMQ bindings.
type BindingInterface interface {
	ListBindings() (string, error)
	ListBindingForAVhost(vhost string) (string, error)
}

// NewBinding creates a new Binding instance with the provided configuration.
func NewBinding(config map[string]interface{}) (BindingInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}

	return &binding{
		client,
	}, nil
}

// ListBindings retrieves a list of all RabbitMQ bindings.
func (b *binding) ListBindings() (string, error) {
	path := "/api/bindings/"
	return b.client.Get(path)
}

// ListBindingForAVhost retrieves a list of bindings for a specific RabbitMQ virtual host.
func (b *binding) ListBindingForAVhost(vhost string) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost)
	return b.client.Get(path)
}
