FROM golang:1.21-alpine3.18

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /user-service

EXPOSE 8080

CMD ["/user-service"]