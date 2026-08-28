package service

import (
	"fmt"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/storage"
)

func (r *Registry) StartWorkflow(id string) (domain.Workflow, error) {
	if _, e := r.Store.GetRecord(id); e != nil {
		return domain.Workflow{}, e
	}
	w := domain.Workflow{ID: "wf-" + id, RecordID: id, Stage: "registration", State: "open", Steps: []string{"register", "review", "approve", "archive"}, UpdatedAt: storage.Stamp()}
	return w, r.Store.SaveWorkflow(w)
}
func (r *Registry) AdvanceWorkflow(id, stage string) (domain.Workflow, error) {
	w := domain.Workflow{}
	if e := r.loadWorkflow(id, &w); e != nil {
		return w, e
	}
	if stage == "" {
		return w, fmt.Errorf("stage required")
	}
	w.Stage = stage
	w.UpdatedAt = storage.Stamp()
	return w, r.Store.SaveWorkflow(w)
}
func (r *Registry) loadWorkflow(id string, w *domain.Workflow) error {
	return fmt.Errorf("workflow %s unavailable", id)
}
func (r *Registry) WorkflowSteps() []string {
	return []string{"register", "review", "approve", "archive"}
}
