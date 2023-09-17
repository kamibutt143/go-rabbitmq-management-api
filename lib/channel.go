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
	ListChannels(pagination map[string]interface{}) (string, error)
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

// ListChannels A list of all open channels. Use pagination parameters to filter channels.
func (c *channel) ListChannels(pagination map[string]interface{}) (string, error) {
	path := "/api/channels"
	query, err := buildPaginationQuery(pagination)
	if err != nil {
		return "", err
	}
	if query != "" {
		path += query
	}
	return c.client.Get(path)
}

// GetAChannel Details about an individual channel.
func (c *channel) GetAChannel(channel string) (string, error) {
	if err := validateParam(channel, "channel"); err != nil {
		return "", err
	}
	path := "/api/channels/" + url.QueryEscape(channel)
	return c.client.Get(path)
}
