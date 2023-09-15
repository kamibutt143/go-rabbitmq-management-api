package go_rabbitmq_management_api

import "github.com/kamibutt143/go-rabbitmq-management-api/lib"

type RabbitMQManager struct {
	Vhost go_rabbitmq_management_api.VhostInterface
}

func NewRabbitMQManager(config map[string]interface{}) (*RabbitMQManager, error) {

	vhost, err := go_rabbitmq_management_api.NewVhost(config)
	if err != nil {
		return nil, err
	}

	return &RabbitMQManager{
		Vhost: vhost,
	}, nil
}
