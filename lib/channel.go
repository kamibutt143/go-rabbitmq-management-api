// Package lib provides an API client for interacting with RabbitMQ channels.
package lib

import (
	"net/url"
)

// channel represents the Channel API client.
type channel struct {
	client RabbitMQAPIClient
}

// ChannelInterface defines the interface for interacting with RabbitMQ channels.
type ChannelInterface interface {
	ListChannels() (string, error)
	GetAChannel(channel string) (string, error)
}

// NewChannel creates a new Channel API client with the provided configuration.
// It returns a ChannelInterface.
func NewChannel(config map[string]interface{}) (ChannelInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}

	return &channel{
		client,
	}, nil
}

// ListChannels retrieves a list of all channels.
func (c *channel) ListChannels() (string, error) {
	return c.client.Get("/api/channels")
}

// GetAChannel retrieves information about a specific channel.
func (c *channel) GetAChannel(channel string) (string, error) {
	if err := validateParam(channel, "channel"); err != nil {
		return "", err
	}
	path := "/api/channels/" + url.QueryEscape(channel)
	return c.client.Get(path)
}
