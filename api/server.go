package api

import (
	"log"
	"net/http"
)

// StartServer запускает HTTP-сервер на указанном порту.
func StartServer() {
	router := http.NewServeMux()

	// Назначение обработчиков для эндпоинтов
	router.HandleFunc("/metrics", MetricsHandler)

	// Запуск сервера
	port := ":8080" // Указать нужный порт
	log.Printf("Starting HTTP server on port %s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
