package gometr

import (
	"fmt"
	"math/rand"
	"time"
)

type GoMetrClient struct {
	URL     string
	Timeout time.Duration
}

type HealthCheck struct {
	ServiceID string
}

func NewGoMetrClient(url string, timeout time.Duration) *GoMetrClient {
	return &GoMetrClient{
		URL:     url,
		Timeout: timeout,
	}
}

func (g *GoMetrClient) GetMetrics() string {
	return "Metrics"
}

func (g *GoMetrClient) Ping() error {
	return nil
}

func (g *GoMetrClient) GetID() string {
	return g.URL
}

func (g *GoMetrClient) Health() bool {
	return rand.Intn(2) == 0
}

func (g *GoMetrClient) String() string {
	return fmt.Sprintf("URL: %s, Timeout: %s", g.URL, g.Timeout)
}
