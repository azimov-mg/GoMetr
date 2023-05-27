package gometr

import (
	"fmt"
	"time"
)

type GoMetrClient struct {
	URL     string
	TimeOut time.Duration
}

func NewGoMetrClient(url string, timeout time.Duration) *GoMetrClient {
	return &GoMetrClient{
		URL:     url,
		TimeOut: timeout,
	}
}

func (g *GoMetrClient) GetMetrics() string {
	return "Metrics for " + g.URL
}

func (g *GoMetrClient) Ping() error {
	// Simulating a ping to the URL
	// ...

	return nil
}

func (g *GoMetrClient) GetID() string {
	return g.URL
}

func (g *GoMetrClient) getHealth() bool {
	// Simulating a health check
	// ...

	return true
}

func (g *GoMetrClient) Health() bool {
	timer := time.NewTimer(g.TimeOut)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("Timeout occurred for", g.URL)
		return false
	default:
		return g.getHealth()
	}
}
