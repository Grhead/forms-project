package transport

import (
	"encoding/json"
	"net/http"
	"tusur-forms/internal/domain"
	"tusur-forms/internal/transport/dto"
)

func (o *Orchestrator) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var newQuestion dto.RequestQuestion

	err := json.NewDecoder(r.Body).Decode(&newQuestion)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var pa = make([]*domain.PossibleAnswer, 0, len(newQuestion.PossibleAnswers))
	for _, p := range newQuestion.PossibleAnswers {
		pa = append(pa, &domain.PossibleAnswer{
			Content: p.Content,
		})
	}
	question, err := o.FormsOrchestrator.CheckoutQuestion(&domain.Question{
		Title:           newQuestion.Title,
		Description:     newQuestion.Description,
		Type:            domain.QuestionType{Title: domain.QuestionTypeTitles(newQuestion.Type)},
		IsRequired:      newQuestion.IsRequired,
		PossibleAnswers: pa,
	})
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}

func (o *Orchestrator) CreateForm(w http.ResponseWriter, r *http.Request) {
	var newForm dto.RequestForm

	err := json.NewDecoder(r.Body).Decode(&newForm)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	var questions = make([]*domain.Question, 0, len(newForm.Questions))
	for _, q := range newForm.Questions {
		var pa = make([]*domain.PossibleAnswer, 0, len(q.PossibleAnswers))
		for _, p := range q.PossibleAnswers {
			pa = append(pa, &domain.PossibleAnswer{
				Content: p.Content,
			})
		}
		questions = append(questions, &domain.Question{
			Title:           q.Title,
			Description:     q.Description,
			Type:            domain.QuestionType{Title: domain.QuestionTypeTitles(q.Type)},
			IsRequired:      q.IsRequired,
			PossibleAnswers: pa,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	form, err := o.FormsOrchestrator.CheckoutForm(newForm.Title, newForm.DocumentTitle, questions)
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fDto := form.ToDTO()
	json.NewEncoder(w).Encode(fDto)
}

func (o *Orchestrator) GenerateXlsx(w http.ResponseWriter, r *http.Request) {

}
