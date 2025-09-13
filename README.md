# Архитектура проекта

```mermaid
flowchart TB
    UserService("UserService (API Producer)") L_UserService_Kafka_0@-- Produce message --> Kafka[("Kafka")]
    Kafka L_Kafka_EmailService_0@-- Email Events --> EmailService("EmailService (API Consumer)")
    Kafka L_Kafka_SmsService_0@-- Phone Events --> SmsService("SmsService (API Consumer)")
    Kafka L_Kafka_AnalyticsService_0@-- All Events --> AnalyticsService("AnalyticsService (API Consumer)")
    style UserService fill:#401839,color:white, stroke:black, stroke-width:3
    style Kafka fill:#154a3f,color:white, stroke:black, stroke-width:3
    style EmailService fill:#401839,color:white, stroke:black, stroke-width:3
    style SmsService fill:#401839,color:white, stroke:black, stroke-width:3
    style AnalyticsService fill:#401839,color:white, stroke:black, stroke-width:3
    L_UserService_Kafka_0@{ animation: slow } 
    L_Kafka_EmailService_0@{ animation: slow } 
    L_Kafka_SmsService_0@{ animation: slow } 
    L_Kafka_AnalyticsService_0@{ animation: slow }
```

## Компоненты системы

| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Сервис &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; | &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Информация  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|
|:-------|:-----------|
| **UserService** | Генерирует события пользователей |
| **Kafka** | Брокер сообщений |
| **EmailService** | Обрабатывает email-события |
| **SmsService** | Обрабатывает SMS-события |
| **AnalyticsService** | Анализирует все события 
