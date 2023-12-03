package main

import (
	"context"
	"flag"
	"log"

	monitorpb "github.com/quamejnr/seshat/gen/monitor/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:9091", "the address to connect to")
	url  = flag.String("url", "https://google.com", "the url to make request to")
	n    = flag.Int("n", 3, "the number of requests to be made")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := monitorpb.NewMonitorServiceClient(conn)

	ctx := context.Background()

	r, err := c.GetLatency(ctx, &monitorpb.GetLatencyRequest{Url: *url, NumOfRequests: int32(*n)})
	if err != nil {
		log.Fatalf("could not get latency: %v", err)
	}
	log.Printf("Response Time: %f", r.GetRequestLatencyMs())
}
