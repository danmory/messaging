FROM golang:1.19

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/local/bin/app .

ENV QUEUE_HOST=rabbitmq
ENV QUEUE_PORT=5672
ENV QUEUE_USER=guest
ENV QUEUE_PASS=guest
ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASS=postgres
ENV DB_NAME=messages

CMD ["app"]
