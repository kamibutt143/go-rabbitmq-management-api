package lib

import (
	"net/url"
)

// connection represents the RabbitMQ connection API client.
type connection struct {
	client RabbitMQAPIClient
}

// ConnectionInterface defines the interface for interacting with RabbitMQ connections.
type ConnectionInterface interface {
	ListConnections(pagination map[string]interface{}) (string, error)
	GetAConnection(connection string) (string, error)
	CloseConnection(connection string) (string, error)
	ListChannelsForAConnection(connection string) (string, error)
	ListOpenConnectionsForAUser(username string, pagination map[string]interface{}) (string, error)
	DeleteOpenConnectionsForAUser(username string) (string, error)
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

// ListConnections retrieves a list of all RabbitMQ connections. Use pagination parameters to filter connections.
func (c *connection) ListConnections(pagination map[string]interface{}) (string, error) {
	path := "/api/connections/"

	query, err := buildPaginationQuery(pagination)
	if err != nil {
		return "", err
	}
	if query != "" {
		path += query
	}

	return c.client.Get(path)
}

// GetAConnection retrieves information about a specific RabbitMQ connection.
func (c *connection) GetAConnection(connection string) (string, error) {
	if err := validateParam(connection, "connection"); err != nil {
		return "", err
	}

	path := "/api/connections/" + url.QueryEscape(connection)
	return c.client.Get(path)
}

// CloseConnection closes a specific RabbitMQ connection.
func (c *connection) CloseConnection(connection string) (string, error) {
	if err := validateParam(connection, "connection"); err != nil {
		return "", err
	}

	path := "/api/connections/" + url.QueryEscape(connection)
	return c.client.Delete(path)
}

// ListChannelsForAConnection List of all channels for a given connection.
func (c *connection) ListChannelsForAConnection(connection string) (string, error) {
	if err := validateParam(connection, "connection"); err != nil {
		return "", err
	}

	path := "/api/connections/" + url.QueryEscape(connection) + "/channels"
	return c.client.Get(path)
}

// ListOpenConnectionsForAUser A list of all open connections for a specific username. Use pagination parameters to filter connections.
func (c *connection) ListOpenConnectionsForAUser(username string, pagination map[string]interface{}) (string, error) {
	if err := validateParam(username, "username"); err != nil {
		return "", err
	}
	path := "/api/connections/username/" + url.QueryEscape(username)
	query, err := buildPaginationQuery(pagination)
	if err != nil {
		return "", err
	}
	if query != "" {
		path += query
	}

	return c.client.Get(path)
}

// DeleteOpenConnectionsForAUser Delete a resource will close all the connections for a username.
func (c *connection) DeleteOpenConnectionsForAUser(username string) (string, error) {
	if err := validateParam(username, "username"); err != nil {
		return "", err
	}

	path := "/api/connections/username/" + url.QueryEscape(username)
	return c.client.Delete(path)
}
