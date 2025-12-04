package transport

import (
	"encoding/json"
	"net/http"
	"tusur-forms/internal/transport/dto"
)

// GetQuestions
// @Summary Get questions
// @Description Get all questions
// @Accept json
// @Produce json
// @Success 200 {array} dto.ResponseQuestion
// @Failure 500 {string} string "Internal error"
// @Router /questions [get]
func (o *Orchestrator) GetQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	questions, err := o.FormsOrchestrator.GetQuestions()
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}
	var qs = make([]*dto.ResponseQuestion, 0, len(questions))
	for _, q := range questions {
		qs = append(qs, q.ToDTO())
	}
	json.NewEncoder(w).Encode(qs)
}

// GetForm
// @Summary Get form
// @Description Get form detail information
// @Accept json
// @Produce json
// @Param form_id query string true "Form External ID"
// @Success 200 {object} dto.ResponseForm
// @Failure 404 {string} string "Not found error"
// @Failure 500 {string} string "Internal error"
// @Router /form [get]
func (o *Orchestrator) GetForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryVals := r.URL.Query()
	id := queryVals.Get("form_id")
	form, err := o.FormsOrchestrator.GetForm(id, true)
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}
	if form == nil {
		http.Error(w, "Form not found", 404)
		return
	}
	f := form.ToDTO()
	json.NewEncoder(w).Encode(f)
}

// GetForms
// @Summary Get forms
// @Description Get all forms
// @Accept json
// @Produce json
// @Success 200 {array} dto.ResponseForm
// @Failure 500 {string} string "Internal error"
// @Router /forms [get]
func (o *Orchestrator) GetForms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	forms, err := o.FormsOrchestrator.GetForms()
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}
	var fs = make([]*dto.ResponseForm, 0, len(forms))
	for _, f := range forms {
		fs = append(fs, f.ToDTO())
	}
	json.NewEncoder(w).Encode(fs)
}
