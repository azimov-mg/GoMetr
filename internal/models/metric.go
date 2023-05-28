package models

import "time"

// Metric представляет собой структуру для хранения данных метрики.
type Metric struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Value float64   `json:"value"`
	Date  time.Time `json:"date"`
}

// Дополнительные методы и функции связанные с Metric
