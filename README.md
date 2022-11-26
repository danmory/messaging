# Messaging application

## Description

This is a messsaging application that consists of:

1. Producer
2. Consumer
3. Postgres database
4. RabbitMQ

Producer sends messages to RabbitMQ and Consumer consumes messages
from RabbitMQ and stores them in Postgres database.

## How to run

1. Clone the repository
2. Run the command

    `` $ POSTGRES_PASSWORD=postgres POSTGRES_USER=postgres docker compose up --build ``
