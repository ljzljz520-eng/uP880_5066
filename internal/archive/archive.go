package archive

import (
	"fmt"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/storage"
)

type Manager struct{ Store *storage.Store }

func New(s *storage.Store) *Manager { return &Manager{Store: s} }
func (m *Manager) Archive(id, actor string) (domain.Record, error) {
	r, e := m.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if !r.CanArchive() {
		return r, fmt.Errorf("cannot archive")
	}
	r.Status = "archived"
	r.Version++
	r.UpdatedAt = storage.Stamp()
	if e = m.Store.SaveRecord(r); e != nil {
		return r, e
	}
	return r, m.Store.SaveAudit(domain.AuditEvent{ID: id + "-archive", RecordID: id, Action: "archive", Actor: actor, At: storage.Stamp()})
}
func (m *Manager) Restore(id, actor string) (domain.Record, error) {
	r, e := m.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if r.Status != "archived" {
		return r, fmt.Errorf("not archived")
	}
	r.Status = "approved"
	r.Version++
	r.UpdatedAt = storage.Stamp()
	e = m.Store.SaveRecord(r)
	return r, e
}
func (m *Manager) Archived(f domain.Filter) (domain.Page, error) {
	f.Status = "archived"
	return m.Store.ListRecords(f)
}
