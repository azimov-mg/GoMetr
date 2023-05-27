package main

import (
	"fmt"
	"github.com/azimov-mg/gometr/checker"
	"github.com/azimov-mg/gometr/gometr"
	"time"
)

func main() {
	g1 := gometr.NewGoMetrClient("https://example.com", 5*time.Second)
	g2 := gometr.NewGoMetrClient("https://api.example.com", 3*time.Second)

	c := checker.NewChecker()
	c.Add(g1)
	c.Add(g2)

	fmt.Print(c)

	c.Check()

	fmt.Println("Program exited")
}
