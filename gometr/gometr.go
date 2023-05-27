package gometr

import (
	"time"
)

type GoMetrClient struct {
	URL     string
	TimeOut time.Duration
}

type HealthCheck struct {
	ServiceID string
	Healthy   bool
}

func NewGoMetrClient(url string, timeout time.Duration) *GoMetrClient {
	return &GoMetrClient{
		URL:     url,
		TimeOut: timeout,
	}
}

func (g *GoMetrClient) getHealth() HealthCheck {
	// Логика получения состояния здоровья сервиса
	// ...

	return HealthCheck{
		ServiceID: g.URL,
		Healthy:   true,
	}
}

func (g *GoMetrClient) Health() bool {
	timer := time.NewTimer(g.TimeOut)
	defer timer.Stop()

	select {
	case <-timer.C:
		return false
	default:
		return g.getHealth().Healthy
	}
}
