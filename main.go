package go_rabbitmq_management_api

import "github.com/kamibutt143/go-rabbitmq-management-api/lib"

type RabbitMQManager struct {
	vhost lib.VhostInterface
}

func NewRabbitMQManager(config map[string]interface{}) (*RabbitMQManager, error) {

	vhost, err := lib.NewVhost(config)
	if err != nil {
		return nil, err
	}

	return &RabbitMQManager{
		vhost,
	}, nil
}
