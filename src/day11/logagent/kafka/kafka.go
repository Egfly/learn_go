package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	KafkaClient sarama.SyncProducer
)

func InitKafka(addr string) (err error) {
	// 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 连接客户端
	KafkaClient, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error(" init kafka producer failed, err:", err)
		return
	}
	logs.Debug("init kafka success")
	return
}

func Send(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}
	pid, offset, sendErr := KafkaClient.SendMessage(msg)
	if sendErr != nil {
		logs.Error("send message failed, err:", sendErr)
	}
	logs.Info("send success, pid:%v, offset:%v\n", pid, offset)

	return
}
