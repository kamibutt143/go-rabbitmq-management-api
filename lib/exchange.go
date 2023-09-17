// Package lib provides an API client for interacting with RabbitMQ exchanges.
package lib

import (
	"encoding/json"
	"net/url"
)

// exchange represents the Exchange API client.
type exchange struct {
	client RabbitMQAPIClient
}

// ExchangeInterface defines the interface for interacting with RabbitMQ exchanges and bindings.
type ExchangeInterface interface {
	ListExchanges(pagination map[string]interface{}) (string, error)
	ListExchangesForAVhost(vhost string, pagination map[string]interface{}) (string, error)
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

// ListExchanges A list of all exchanges. Use pagination parameters to filter exchanges.
func (e *exchange) ListExchanges(pagination map[string]interface{}) (string, error) {
	path := "/api/exchanges"
	query, err := buildPaginationQuery(pagination)
	if err != nil {
		return "", err
	}
	if query != "" {
		path += query
	}

	return e.client.Get(path)
}

// ListExchangesForAVhost A list of all exchanges in a given virtual host. Use pagination parameters to filter exchanges.
func (e *exchange) ListExchangesForAVhost(vhost string, pagination map[string]interface{}) (string, error) {
	if err := validateExchangeParams(vhost, "", ""); err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost)
	query, err := buildPaginationQuery(pagination)
	if err != nil {
		return "", err
	}
	if query != "" {
		path += query
	}
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

// GetBindingsForSource A list of all bindings in which a given exchange is the source.
func (e *exchange) GetBindingsForSource(vhost string, exchange string) (string, error) {
	if err := validateExchangeParams(vhost, exchange, ""); err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(exchange) + "/bindings/source"
	return e.client.Get(path)
}

// GetBindingsForDestination A list of all bindings in which a given exchange is the destination.
func (e *exchange) GetBindingsForDestination(vhost string, exchange string) (string, error) {
	if err := validateExchangeParams(vhost, exchange, ""); err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(exchange) + "/bindings/destination"
	return e.client.Get(path)
}

// PublishMessage sends a message to the specified exchange in the given virtual host (vhost).
// You will need a body looking something like:
// {"properties":{},"routing_key":"my key","payload":"my body","payload_encoding":"string"}
func (e *exchange) PublishMessage(vhost string, exchange string, options map[string]string) (string, error) {
	if err := validateExchangeParams(vhost, exchange, ""); err != nil {
		return "", err
	}
	if err := validateParam(options["properties"], "properties"); err != nil {
		return "", err
	}
	if err := validateParam(options["routing_key"], "routing_key"); err != nil {
		return "", err
	}
	if err := validateParam(options["payload"], "payload"); err != nil {
		return "", err
	}
	if err := validateParam(options["payload_encoding"], "payload_encoding"); err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(options)
	if err != nil {
		return "", err
	}

	path := "/api/exchanges/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(exchange) + "/publish"
	return e.client.Post(path, string(jsonData))
}
