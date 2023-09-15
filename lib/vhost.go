package lib

import (
	"fmt"
	"net/url"
)

// Vhost represents a RabbitMQ virtual host client.
type Vhost struct {
	config map[string]interface{}
	client RabbitMQAPIClient
}

// VhostInterface defines the interface for interacting with RabbitMQ virtual hosts.
type VhostInterface interface {
	ListVhosts() (string, error)
	GetAVhost(vhost string) (string, error)
	CreateAVhost(vhost string) (string, error)
	GetVhostPermissions(vhost string) (string, error)
	DeleteVhost(vhost string) (string, error)
}

// NewVhost creates a new Vhost instance with the provided configuration.
func NewVhost(config map[string]interface{}) (VhostInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}
	return &Vhost{
		config,
		client,
	}, nil
}

// ListVhosts retrieves a list of all RabbitMQ virtual hosts.
func (v *Vhost) ListVhosts() (string, error) {
	return v.client.Get("/api/vhosts")
}

// GetAVhost retrieves information about a specific RabbitMQ virtual host.
func (v *Vhost) GetAVhost(vhost string) (string, error) {
	if vhost == "" {
		return "", fmt.Errorf("missing vhost parameter")
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost)
	return v.client.Get(path)
}

// CreateAVhost creates a new RabbitMQ virtual host.
func (v *Vhost) CreateAVhost(vhost string) (string, error) {
	if vhost == "" {
		return "", fmt.Errorf("missing vhost parameter")
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost)
	return v.client.Put(path, "")
}

// GetVhostPermissions retrieves permissions for a specific RabbitMQ virtual host.
func (v *Vhost) GetVhostPermissions(vhost string) (string, error) {
	if vhost == "" {
		return "", fmt.Errorf("missing vhost parameter")
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost) + "/permissions"
	return v.client.Get(path)
}

// DeleteVhost deletes a RabbitMQ virtual host.
func (v *Vhost) DeleteVhost(vhost string) (string, error) {
	if vhost == "" {
		return "", fmt.Errorf("missing vhost parameter")
	}

	path := "/api/vhosts/" + url.QueryEscape(vhost)
	return v.client.Delete(path)
}
