# Go RabbitMQ Management API Client

The Go RabbitMQ Management API Client is a powerful library that allows you to interact with RabbitMQ, a leading message broker, via its HTTP API. It simplifies the process of managing various RabbitMQ components such as exchanges, queues, connections, nodes, and more, all from within your Go applications.

## Features

- **Easy Setup**: Get started quickly by creating a client instance with your RabbitMQ server details.
- **Comprehensive**: This library supports a wide range of RabbitMQ management functions, from viewing an overview of your RabbitMQ setup to managing individual components like queues and exchanges.
- **Flexible**: Tailor your requests with filters and parameters to precisely control which data you retrieve or actions you perform.
- **Error Handling**: Handle errors gracefully with callback functions, ensuring your application stays robust.
- **Support**: This client library is designed to work with RabbitMQ 3.x versions and requires the RabbitMQ Management UI plugin to be installed and enabled.

## Categories
- **Binding** provides an API client for interacting with RabbitMQ binding.
- **Channel** provides an API client for interacting with RabbitMQ channel.
- **Cluster** provides an API client for interacting with RabbitMQ cluster.
- **Connection** provides an API client for interacting with RabbitMQ connection.
- **Definition** provides an API client for interacting with RabbitMQ definition.
- **Exchange** provides an API client for interacting with RabbitMQ exchange.
- **Node** provides an API client for interacting with RabbitMQ node.
- **Queue** provides an API client for interacting with RabbitMQ queue.
- **Vhost** provides an API client for interacting with RabbitMQ vhost.

## Installation

You can easily integrate this library into your Go project using Go modules:

```bash
go get github.com/kamibutt143/go-rabbitmq-management-api
```

## Supported RabbitMQ Versions
Our RabbitMQ HTTP API Client is compatible with RabbitMQ 3.x versions. Please note that to use this library, you must have the [RabbitMQ Management UI plugin](http://www.rabbitmq.com/management.html) installed and enabled for your RabbitMQ instance.

## Usage

```go
package main

import (
	"fmt"
	m "github.com/kamibutt143/go-rabbitmq-management-api"
)

func main() {

	//Usage example:
	config := map[string]interface{}{
		"host":     "https://localhost",
		"port":     15671,
		"user":     "guest",
		"password": "guest",
		"timeout":  30,
	}

	manager, err := m.NewRabbitMQManager(config)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Example 1: List virtual hosts
	vhosts, err := manager.Vhost.ListVhosts()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(vhosts)
}
```
## License
This Go RabbitMQ Management API Client is open-source software licensed under the MIT License.