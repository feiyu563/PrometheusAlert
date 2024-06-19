package controllers

import (
	"github.com/IBM/sarama"
	"github.com/astaxie/beego"
	"time"
)

func GetKafkaProducer() sarama.SyncProducer {
	kafka_server := beego.AppConfig.Strings("kafka_server")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 等待所有follower都回复ack，确保Kafka不会丢消息
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner

	// 对Key进行Hash，同样的Key每次都落到一个分区，这样消息是有序的
	producer, err := sarama.NewSyncProducer(kafka_server, config)

	if err != nil {
		panic(err.Error())
	}

	return producer
}

func SendKafka(message, logsign string) string {
	//发送kafka
	open := beego.AppConfig.String("open-kafka")
	if open != "1" {
		beego.Info(logsign, "[kafka]", "kafka未配置未开启状态,请先配置open-kafka为1")
		return "kafka未配置未开启状态,请先配置open-kafka为1"
	}
	t1 := time.Now().UnixMilli()
	producer := GetKafkaProducer()
	kafka_topic := beego.AppConfig.String("kafka_topic")
	Key_string := beego.AppConfig.String("kafka_key") + "-" + logsign
	msg := &sarama.ProducerMessage{
		Topic: kafka_topic,
		Value: sarama.StringEncoder(message),
		Key:   sarama.StringEncoder(Key_string),
	}

	partition, offset, err := producer.SendMessage(msg)
	defer producer.Close()
	t2 := time.Now().UnixMilli()
	if err == nil {
		beego.Debug("发送kafka消息:", Key_string, "成功, partition:", partition, ",offset:", offset, ",cost:", t2-t1, " ms")
		return "发送kafka消息:" + Key_string + "成功"
	} else {
		beego.Error(err)
		return err.Error()
	}

}
