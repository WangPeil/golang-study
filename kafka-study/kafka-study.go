package kafka_study

import (
	"context"
	"flag"
	"github.com/Shopify/sarama"
	_ "github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

var (
	brokers  = "localhost:9092"
	version  = "2.2.2"
	group    = ""
	topics   = ""
	assignor = ""
	oldest   = true
	verbose  = false
)

func Start() {
	flag.StringVar(&brokers, "brokers", "localhost:9092", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&version, "version", "2.2.2", "Kafka cluster version")
	flag.StringVar(&group, "group", "go", "Kafka consumer group definition")
	flag.StringVar(&topics, "topics", "test-topic", "Kafka topics to be consumed, as a comma separated list")
	flag.StringVar(&assignor, "assignor", "range", "Consumer group partition assignment strategy")
	flag.BoolVar(&oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVar(&verbose, "verbose", false, "Sarama logging")

	if len(brokers) == 0 {
		panic("xxx")
	}
	if len(topics) == 0 {
		panic("xxx")
	}
	if len(group) == 0 {
		panic("xxx")
	}

	log.Println("Starting a new Sarama consumer")
	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama]", log.LstdFlags)
		kafkaVersion, err := sarama.ParseKafkaVersion(version)
		if err != nil {
			log.Panicf("xxx")
		}
		/**
		 * Construct a new Sarama configuration
		 */
		config := sarama.NewConfig()
		config.Version = kafkaVersion

		switch assignor {
		case "sticky":
			config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
		case "range":
			config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
		case "roundrobin":
			config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
		}
		if oldest {
			config.Consumer.Offsets.Initial = sarama.OffsetOldest
		}
		consumer := Consumer{ready: make(chan bool)}
		ctx, cancel := context.WithCancel(context.Background())
		client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
		if err != nil {
			log.Panicf("xxx")
		}
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if err := client.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
					log.Panicf("xxx")
				}
				if ctx.Err() != nil {
					return
				}
				consumer.ready = make(chan bool)
			}
		}()
		<-consumer.ready
		log.Println("Sarama consumer up and running")

		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			log.Println("...")
		case <-sigterm:
			log.Println("xxx")

		}
		cancel()
		wg.Wait()
		if err = client.Close(); err != nil {
			log.Panicf("xxx")
		}

	}

}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Println("message claimed", message.Value)
		session.MarkMessage(message, "")
	}
	return nil
}
