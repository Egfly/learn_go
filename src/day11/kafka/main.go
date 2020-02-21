package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

/**
go kafka操作示例
*/
func main() {
	// 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 构建消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "first"
	msg.Value = sarama.StringEncoder("this is a message send by a golang project")

	// 连接客户端
	client, err := sarama.NewSyncProducer([]string{"192.168.31.239:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()

	//发送消息
	pid, offset, sendErr := client.SendMessage(msg)
	if sendErr != nil {
		fmt.Println("send message failed, err:", sendErr)
	}

	fmt.Printf("pid: %v, offset: %v\n", pid, offset)
}
