package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
	monitorpb "github.com/quamejnr/seshat/gen/monitor/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = flag.Int("port", 9091, "Port number of server")

type monitorService struct {
	monitorpb.UnimplementedMonitorServiceServer
}

func (s *monitorService) GetLatency(ctx context.Context, req *monitorpb.GetLatencyRequest) (*monitorpb.GetLatencyResponse, error) {
	url := req.Url
	numOfRequests := req.NumOfRequests
	var i int32
	var count, totalElapsedTime float32

	client := http.Client{Timeout: 5 * time.Second}

	for i = 0; i < numOfRequests; i++ {
		start := time.Now()
		log.Infof("Making request to %s", url)
		resp, err := client.Get(url)
		if os.IsTimeout(err) {
			log.Errorf("Request to %s timed out. %s", url, err.Error())
			continue
		}

		if err != nil {
			log.Fatalf("Error reaching url. Check if url is correct: %s", err.Error())
		}

		if resp.StatusCode == 200 {
			count++
			elapsedTime := time.Since(start)
			elapsedTimeMs := float32(elapsedTime / time.Millisecond)
			log.Infof("Response time %vms", elapsedTime)
			totalElapsedTime += float32(elapsedTimeMs)
		}
	}

	avgElapsedTimeMs := totalElapsedTime / count
	if math.IsNaN(float64(avgElapsedTimeMs)) {
		avgElapsedTimeMs = 0
	}
	return &monitorpb.GetLatencyResponse{Latency: avgElapsedTimeMs}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal("Failed to serve: %s", err.Error())
	}

	srv := grpc.NewServer()

	// Register services
	monitorpb.RegisterMonitorServiceServer(srv, &monitorService{})

	reflection.Register(srv)

	log.Infof("Server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}
}
