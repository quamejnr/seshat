# Seshat

A gRPC server that keeps metrics of your services.

## Motivation

I built Seshat in an attempt to deapen my understanding of gRPC although it
currently has the potential to be a library independent of gRPC. The idea of
Seshat is to help you easily create metrics for your services. You can easily
configure your gRPC server to send ping all your services and keep records of
the latency for later analysis.

## Usage

1. Run the command `git clone https://github.com/quamejnr/seshat`
2. From the root directory, run the command `go run metrics-server/main.go` to
   start the server
3. Open another terminal and run the command `go run metrics-client/main.go` to
   start the client. You can provide your own url as a flag, it defaults to
   `https://google.com`. Run `go run metrics-client/main.go` to see the flags
   available.
4. Make a couple of requests and then run
   `curl localhost:8000/metrics | json_pp` to see some metrics.

### Create metrics

There are currently two categories of metrics: - `CounterMetric` which can be
used to keep count or fixed value. - `MapMetric` which can be used to keep track
of a key, value pair.

#### MapMetric

> The below example has already been implemented in this code.

1. You can create a `MapMetric` by creating an instance of `MapMetric` and
   providing the name of your metric. For example:

```go
var l = NewMapMetric("LatencyMetrics")
```

2. You can then write your own function and use this instance created.
   Currently, `MapMetric` has only one method
   `SetKeyValue(key string, value any)`. For example;

```go
var l = NewMapMetric("LatencyMetrics")

func recordLatency(url string, responseTime float32) {
    l.SetKeyValue(url, responseTime)
}
```

3. You can use your function to keep metrics of your requests latency. You
   should see your created metric in your metrics when you run
   `curl localhost:8000/metrics | json_pp`.

#### CounterMetric

> The below example has already been implemented in this code.

1. You can create a `CounterMetric` using the function `NewCounterMetric`.

```go
var c = NewCounterMetric{"requestCounter"}
```

2. You can write your own function leveraging the methods from the instance
   created. Currently, `CounterMetric` has 3 methods; `SetKey`, `Increment`,
   `Decrement`.

```go
var c = NewCounterMetric{"requestCounter"}

func recordRequests() {
    c.Increment()
}
```

3. You can use your function to keep metrics of your requests latency. You
   should see your created metric in your metrics when you run
   `curl localhost:8000/metrics | json_pp`.

## Future Plans

- [ ] Add tests
- [ ] Make it a standalone library.

## Contributing

Contributions to Seshat are welcome! If you find a bug, have an idea for an
improvement, or want to add a new feature, please open an issue or create a pull
request on the Heimdall GitHub repository.

## License

Seshat is open-source software licensed under the MIT License. You are free to
use, modify, and distribute this tool according to the terms of the license.
