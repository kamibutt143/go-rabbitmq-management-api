// Package lib provides an API client for interacting with RabbitMQ node.
package lib

import (
	"fmt"
	"net/url"
)

// node represents a RabbitMQ node client.
type node struct {
	client RabbitMQAPIClient
}

// NodeInterface defines the interface for interacting with RabbitMQ node.
type NodeInterface interface {
	ListNodes() (string, error)
	GetANode(node string) (string, error)
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

// GetANode retrieve a node information
func (n *node) GetANode(node string) (string, error) {
	if node == "" {
		return "", fmt.Errorf("missing node parameter")
	}

	path := "/api/nodes/" + url.QueryEscape(node)
	return n.client.Get(path)
}
