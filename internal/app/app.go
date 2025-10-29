package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"tusur-forms/internal/config"

	"github.com/go-chi/chi/v5"
)

type formHandler struct{}

func Run() {
	ctx := context.Background()
	provider := &config.EnvProvider{}
	cfg, err := provider.NewConfig()
	if err != nil {
		panic(err)
	}
	router := chi.NewRouter()
	rs := formHandler{}
	router.Route("/api/v1/forms", func(r chi.Router) {
		// GET /api/v1/forms/{id}
		r.Get("/{id}", rs.GetFormByID)
		// POST /api/v1/forms
		r.Post("/yu", rs.CreateForm)
	})

	// Группировка маршрутов для Answers
	router.Route("/api/v1/answers", func(r chi.Router) {
		// POST /api/v1/answers (для отправки ответов)
		r.Post("/", rs.SubmitAnswer)
	})

	// --- 5. Запуск HTTP-Сервера ---
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func (rs formHandler) GetFormByID(w http.ResponseWriter, r *http.Request) {

}

func (rs formHandler) CreateForm(w http.ResponseWriter, r *http.Request) {

}

func (rs formHandler) SubmitAnswer(w http.ResponseWriter, r *http.Request) {

}
