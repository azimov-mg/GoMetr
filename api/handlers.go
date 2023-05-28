package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/azimov-mg/gometr/internal/db"
)

// MetricsHandler обрабатывает запрос на получение списка метрик.
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметров из запроса
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	timeFromStr := r.URL.Query().Get("time_from")
	timeToStr := r.URL.Query().Get("time_to")
	name := r.URL.Query().Get("name")

	// Преобразование параметров в нужный формат
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	timeFrom, _ := time.Parse(time.RFC3339, timeFromStr)
	timeTo, _ := time.Parse(time.RFC3339, timeToStr)

	// Получение списка метрик из базы данных
	storage := db.NewStorage() // Предполагается, что класс Storage уже реализован
	metrics := storage.GetMetricsList(limit, offset, timeFrom, timeTo, name)

	// Преобразование списка метрик в JSON
	response, _ := json.Marshal(metrics)

	// Установка заголовков и отправка ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Другие обработчики API
