package lib

import (
	"encoding/json"
	"net/url"
)

// definition represents the RabbitMQ definition API client.
type definition struct {
	client RabbitMQAPIClient
}

// DefinitionInterface defines the interface for interacting with RabbitMQ definitions.
type DefinitionInterface interface {
	ListDefinitions() (string, error)
	CreateDefinition(definition string) (string, error)
	ListDefinitionsForAVhost(vhost string) (string, error)
}

// NewDefinition creates a new RabbitMQ definition API client with the provided configuration.
// It returns a DefinitionInterface.
func NewDefinition(config map[string]interface{}) (DefinitionInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}

	return &definition{
		client,
	}, nil
}

// ListDefinitions The server definitions - exchanges, queues, bindings, users, virtual hosts, permissions, topic permissions, and parameters.
func (d *definition) ListDefinitions() (string, error) {
	return d.client.Get("/api/definitions")
}

// CreateDefinition sets RabbitMQ definitions using the provided definition string in JSON format.
func (d *definition) CreateDefinition(definition string) (string, error) {
	if err := validateParam(definition, "definition"); err != nil {
		return "", err
	}

	// Create a map for the JSON data
	data := map[string]string{"file": definition}

	// Marshal the map into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	path := "/api/definitions"
	return d.client.Post(path, string(jsonData))
}

// ListDefinitionsForAVhost The server definitions for a given virtual host - exchanges, queues, bindings and policies.
func (d *definition) ListDefinitionsForAVhost(vhost string) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}

	path := "/api/definitions/" + url.QueryEscape(vhost)
	return d.client.Get(path)
}
