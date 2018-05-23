# Go - Kafka Exercise

This repository contains a [POC](https://en.wikipedia.org/wiki/Proof_of_concept) using [Golang](https://golang.org/) and [Kafka](https://kafka.apache.org/). The excercise consists in two parts producer and consumer, the producer reads the lines of a file and publishes them in a Kafka topic. The consumer receives the messages and writes them to a different file.

We are using [the Confluent's Golang Client for Apache KafkaTM](https://github.com/confluentinc/confluent-kafka-go#confluents-golang-client-for-apache-kafkatm)

## Usage

Just clone the repository into your local machine, to run the producer, you can type the command below:

```
go run cmd/producer/main.go source.txt
```

And in order to run the consumer, you can type:

```
go run cmd/consumer/main.go
```

## License

This project is licensed under the **MIT License**.

## Contributions
Feel free to make any comment, pull request, code review, shared post, fork or feedback. Everything is welcome.

## Authors

**Sebastian Zapata Mardini** - [GitHub profile](https://github.com/Mardiniii)
