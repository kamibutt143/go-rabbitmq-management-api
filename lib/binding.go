// Package lib provides an API client for interacting with RabbitMQ bindings.
package lib

import (
	"encoding/json"
	"fmt"
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
	ListBindingForAVhostExchangeAndQueue(vhost, exchange, queue string) (string, error)
	CreateBindingForAVhostExchangeAndQueue(vhost, exchange, queue string, options map[string]string) (string, error)
	GetBindingForAVhostExchangeAndQueue(vhost, exchange, queue, props string) (string, error)
	ListBindingsForAVhostBetweenTwoExchanges(vhost, source, destination string) (string, error)
	CreateBindingForAVhostBetweenTwoExchanges(vhost, source, destination string, options map[string]string) (string, error)
	GetBindingForAVhostBetweenTwoExchanges(vhost, source, destination, props string) (string, error)
	DeleteBindingForAVhostBetweenTwoExchanges(vhost, source, destination, props string) (string, error)
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

// ListBindings A list of all bindings.
func (b *binding) ListBindings() (string, error) {
	path := "/api/bindings/"
	return b.client.Get(path)
}

// ListBindingForAVhost A list of all bindings in a given virtual host.
func (b *binding) ListBindingForAVhost(vhost string) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost)
	return b.client.Get(path)
}

// ListBindingForAVhostExchangeAndQueue A list of all bindings in a given virtual host.
func (b *binding) ListBindingForAVhostExchangeAndQueue(vhost, exchange, queue string) (string, error) {
	params := map[string]string{
		"vhost":    vhost,
		"exchange": exchange,
		"queue":    queue,
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost) + "/e/" + url.QueryEscape(exchange) + "/q/" + url.QueryEscape(queue)
	return b.client.Get(path)
}

// CreateBindingForAVhostExchangeAndQueue To create a new binding, POST to this URI.
// Request body should be a JSON object optionally containing two fields, routing_key (a string) and arguments (a map of optional arguments):
// {"routing_key":"my_routing_key", "arguments":{"x-arg": "value"}}
// All keys are optional. The response will contain a Location header telling you the URI of your new binding.
func (b *binding) CreateBindingForAVhostExchangeAndQueue(vhost, exchange, queue string, options map[string]string) (string, error) {
	params := map[string]string{
		"vhost":    vhost,
		"exchange": exchange,
		"queue":    queue,
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	// Marshal the map into a JSON string
	jsonData, err := json.Marshal(options)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost) + "/e/" + url.QueryEscape(exchange) + "/q/" + url.QueryEscape(queue)
	return b.client.Post(path, string(jsonData))
}

// GetBindingForAVhostExchangeAndQueue An individual binding between an exchange and a queue.
// The props part of the URI is a "name" for the binding composed of its routing key and a hash of its arguments.
// props is the field named "properties_key" from a bindings listing response.
func (b *binding) GetBindingForAVhostExchangeAndQueue(vhost, exchange, queue, props string) (string, error) {
	params := map[string]string{
		"vhost":    vhost,
		"exchange": exchange,
		"queue":    queue,
		"props":    props,
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost) + "/e/" + url.QueryEscape(exchange) + "/q/" + url.QueryEscape(queue) + "/" + url.QueryEscape(props)
	return b.client.Get(path)
}

// DeleteBindingForAVhostExchangeAndQueue Delete an individual binding between an exchange and a queue.
func (b *binding) DeleteBindingForAVhostExchangeAndQueue(vhost, exchange, queue, props string) (string, error) {
	params := map[string]string{
		"vhost":    vhost,
		"exchange": exchange,
		"queue":    queue,
		"props":    props,
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost) + "/e/" + url.QueryEscape(exchange) + "/q/" + url.QueryEscape(queue) + "/" + url.QueryEscape(props)
	return b.client.Delete(path)
}

// ListBindingsForAVhostBetweenTwoExchanges A list of all bindings between two exchanges
func (b *binding) ListBindingsForAVhostBetweenTwoExchanges(vhost, source, destination string) (string, error) {
	params := map[string]string{
		"vhost":    vhost,
		"exchange": source,
		"queue":    destination,
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost) + "/e/" + url.QueryEscape(source) + "/e/" + url.QueryEscape(destination)
	return b.client.Get(path)
}

// CreateBindingForAVhostBetweenTwoExchanges To create a new binding, POST to this URI.
// Request body should be a JSON object optionally containing two fields, routing_key (a string) and arguments (a map of optional arguments):
// {"routing_key":"my_routing_key", "arguments":{"x-arg": "value"}}
// All keys are optional. The response will contain a Location header telling you the URI of your new binding.
func (b *binding) CreateBindingForAVhostBetweenTwoExchanges(vhost, source, destination string, options map[string]string) (string, error) {
	params := map[string]string{
		"vhost":       vhost,
		"source":      source,
		"destination": destination,
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	// Marshal the map into a JSON string
	jsonData, err := json.Marshal(options)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost) + "/e/" + url.QueryEscape(source) + "/e/" + url.QueryEscape(destination)
	return b.client.Post(path, string(jsonData))
}

// GetBindingForAVhostBetweenTwoExchanges An individual binding between two exchanges.
// The props part of the URI is a "name" for the binding composed of its routing key and a hash of its arguments.
// props is the field named "properties_key" from a bindings listing response.
func (b *binding) GetBindingForAVhostBetweenTwoExchanges(vhost, source, destination, props string) (string, error) {
	params := map[string]string{
		"vhost":    vhost,
		"exchange": source,
		"queue":    destination,
		"props":    props,
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost) + "/e/" + url.QueryEscape(source) + "/e/" + url.QueryEscape(destination) + "/" + url.QueryEscape(props)
	return b.client.Get(path)
}

// DeleteBindingForAVhostBetweenTwoExchanges Delete an individual binding between two exchanges
func (b *binding) DeleteBindingForAVhostBetweenTwoExchanges(vhost, source, destination, props string) (string, error) {
	params := map[string]string{
		"vhost":       vhost,
		"source":      source,
		"destination": destination,
		"props":       props,
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	path := "/api/bindings/" + url.QueryEscape(vhost) + "/e/" + url.QueryEscape(source) + "/e/" + url.QueryEscape(destination) + "/" + url.QueryEscape(props)
	return b.client.Delete(path)
}
