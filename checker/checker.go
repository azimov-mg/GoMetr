package checker

import (
	"fmt"
	"strings"
	"sync"
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
	mutex sync.Mutex
}

func NewChecker() *Checker {
	return &Checker{
		items: make([]Checkable, 0),
	}
}

func (c *Checker) Add(item Checkable) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items = append(c.items, item)
}

func (c *Checker) String() string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var ids []string
	for _, item := range c.items {
		ids = append(ids, item.GetID())
	}

	return strings.Join(ids, ", ")
}

func (c *Checker) Check() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, item := range c.items {
		if !item.Health() {
			fmt.Println(item.GetID(), "не работает")
		}
	}
}

func (c *Checker) Run() {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		c.Check()
	}
}

func (c *Checker) Stop() {
	fmt.Println("Проверки остановлены")
}
