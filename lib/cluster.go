// Package lib provides an API client for interacting with RabbitMQ clusters.
package lib

import (
	"encoding/json"
	"fmt"
)

// cluster represents the Cluster API client.
type cluster struct {
	client RabbitMQAPIClient
}

// ClusterInterface defines the interface for interacting with RabbitMQ clusters.
type ClusterInterface interface {
	GetClusterName() (string, error)
	SetClusterName(clusterName string) (string, error)
}

// NewCluster creates a new Cluster API client with the provided configuration.
// It returns a ClusterInterface.
func NewCluster(config map[string]interface{}) (ClusterInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}

	return &cluster{
		client,
	}, nil
}

// GetClusterName retrieves the name of the RabbitMQ cluster.
func (c *cluster) GetClusterName() (string, error) {
	return c.client.Get("/api/cluster-name")
}

// SetClusterName sets the name of the RabbitMQ cluster.
func (c *cluster) SetClusterName(clusterName string) (string, error) {
	if clusterName == "" {
		return "", fmt.Errorf("missing clusterName parameter")
	}

	// Create a map for the JSON data
	data := map[string]string{"name": clusterName}

	// Marshal the map into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	path := "/api/cluster-name/"
	return c.client.Put(path, string(jsonData))
}
