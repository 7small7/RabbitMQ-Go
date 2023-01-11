package work

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"

	"rabbit_mq/logs"
)

func Send() {
	channel := connection()

	// 1. 声明队列
	declare, err := channel.QueueDeclare("logs", false, false, false, false, nil)
	if err != nil {
		logs.FailOnError(err, "Hello队列声明失败")
	}

	// 2. 组装和发送数据
	for i := 0; i < 1000; i++ {
		marshal, err := json.Marshal(map[string]interface{}{
			"name":   "公众号：菜鸟成长学习笔记",
			"author": fmt.Sprintf("%s%d", "kert---", i),
			"id":     1,
		})
		if err != nil {
			logs.FailOnError(err, "消息JSON序列化失败")
		}
		err = channel.Publish(
			"",
			declare.Name,
			false,
			false,
			amqp.Publishing{
				ContentType:  "text/plain", // 消息类型类型
				Body:         marshal,      // 字节数组消息体
				DeliveryMode: amqp.Persistent,
			})
		if err != nil {
			logs.FailOnError(err, "消息发送失败")
		}
		fmt.Println("第", i, "条消息发送成功")
		//time.Sleep(time.Second * 3)
	}
}
