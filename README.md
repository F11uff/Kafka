# Архитектура проекта

Взаимодействие микросервисов через Kafka:

```mermaid
flowchart TB
    UserService("UserService (API Producer)") -->|Produce message| Kafka[(Kafka)]
    Kafka --> |Email Events| EmailService("EmailService (API Consumer)")
    Kafka --> |Phone Events| SmsService("SmsService (API Consumer)")
    Kafka --> |All Events| AnalyticsService("AnalyticsService (API Consumer)")
    style UserService fill:#401839,color:white, stroke:black, stroke-width:3
    style AnalyticsService fill:#401839,color:white, stroke:black, stroke-width:3
    style EmailService fill:#401839,color:white, stroke:black, stroke-width:3
    style SmsService fill:#401839,color:white, stroke:black, stroke-width:3
    style Kafka fill:#154a3f,color:white, stroke:black, stroke-width:3
