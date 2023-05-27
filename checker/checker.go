package checker

import (
	"fmt"
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

type Checker struct {
	items []Checkable
}

func NewChecker() *Checker {
	return &Checker{}
}

func (c *Checker) Add(item Checkable) {
	c.items = append(c.items, item)
}

func (c *Checker) String() string {
	var result string
	for _, item := range c.items {
		result += item.GetID() + "\n"
	}
	return result
}

func (c *Checker) Check() {
	for _, item := range c.items {
		if !item.Health() {
			fmt.Println(item.GetID() + " is not working")
		}
	}
}
