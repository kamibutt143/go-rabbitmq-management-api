// Package lib provides an API client for interacting with RabbitMQ queue.
package lib

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// queue represents the Queue API client.
type queue struct {
	client RabbitMQAPIClient
}

type QueueInterface interface {
	ListQueues(pagination map[string]interface{}) (string, error)
	ListQueuesForAVhost(vhost string, pagination map[string]interface{}) (string, error)
	GetAQueueForAVhost(vhost string, queue string) (string, error)
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

// ListQueues A list of all queues. Use pagination parameters to filter queues.
func (b *queue) ListQueues(pagination map[string]interface{}) (string, error) {
	path := "/api/queues/"
	query, err := buildPaginationQuery(pagination)
	if err != nil {
		return "", err
	}
	if query != "" {
		path += query
	}

	return b.client.Get(path)
}

// ListQueuesForAVhost A list of all queues in a given virtual host. Use pagination parameters to filter queues.
func (b *queue) ListQueuesForAVhost(vhost string, pagination map[string]interface{}) (string, error) {
	if err := validateParam(vhost, "vhost"); err != nil {
		return "", err
	}
	path := "/api/queues/" + url.QueryEscape(vhost)
	query, err := buildPaginationQuery(pagination)
	if err != nil {
		return "", err
	}
	if query != "" {
		path += query
	}

	return b.client.Get(path)
}

// GetAQueueForAVhost retrieves a queue for given vhost and queue name
func (b *queue) GetAQueueForAVhost(vhost string, queue string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}

	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue)
	return b.client.Get(path)
}

// CreateQueue creates a new queue with the given parameters and options.
// To PUT a queue, you will need a body looking something like this:
// {"auto_delete":false,"durable":true,"arguments":{},"node":"rabbit@smacmullen"}
// All keys are optional.
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

// GetQueueBinding A list of all bindings on a given queue.
func (b *queue) GetQueueBinding(vhost string, queue string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}
	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue) + "/bindings"
	return b.client.Get(path)
}

// PurgeQueue Contents of a queue. DELETE to purge. Note you can't GET this.
func (b *queue) PurgeQueue(vhost string, queue string) (string, error) {
	if err := validateQueueParams(vhost, queue); err != nil {
		return "", err
	}
	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue) + "/contents"
	return b.client.Delete(path)
}

// SetQueueActions Actions that can be taken on a queue. POST a body like:
// {"action":"sync"}
func (b *queue) SetQueueActions(vhost string, queue string, action string) (string, error) {
	params := map[string]string{
		"vhost":  vhost,
		"queue":  queue,
		"action": action,
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	if action != "sync" && action != "cancel_sync" {
		return "", fmt.Errorf("provided value is ('%s') is not a valid value", action)
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

// GetMessages Get messages from a queue. (This is not an HTTP GET as it will alter the state of the queue.)
// You should post a body looking like:
// {"count":5,"ackmode":"ack_requeue_true","encoding":"auto","truncate":50000}
func (b *queue) GetMessages(vhost string, queue string, options map[string]string) (string, error) {
	params := map[string]string{
		"vhost":    vhost,
		"queue":    queue,
		"count":    options["count"],
		"ackmode":  options["ackmode"],
		"encoding": options["encoding"],
	}

	if err := validateParams(params); err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(options)
	if err != nil {
		return "", err
	}

	path := "/api/queues/" + url.QueryEscape(vhost) + "/" + url.QueryEscape(queue) + "/get"
	return b.client.Post(path, string(jsonData))
}
