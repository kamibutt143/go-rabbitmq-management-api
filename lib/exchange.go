// Package lib provides an API client for interacting with RabbitMQ exchanges.
package lib

import (
	"encoding/json"
	"errors"
	"net/url"
)

// exchange represents the Exchange API client.
type exchange struct {
	client RabbitMQAPIClient
}

// ExchangeInterface defines the interface for interacting with RabbitMQ exchanges and bindings.
type ExchangeInterface interface {
	ListExchanges() (string, error)
	ListExchangesForAVhost(vhost string) (string, error)
	GetAExchange(vhost string, exchange string) (string, error)
	CreateExchange(vhost string, exchange string, exchangeType string, options map[string]string) (string, error)
	DeleteExchange(vhost string, exchange string) (string, error)
	GetBindingsForSource(vhost string, exchange string) (string, error)
	GetBindingsForDestination(vhost string, exchange string) (string, error)
	PublishMessage(vhost string, exchange string, options map[string]string) (string, error)
}

// NewExchange creates a new Exchange API client with the provided configuration.
// It returns an ExchangeInterface.
func NewExchange(config map[string]interface{}) (ExchangeInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}

	return &exchange{
		client,
	}, nil
}

// ListExchanges retrieves a list of all exchanges.
func (e *exchange) ListExchanges() (string, error) {
	return e.client.Get("/api/exchanges")
}

// ListExchangesForAVhost retrieves a list of exchanges for a specific virtual host.
func (e *exchange) ListExchangesForAVhost(vhost string) (string, error) {
	if err := validateExchangeParams(vhost, "", ""); err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost)
	return e.client.Get(path)
}

// GetAExchange retrieves information about a specific exchange.
func (e *exchange) GetAExchange(vhost string, exchange string) (string, error) {
	if err := validateExchangeParams(vhost, exchange, ""); err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(exchange)
	return e.client.Get(path)
}

// CreateExchange creates a new exchange with the given parameters and options.
func (e *exchange) CreateExchange(vhost string, exchange string, exchangeType string, options map[string]string) (string, error) {
	if err := validateExchangeParams(vhost, exchange, exchangeType); err != nil {
		return "", err
	}

	options["type"] = exchangeType
	jsonData, err := json.Marshal(options)
	if err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(exchange)
	return e.client.Put(path, string(jsonData))
}

// DeleteExchange deletes a specific exchange.
func (e *exchange) DeleteExchange(vhost string, exchange string) (string, error) {
	if err := validateExchangeParams(vhost, exchange, ""); err != nil {
		return "", err
	}
	path := "/api/exchanges/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(exchange)
	return e.client.Delete(path)
}

// GetBindingsForSource retrieves bindings for which the specified exchange is the source.
func (e *exchange) GetBindingsForSource(vhost string, exchange string) (string, error) {
	if err := validateExchangeParams(vhost, exchange, ""); err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(exchange) + "/bindings/source"
	return e.client.Get(path)
}

// GetBindingsForDestination retrieves bindings for which the specified exchange is the destination.
func (e *exchange) GetBindingsForDestination(vhost string, exchange string) (string, error) {
	if err := validateExchangeParams(vhost, exchange, ""); err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(exchange) + "/bindings/destination"
	return e.client.Get(path)
}

// PublishMessage sends a message to the specified exchange in the given virtual host (vhost).
// Returns a string response and an error if the request fails or if any required parameters are missing.
func (e *exchange) PublishMessage(vhost string, exchange string, options map[string]string) (string, error) {
	if err := validateExchangeParams(vhost, exchange, ""); err != nil {
		return "", err
	}

	if options["properties"] == "" {
		return "", errors.New("missing properties parameter")
	}
	if options["routing_key"] == "" {
		return "", errors.New("missing routing_key parameter")
	}
	if options["payload"] == "" {
		return "", errors.New("missing payload parameter")
	}
	if options["payload_encoding"] == "" {
		return "", errors.New("missing payload_encoding parameter")
	}

	jsonData, err := json.Marshal(options)
	if err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(exchange) + "/bindings/destination"
	return e.client.Post(path, string(jsonData))
}
