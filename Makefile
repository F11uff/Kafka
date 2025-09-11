all:

build:
	docker-compose build kafka kafka-setup zookeeper user-service email-service sms-service

up:
	docker-compose up kafka kafka-setup zookeeper user-service email-service sms-service

down:
	docker-compose down

logs:
	docker-compose logs -f sms-service user-service email-service

request:
	curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"email": "test@example.com", "phone": "+1234567890"}'