package db

import (
	"database/sql"
	"github.com/azimov-mg/gometr/internal/models"
	"time"
)

// Storage представляет собой класс для работы с базой данных.
type Storage struct {
	db *sql.DB
}

// NewStorage создает новый экземпляр Storage.
func NewStorage() *Storage {
	// Установка соединения с базой данных
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=gometr password=secret")
	if err != nil {
		panic(err)
	}

	return &Storage{
		db: db,
	}
}

// GetMetricsList возвращает список метрик с возможностью фильтрации.
func (s *Storage) GetMetricsList(limit, offset int, timeFrom, timeTo time.Time, name string) []*models.Metric {
	// Здесь должна быть логика выполнения запроса к базе данных и возврата списка метрик
	// Предполагается, что метод возвращает []*models.Metric
	return nil
}

// Другие методы работы с базой данных
