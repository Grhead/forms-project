package transport

import "tusur-forms/internal/services/orchectrators"

type Orchestrator struct {
	*orchectrators.FormsOrchestrator
}

func NewOrchestrator(o *orchectrators.FormsOrchestrator) *Orchestrator {
	return &Orchestrator{
		FormsOrchestrator: o,
	}
}
