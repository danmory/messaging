FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/local/bin/app .

ENV QUEUE_HOST=rabbitmq
ENV QUEUE_PORT=5672
ENV QUEUE_USER=guest
ENV QUEUE_PASS=guest

CMD ["app"]
