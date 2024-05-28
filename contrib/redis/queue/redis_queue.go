package queue

import (
	"github.com/redis/go-redis/v9"
	"github.com/sql668/go-toolkit/storage"
	gtsm "github.com/sql668/go-toolkit/storage/message"
)

// NewRedis redis模式
func NewRedis(
	producerOptions *ProducerOptions,
	consumerOptions *ConsumerOptions,
) (*Redis, error) {
	var err error
	r := &Redis{}
	r.producer, err = r.newProducer(producerOptions)
	if err != nil {
		return nil, err
	}
	r.consumer, err = r.newConsumer(consumerOptions)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Redis cache implement
type Redis struct {
	client   *redis.Client
	consumer *Consumer
	producer *Producer
}

func (Redis) String() string {
	return "redis"
}

func (r *Redis) newConsumer(options *ConsumerOptions) (*Consumer, error) {
	if options == nil {
		options = &ConsumerOptions{}
	}
	return NewConsumerWithOptions(options)
}

func (r *Redis) newProducer(options *ProducerOptions) (*Producer, error) {
	if options == nil {
		options = &ProducerOptions{}
	}
	return NewProducerWithOptions(options)
}

func (r *Redis) Append(message storage.Messager) error {
	err := r.producer.Enqueue(&gtsm.Message{
		ID:     message.GetID(),
		Stream: message.GetStream(),
		Values: message.GetValues(),
	})
	return err
}

func (r *Redis) Register(name string, f storage.ConsumerFunc) {
	r.consumer.Register(name, func(message *gtsm.Message) error {
		m := new(gtsm.Message)
		m.SetValues(message.Values)
		m.SetStream(message.Stream)
		m.SetID(message.ID)
		return f(m)
	})
}

func (r *Redis) Run() {
	r.consumer.Run()
}

func (r *Redis) Shutdown() {
	r.consumer.Shutdown()
}
