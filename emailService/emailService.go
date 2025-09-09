package main

import "os"

func main() {
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_TOPIC")
}
