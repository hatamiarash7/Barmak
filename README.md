# Barmak

This is a simple project to test 4 popular [Kafka](https://kafka.apache.org/) libraries for [Go](https://go.dev/). Our libraries are :

-   [Confluent](https://github.com/confluentinc/confluent-kafka-go)
-   [Sarama](https://github.com/Shopify/sarama)
-   [Goka](https://github.com/lovoo/goka)
-   [Kafka-Go](https://github.com/segmentio/kafka-go)

## How to

We can get a simple benchmark from these libraries.

```go
func BenchmarkProduceConfluent(b *testing.B) {
    count := b.N

    messages := make([]Message, count)

    for i := 0; i < count; i++ {
        messages[i] = Message{Key: "Confluent", Value: fmt.Sprint(i)}
    }

    Confluent(messages)
}
```

## Run

There is a `Makefile` for you to run all commands.

```no-lang
Usage: make [target]

  help        Show this help message
  build       Build app's binary
  run         Run the app
  test        Run unit tests
  benchmark   Run benchmark tests
```

Execute the following command to run the benchmark:

```bash
$ make benchmark

goos: darwin
goarch: arm64
pkg: arash-hatami.ir/Barmak/kafka
BenchmarkProduceSegmentio-8          367           3235080 ns/op
BenchmarkProduceConfluent-8         1137           1006395 ns/op
BenchmarkProduceSarama-8             496           2136077 ns/op
BenchmarkProduceGoka-8                10         104631917 ns/op
PASS
ok      arash-hatami.ir/Barmak/kafka    6.299s
'graphic/operations.png' and 'graphic/time_operations.png' graphics were generated.
```

And you can see the results in `graphic` directory.

![operations](.github/operations.png)
![operations](.github/time_operations.png)

Default process will be handled by the Go [testing](https://pkg.go.dev/testing#hdr-Benchmarks) package. We can also use `-benchtime=100x` argument to set the minimum amount of time that the benchmark function will run.
