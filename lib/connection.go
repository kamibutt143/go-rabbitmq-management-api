package lib

import (
	"fmt"
	"net/url"
)

// connection represents the RabbitMQ connection API client.
type connection struct {
	client RabbitMQAPIClient
}

// ConnectionInterface defines the interface for interacting with RabbitMQ connections.
type ConnectionInterface interface {
	ListConnections() (string, error)
	GetAConnection(connection string) (string, error)
	CloseConnection(connection string) (string, error)
}

// NewConnection creates a new RabbitMQ connection API client with the provided configuration.
// It returns a ConnectionInterface.
func NewConnection(config map[string]interface{}) (ConnectionInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}

	return &connection{
		client,
	}, nil
}

// ListConnections retrieves a list of all RabbitMQ connections.
func (c *connection) ListConnections() (string, error) {
	path := "/api/connections/"
	return c.client.Get(path)
}

// GetAConnection retrieves information about a specific RabbitMQ connection.
func (c *connection) GetAConnection(connection string) (string, error) {
	if connection == "" {
		return "", fmt.Errorf("missing connection parameter")
	}

	path := "/api/connections/" + url.QueryEscape(connection)
	return c.client.Get(path)
}

// CloseConnection closes a specific RabbitMQ connection.
func (c *connection) CloseConnection(connection string) (string, error) {
	if connection == "" {
		return "", fmt.Errorf("missing connection parameter")
	}

	path := "/api/connections/" + url.QueryEscape(connection)
	return c.client.Delete(path)
}
