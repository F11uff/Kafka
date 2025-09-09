#!/bin/bash

sleep 10

echo "Waiting for Kafka to ready"

kafka-topics --bootstrap-server ${KAFKA_BROKERS} --create --topic ${KAFKA_TOPIC_USER_SERVICE_EVENT} --paritions 6 --replication-factor 1