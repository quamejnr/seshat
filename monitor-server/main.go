package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"

	"context"
	"math"
	"os"
	"time"

	"github.com/charmbracelet/log"
	monitorpb "github.com/quamejnr/seshat/gen/monitor/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"encoding/json"
)

var l = NewMapMetric("LatencyMetrics")
var c = NewCounterMetric("requestCounter")

func recordRequests() {
	c.Increment()
}

func recordLatency(url string, responseTime float32) {
	l.SetLabelValue(url, responseTime)
}

type monitorService struct {
	monitorpb.UnimplementedMonitorServiceServer
}

func (s *monitorService) GetLatency(ctx context.Context, req *monitorpb.GetLatencyRequest) (*monitorpb.GetLatencyResponse, error) {
	url := req.Url
	numOfRequests := req.NumOfRequests
	var i int32
	var count, totalElapsedTime, avgElapsedTimeMs float32

	client := http.Client{Timeout: 5 * time.Second}

	for i = 0; i < numOfRequests; i++ {
		start := time.Now()
		log.Infof("Making request to %s", url)
		resp, err := client.Get(url)
		recordRequests()
		if os.IsTimeout(err) {
			log.Errorf("Request to %s timed out. %s", url, err.Error())
			continue
		}

		if err != nil {
			log.Errorf("Error reaching url. Check if url is correct: %s", err.Error())
			break
		}

		if resp.StatusCode == 200 {
			count++
			elapsedTime := time.Since(start)
			elapsedTimeMs := float32(elapsedTime / time.Millisecond)
			log.Infof("Response time %vms", elapsedTimeMs)
			totalElapsedTime += float32(elapsedTimeMs)
		}
	}

	avgElapsedTimeMs = totalElapsedTime / count

	if math.IsNaN(float64(avgElapsedTimeMs)) {
		avgElapsedTimeMs = 0
	}

	recordLatency(url, avgElapsedTimeMs)

	return &monitorpb.GetLatencyResponse{RequestLatencyMs: avgElapsedTimeMs}, nil
}

var (
	port = flag.Int("port", 9091, "Port number of server")
)

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

	// start gRPC server
	log.Infof("Server listening at %v", lis.Addr())
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %s", err.Error())
		}
	}()

	// Handle metrics
	http.HandleFunc("/metrics", metricHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var metrics []MetricInterface

type envelope map[string]interface{}

type MetricInterface interface {
	Envelope() envelope
}

type CounterMetric struct {
	namespace string
	value     int64
}

func NewCounterMetric(namespace string) *CounterMetric {
	res := &CounterMetric{
		namespace: namespace,
	}
	metrics = append(metrics, res)
	return res
}

func (c *CounterMetric) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *CounterMetric) Decrement() {
	atomic.AddInt64(&c.value, -1)
}

func (c *CounterMetric) SetValue(value int64) {
	atomic.StoreInt64(&c.value, value)
}

func (c *CounterMetric) Envelope() envelope {
	return envelope{c.namespace: atomic.LoadInt64(&c.value)}

}

type MapMetric struct {
	namespace string
	value     map[string]any
}

func NewMapMetric(namespace string) *MapMetric {
	res := &MapMetric{
		namespace: namespace,
		value:     map[string]any{},
	}
	metrics = append(metrics, res)
	return res
}

func (m *MapMetric) SetLabelValue(label string, value any) {
	m.value[label] = value
}

func (m *MapMetric) Envelope() envelope {
	type data map[string]interface{}
	dataList := []data{}
	for k, v := range m.value {
		d := data{
			"label": k,
			"value": v,
		}
		dataList = append(dataList, d)
	}
	return envelope{m.namespace: dataList}
}

func metricHandler(w http.ResponseWriter, r *http.Request) {
	data := []envelope{}

	for _, v := range metrics {
		env := v.Envelope()
		data = append(data, env)
	}

	type respData map[string]interface{}

	res := respData{
		"status":  http.StatusOK,
		"metrics": data,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving data. Contact support."))
		log.Fatal("Error marshalling response")
	}
}
