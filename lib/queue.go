// Package lib provides an API client for interacting with RabbitMQ queue.
package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// queue represents the Queue API client.
type queue struct {
	client RabbitMQAPIClient
}

type QueueInterface interface {
	ListQueues() (string, error)
	GetAQueue(vhost string, queue string) (string, error)
	CreateQueue(vhost string, queue string, options map[string]string) (string, error)
	DeleteQueue(vhost string, queue string) (string, error)
	GetQueueBinding(vhost string, queue string) (string, error)
	PurgeQueue(vhost string, queue string) (string, error)
	SetQueueActions(vhost string, queue string, action string) (string, error)
	GetMessages(vhost string, queue string, options map[string]string) (string, error)
}

// NewQueue creates a new Queue API client with the provided configuration.
// It returns an QueueInterface.
func NewQueue(config map[string]interface{}) (QueueInterface, error) {
	client, err := NewRabbitMQAPIClient(config)
	if err != nil {
		return nil, err
	}

	return &queue{
		client,
	}, nil
}

// ListQueues retrieves a list of all queues.
func (b *queue) ListQueues() (string, error) {
	path := "/api/queues/"
	return b.client.Get(path)
}

// GetAQueue retrieves a queue for given vhost and queue name
func (b *queue) GetAQueue(vhost string, queue string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}
	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue)
	return b.client.Get(path)
}

// CreateQueue creates a new queue with the given parameters and options.
func (b *queue) CreateQueue(vhost string, queue string, options map[string]string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}
	// Marshal the map into a JSON string
	jsonData, err := json.Marshal(options)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue)
	return b.client.Put(path, string(jsonData))
}

// DeleteQueue deletes a specific queue.
func (b *queue) DeleteQueue(vhost string, queue string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}
	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue)
	return b.client.Delete(path)
}

// GetQueueBinding retrieves bindings for a queue.
func (b *queue) GetQueueBinding(vhost string, queue string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}
	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue) + "/bindings"
	return b.client.Get(path)
}

// PurgeQueue purge the messages of the queue
func (b *queue) PurgeQueue(vhost string, queue string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}
	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue) + "/contents"
	return b.client.Delete(path)
}

// SetQueueActions set the actions of the queue
func (b *queue) SetQueueActions(vhost string, queue string, action string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}
	if action == "" {
		return "", fmt.Errorf("missing action parameter")
	}
	// Create a map for the JSON data
	data := map[string]string{"action": action}

	// Marshal the map into a JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue) + "/actions"
	return b.client.Post(path, string(jsonData))
}

// GetMessages retrieves the message from the queue for the given parameters and options.
func (b *queue) GetMessages(vhost string, queue string, options map[string]string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}

	if options["count"] == "" {
		return "", errors.New("missing count parameter")
	}
	if options["requeue"] == "" {
		return "", errors.New("missing requeue parameter")
	}
	if options["encoding"] == "" {
		return "", errors.New("missing encoding parameter")
	}

	jsonData, err := json.Marshal(options)
	if err != nil {
		return "", err
	}

	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue) + "/get"
	return b.client.Post(path, string(jsonData))
}

// validateQueueParams checks if required parameters are missing.
func validateQueueParams(vhost, queue string) error {
	if vhost == "" {
		return errors.New("missing vhost parameter")
	}
	if queue == "" {
		return errors.New("missing queue parameter")
	}
	return nil
}
