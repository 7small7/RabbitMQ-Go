package hello

import (
	"github.com/streadway/amqp"

	"rabbit_mq/logs"
)

func connection() *amqp.Channel {
	// `amqp://RabbitMQ服务密码:RabbitMQ服务密码@RabbitMQ主机地址:RabbitMQ服务端口/`
	dial, err := amqp.Dial("amqp://myuser:mypass@192.168.2.100:5672/")
	if err != nil {
		logs.FailOnError(err, "RabbitMQ TCP 连接失败")
	}

	channel, err := dial.Channel()
	if err != nil {
		logs.FailOnError(err, "RabbitMQ Channel 连接失败")
	}

	return channel
}
