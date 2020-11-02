package rabbitmq

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://root:1234@127.0.0.1:5672/"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// Key
	Key string
	// 连接信息
	Mqurl string
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	var err error
	// 创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		log.Fatalf("%s:%s", "创建连接错误！", err)
	}

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		log.Fatalf("%s:%s", "获取channel失败！", err)
	}

	return rabbitmq
}

// Destroy 断开channel和connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// NewRabbitMQSimple
// 简单模式Step 1.创建简单模式下的RabbitMq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式Step 2:简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 声明队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		// 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

// ConsumeSimple 使用 goroutine 消费消息
func (r *RabbitMQ) ConsumeSimple() {
	// 声明队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 队列消费是否阻塞
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 协程处理
	go func() {
		for d := range msgs {
			// 实现我们要实现的逻辑函数
			//log.Printf("Received a message: %s", d.Body)
			fmt.Println("Consumer:", string(d.Body))
			if string(d.Body) == "sleep" {
				time.Sleep(5 * time.Second)
			}
		}
	}()
	log.Printf("...Waiting for message")
	<-forever
}
