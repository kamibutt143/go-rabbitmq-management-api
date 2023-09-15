// Package lib provides an API client for interacting with RabbitMQ services.
package go_rabbitmq_management_api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// rabbitMQAPIClient represents the RabbitMQ API client.
type rabbitMQAPIClient struct {
	config map[string]interface{}
	client *http.Client
}

// RabbitMQAPIClient defines the interface for interacting with RabbitMQ services.
type RabbitMQAPIClient interface {
	Get(path string) (string, error)
	Put(path string, body string) (string, error)
	Patch(path string, body string) (string, error)
	Post(path string, body string) (string, error)
	Delete(path string) (string, error)
}

// NewRabbitMQAPIClient New creates a new RabbitMQ API client with the provided configuration.
// It returns a RabbitMQAPIClient interface.
func NewRabbitMQAPIClient(config map[string]interface{}) (RabbitMQAPIClient, error) {
	requiredKeys := []string{"host", "port", "user", "password"}

	for _, key := range requiredKeys {
		if _, exists := config[key]; !exists {
			err := fmt.Errorf("config key '%s' is missing", key)
			fmt.Println("Error:", err)
			return nil, err
		}
	}

	value, exists := config["timeout"].(int)
	if !exists || value <= 0 {
		value = 30000
	}
	timeout := time.Duration(value) * time.Millisecond

	return &rabbitMQAPIClient{
		config: config,
		client: &http.Client{Timeout: timeout},
	}, nil
}

// getUrl constructs a full URL from the given path using the client's configuration.
func (c *rabbitMQAPIClient) getUrl(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return c.config["host"].(string) + ":" + strconv.Itoa(c.config["port"].(int)) + "/" + path
}

// request performs an HTTP request with the specified method, path, and optional body.
func (c *rabbitMQAPIClient) request(path string, method string, body string) (string, error) {
	// Define the data you want to send in the POST request body as a byte slice
	var postData *bytes.Buffer
	if body != "" {
		postData = bytes.NewBuffer([]byte(body))
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, c.getUrl(path), postData)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}
	// Set the content type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// Set basic authentication
	req.SetBasicAuth(c.config["username"].(string), c.config["password"].(string))

	// Send the request
	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}

	// Make sure to close the response body if it's not nil
	err = resp.Body.Close()
	if err != nil {
		fmt.Println("Error coming in closing body of response:", err)
		return "", err
	}

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP Request failed with status: '%s'", resp.Status)
		fmt.Println("Error:", err)
		return "", err
	}

	// Print the response body (JSON)
	fmt.Println(string(responseBody))

	return string(responseBody), nil
}

// Get sends an HTTP GET request to the specified path.
func (c *rabbitMQAPIClient) Get(path string) (string, error) {
	return c.request(path, "GET", "")
}

// Post sends an HTTP POST request to the specified path with the given body.
func (c *rabbitMQAPIClient) Post(path string, body string) (string, error) {
	return c.request(path, "POST", body)
}

// Put sends an HTTP PUT request to the specified path with the given body.
func (c *rabbitMQAPIClient) Put(path string, body string) (string, error) {
	return c.request(path, "PUT", body)
}

// Patch sends an HTTP PATCH request to the specified path with the given body.
func (c *rabbitMQAPIClient) Patch(path string, body string) (string, error) {
	return c.request(path, "PATCH", body)
}

// Delete sends an HTTP DELETE request to the specified path.
func (c *rabbitMQAPIClient) Delete(path string) (string, error) {
	return c.request(path, "DELETE", "")
}
