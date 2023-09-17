// Package lib provides an API client for interacting with RabbitMQ node.
package lib

import (
	"net/url"
)

// node represents a RabbitMQ node client.
type node struct {
	client RabbitMQAPIClient
}

// NodeInterface defines the interface for interacting with RabbitMQ node.
type NodeInterface interface {
	ListNodes() (string, error)
	GetANode(node string, options map[string]interface{}) (string, error)
}

// NewNode creates a new node instance with the provided configuration.
func NewNode(config map[string]interface{}) (NodeInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}

	return &node{
		client,
	}, nil
}

// ListNodes retrieves a list of all RabbitMQ nodes.
func (n *node) ListNodes() (string, error) {
	path := "/api/nodes"
	return n.client.Get(path)
}

// GetANode An individual node in the RabbitMQ cluster.
// Add "?memory=true" to get memory statistics, and "?binary=true" to get a breakdown of binary memory use (may be expensive if there are many small binaries in the system).
func (n *node) GetANode(node string, options map[string]interface{}) (string, error) {
	if err := validateParam(node, "node"); err != nil {
		return "", err
	}

	path := "/api/nodes/" + url.QueryEscape(node)
	query, err := buildQuery(options)
	if err != nil {
		return "", err
	}
	if query != "" {
		path += query
	}

	return n.client.Get(path)
}
