package workercore

import (
	"time"
	
	"github.com/riftbit/golif"
	"github.com/streadway/amqp"
)

// Configuration is basic configureation for worker
type Configuration struct {
	AsyncWorker           bool
	AsyncPoolSize           int32
	ConnectionString      string
	ConnectionConfig      amqp.Config
	ExchangeName          string
	QueueName             string
	QueueDurable          bool
	QueueArguments        map[string]interface{}
	QueueAutoDelete       bool
	QueueExclusive        bool
	QueueNoWait           bool
	QueueAutoACK          bool
	PrefetchCount         int
	UseDelayedQueue       bool
	DelayedQueueArguments map[string]interface{}
	DefaultRetryCount     int32
	DefaultRetryDelay     time.Duration
}

// ProcessFunction 
// service field needs to pass application data, services, configs, etc.
type ProcessFunction func(service interface{}, amqpMSG *amqp.Delivery) (retryCnt int32, retryDelay time.Duration, err error)

// Worker basic client
type Worker struct {
	config                    *Configuration
	logger                    golif.Logger
	amqpConnection            *amqp.Connection
	amqpChannel               *amqp.Channel
	amqpQueue                 amqp.Queue
	amqpQueueDelayed          amqp.Queue
	amqpMessages              <-chan amqp.Delivery
	amqpNotifyCloseConnection chan *amqp.Error
	amqpQueueDelayedName      string
	shutdownCh                chan bool
	processFunction           ProcessFunction
	externalService           interface{}
}
