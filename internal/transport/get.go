package transport

import (
	"encoding/json"
	"net/http"
	"tusur-forms/internal/transport/dto"
)

func (o *Orchestrator) GetQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	questions, err := o.FormsOrchestrator.GetQuestions()
	if err != nil {
		http.Error(w, "Internal error", 500)
	}
	var qs = make([]*dto.Question, 0, len(questions))
	for _, q := range questions {
		qs = append(qs, q.ToDTO())
	}
	json.NewEncoder(w).Encode(qs)
}

func (o *Orchestrator) GetForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryVals := r.URL.Query()
	id := queryVals.Get("form_id")
	form, err := o.FormsOrchestrator.GetForm(id, true)
	if err != nil {
		return
	}
	f := form.ToDTO()
	json.NewEncoder(w).Encode(f)
}

func (o *Orchestrator) GetForms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	forms, err := o.FormsOrchestrator.GetForms()
	if err != nil {
		http.Error(w, "Internal error", 500)
	}
	var fs = make([]*dto.Form, 0, len(forms))
	for _, f := range forms {
		fs = append(fs, f.ToDTO())
	}
	json.NewEncoder(w).Encode(fs)
}
