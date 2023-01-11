package work

import (
	"encoding/json"
	"fmt"
	"time"

	"rabbit_mq/logs"
)

func Consumer() {
	channel := connection()

	// 声明一个队列
	declare, err := channel.QueueDeclare("logs", false, false, false, false, nil)
	if err != nil {
		logs.FailOnError(err, "队列连接失败")
	}

	// 从队列中消费消息
	consume, err := channel.Consume(
		declare.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logs.FailOnError(err, "消息获取失败")
	}

	// 从channel中获取数据(Consume方法返回的是一个channel)
	for val := range consume {
		msg := make(map[string]interface{})
		err = json.Unmarshal(val.Body, &msg)
		fmt.Println(msg)
		if err != nil {
			logs.FailOnError(err, "消息反序列化失败")
		}
		val.Ack(false)
		time.Sleep(time.Second * 2)
	}
}
