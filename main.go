package go_rabbitmq_management_api

import "github.com/kamibutt143/go-rabbitmq-management-api/lib"

type RabbitMQManager struct {
	Vhost      lib.VhostInterface
	Binding    lib.BindingInterface
	Exchange   lib.ExchangeInterface
	Channel    lib.ChannelInterface
	Cluster    lib.ClusterInterface
	Connection lib.ConnectionInterface
	Definition lib.DefinitionInterface
	Node       lib.NodeInterface
	Queue      lib.QueueInterface
}

func NewRabbitMQManager(config map[string]interface{}) (*RabbitMQManager, error) {
	vhost, err := lib.NewVhost(config)
	if err != nil {
		return nil, err
	}
	binding, err := lib.NewBinding(config)
	if err != nil {
		return nil, err
	}
	exchange, err := lib.NewExchange(config)
	if err != nil {
		return nil, err
	}
	channel, err := lib.NewChannel(config)
	if err != nil {
		return nil, err
	}
	cluster, err := lib.NewCluster(config)
	if err != nil {
		return nil, err
	}
	connection, err := lib.NewConnection(config)
	if err != nil {
		return nil, err
	}
	definition, err := lib.NewDefinition(config)
	if err != nil {
		return nil, err
	}
	node, err := lib.NewNode(config)
	if err != nil {
		return nil, err
	}
	queue, err := lib.NewQueue(config)
	if err != nil {
		return nil, err
	}
	return &RabbitMQManager{
		Vhost:      vhost,
		Binding:    binding,
		Exchange:   exchange,
		Channel:    channel,
		Cluster:    cluster,
		Connection: connection,
		Definition: definition,
		Node:       node,
		Queue:      queue,
	}, nil
}
