package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

type UserEvent struct {
	EventType string    `json:"event_type"`
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Timestamp time.Time `json:"timestamp"`
}

type RegisterRequest struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func getKafkaWriter(topic string) *kafka.Writer {
	kafkaBrokers := getEnv("KAFKA_BROKERS", "localhost:9092")
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaBrokers),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	kafkaBrokers := getEnv("KAFKA_BROKERS", "kafka:9092")
	userEventsTopic := getEnv("KAFKA_TOPIC_USER_EVENTS", "user-events")
	analyticsTopic := getEnv("KAFKA_TOPIC_ANALYTICS_EVENTS", "analytics-events")
	env := getEnv("ENVIRONMENT", "development")
	logLevel := getEnv("LOG_LEVEL", "info")

	fmt.Printf("Starting User Service in %s environment\n", env)
	fmt.Printf("Kafka: %s, Topics: %s, %s\n", kafkaBrokers, userEventsTopic, analyticsTopic)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		userEvent := UserEvent{
			EventType: "user_registered",
			UserID:    fmt.Sprintf("user_%d", time.Now().UnixNano()),
			Email:     req.Email,
			Phone:     req.Phone,
			Timestamp: time.Now(),
		}

		// Send to user events
		userWriter := getKafkaWriter(userEventsTopic)
		userEventJSON, _ := json.Marshal(userEvent)
		err := userWriter.WriteMessages(context.Background(), kafka.Message{
			Value: userEventJSON,
		})
		userWriter.Close()

		if err != nil {
			log.Printf("Failed to send user event: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		analyticsWriter := getKafkaWriter(analyticsTopic)
		analyticsEvent := map[string]interface{}{
			"type":      "user_registration",
			"user_id":   userEvent.UserID,
			"timestamp": time.Now(),
		}
		analyticsJSON, _ := json.Marshal(analyticsEvent)
		analyticsWriter.WriteMessages(context.Background(), kafka.Message{
			Value: analyticsJSON,
		})
		analyticsWriter.Close()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"user_id": userEvent.UserID,
		})
	})

	port := getEnv("PORT", "8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
