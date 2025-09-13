```markdown
# Архитектура проекта

```mermaid
flowchart TB
    UserService["UserService\n(API Producer)"] -->|"Produce message"| Kafka[("Kafka Message Broker")]
    Kafka --> |"Email Events"| EmailService["EmailService\n(API Consumer)"]
    Kafka --> |"Phone Events"| SmsService["SmsService\n(API Consumer)"]
    Kafka --> |"All Events"| AnalyticsService["AnalyticsService\n(API Consumer)"]
    
    classDef producer fill:#401839,color:white,stroke:#ff6b6b,stroke-width:2px
    classDef consumer fill:#401839,color:white,stroke:#4ecdc4,stroke-width:2px
    classDef analytics fill:#401839,color:white,stroke:#f9ca24,stroke-width:2px
    classDef broker fill:#154a3f,color:white,stroke:#1dd1a1,stroke-width:2px

    class UserService producer
    class EmailService,SmsService consumer
    class AnalyticsService analytics
    class Kafka broker
