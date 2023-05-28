package metrics

import (
	"github.com/azimov-mg/gometr/internal/db"
)

// MetricsProcessor представляет собой класс для сбора метрик.
type MetricsProcessor struct {
	storage *db.Storage
}

// NewMetricsProcessor создает новый экземпляр MetricsProcessor.
func NewMetricsProcessor(storage *db.Storage) *MetricsProcessor {
	return &MetricsProcessor{
		storage: storage,
	}
}

// Add собирает и сохраняет метрики в базу данных.
func (mp *MetricsProcessor) Add() {
	// Здесь должна быть логика снятия метрик и сохранения их в базу данных
	// Предполагается использование HTTP-клиента для опроса URL metrics
	// и парсинг полученных данных для получения метрик
	// Для сохранения метрик в базу данных используется метод Storage.Add()
}

// Другие методы MetricsProcessor
