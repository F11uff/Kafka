FROM golang:1.21-alpine3.18

WORKDIR /app

RUN apk add --no-cache curl

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /email-service

CMD ["/email-service"]