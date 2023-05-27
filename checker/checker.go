package checker

import (
	"context"
	"fmt"
	"time"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

type HealthCheck struct {
	ServiceID string
	Healthy   bool
}

type Checker struct {
	items []Checkable
	ctx   context.Context
	stop  context.CancelFunc
}

func NewChecker() *Checker {
	ctx, stop := context.WithCancel(context.Background())

	return &Checker{
		ctx:  ctx,
		stop: stop,
	}
}

func (c *Checker) Add(item Checkable) {
	c.items = append(c.items, item)
}

func (c *Checker) Run() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.Check()
		case <-c.ctx.Done():
			fmt.Println("Проверки остановлены")
			return
		}
	}
}

func (c *Checker) Stop() {
	c.stop()
}

func (c *Checker) Check() {
	for _, item := range c.items {
		if !item.Health() {
			fmt.Println(item.GetID() + " не работает")
		}
	}
}

func (c *Checker) String() string {
	var output string

	for _, item := range c.items {
		output += item.GetID() + "\n"
	}

	return output
}
