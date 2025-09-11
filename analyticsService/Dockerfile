FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache curl

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /analytics-service

CMD ["/analytics-service"]