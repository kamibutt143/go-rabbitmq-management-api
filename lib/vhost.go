// Package lib provides an API client for interacting with RabbitMQ VHost
package lib

import (
	"net/url"
)

// Vhost represents a RabbitMQ virtual host client.
type vhost struct {
	client RabbitMQAPIClient
}

// VhostInterface defines the interface for interacting with RabbitMQ virtual hosts.
type VhostInterface interface {
	ListVhosts() (string, error)
	GetAVhost(vhost string) (string, error)
	CreateAVhost(vhost string) (string, error)
	GetVhostPermissions(vhost string) (string, error)
	GetVhostTopicPermissions(vhost string) (string, error)
	GetVhostConnections(vhost string, pagination map[string]interface{}) (string, error)
	GetVhostChannels(vhost string, pagination map[string]interface{}) (string, error)
	DeleteVhost(vhost string) (string, error)
	StartVhostOnANode(vhost string, node string) (string, error)
}

// NewVhost creates a new Vhost instance with the provided configuration.
func NewVhost(config map[string]interface{}) (VhostInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}
	return &vhost{
		client,
	}, nil
}

// ListVhosts retrieves a list of all RabbitMQ virtual hosts.
func (v *vhost) ListVhosts() (string, error) {
	return v.client.Get("/api/vhosts")
}

// GetAVhost retrieves information about a specific RabbitMQ virtual host.
func (v *vhost) GetAVhost(vhost string) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost)
	return v.client.Get(path)
}

// CreateAVhost creates a new RabbitMQ virtual host.
func (v *vhost) CreateAVhost(vhost string) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost)
	return v.client.Put(path, "")
}

// GetVhostPermissions retrieves permissions for a specific RabbitMQ virtual host.
func (v *vhost) GetVhostPermissions(vhost string) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost) + "/permissions"
	return v.client.Get(path)
}

// GetVhostTopicPermissions A list of all topic permissions for a given virtual host.
func (v *vhost) GetVhostTopicPermissions(vhost string) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost) + "/topic-permissions"
	return v.client.Get(path)
}

// GetVhostConnections A list of all open connections in a specific virtual host. Use pagination parameters to filter connections.
func (v *vhost) GetVhostConnections(vhost string, pagination map[string]interface{}) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}
	path := "/api/vhost/" + url.QueryEscape(vhost) + "/connections"

	query, err := buildPaginationQuery(pagination)
	if err != nil {
		return "", err
	}

	if query != "" {
		path += query
	}

	return v.client.Get(path)
}

// GetVhostChannels A list of all open channels in a specific virtual host. Use pagination parameters to filter channels.
func (v *vhost) GetVhostChannels(vhost string, pagination map[string]interface{}) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}
	path := "/api/vhost/" + url.QueryEscape(vhost) + "/channels"

	query, err := buildPaginationQuery(pagination)
	if err != nil {
		return "", err
	}

	if query != "" {
		path += query
	}

	return v.client.Get(path)
}

// DeleteVhost deletes a RabbitMQ virtual host.
func (v *vhost) DeleteVhost(vhost string) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost)
	return v.client.Delete(path)
}

// StartVhostOnANode Starts virtual host name on a specific node.
func (v *vhost) StartVhostOnANode(vhost string, node string) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}
	if err := validateParam(vhost, "node"); err != nil {
		return "", err
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost) + "/start/" + url.QueryEscape(node)
	return v.client.Post(path, "")
}
