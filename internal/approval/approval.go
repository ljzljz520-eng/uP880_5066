package approval

import (
	"fmt"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/storage"
)

type Engine struct{ Store *storage.Store }

func New(s *storage.Store) *Engine { return &Engine{Store: s} }
func (e *Engine) Approve(id, actor string) (domain.Record, error) {
	r, er := e.Store.GetRecord(id)
	if er != nil {
		return r, er
	}
	if !r.CanApprove() {
		return r, fmt.Errorf("not pending")
	}
	r.Status = "approved"
	r.Version++
	r.UpdatedAt = storage.Stamp()
	if er = e.Store.SaveRecord(r); er != nil {
		return r, er
	}
	return r, e.Store.SaveAudit(domain.AuditEvent{ID: id + "-approve", RecordID: id, Action: "approve", Actor: actor, At: storage.Stamp()})
}
func (e *Engine) Reject(id, actor, detail string) (domain.Record, error) {
	r, er := e.Store.GetRecord(id)
	if er != nil {
		return r, er
	}
	if r.Status != "pending" {
		return r, fmt.Errorf("not pending")
	}
	r.Status = "rejected"
	r.Version++
	r.UpdatedAt = storage.Stamp()
	er = e.Store.SaveRecord(r)
	if er != nil {
		return r, er
	}
	return r, e.Store.SaveAudit(domain.AuditEvent{ID: id + "-reject", RecordID: id, Action: "reject", Actor: actor, Detail: detail, At: storage.Stamp()})
}
func (e *Engine) Pending(f domain.Filter) (domain.Page, error) {
	f.Status = "pending"
	return e.Store.ListRecords(f)
}
func (e *Engine) IsAllowed(r domain.Record) bool { return r.Status == "approved" }
