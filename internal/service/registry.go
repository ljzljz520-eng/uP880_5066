package service

import (
	"fmt"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/storage"
)

type Registry struct{ Store *storage.Store }

func New(s *storage.Store) *Registry { return &Registry{Store: s} }
func (r *Registry) Register(id, supplier, warehouse, permission, desc string) (domain.Record, error) {
	if id == "" || supplier == "" {
		return domain.Record{}, fmt.Errorf("id and supplier required")
	}
	x := domain.NewRecord(id, supplier, warehouse, permission, desc)
	if e := r.Store.SaveRecord(x); e != nil {
		return x, e
	}
	return x, nil
}
func (r *Registry) Find(f domain.Filter) (domain.Page, error) { return r.Store.ListRecords(f) }
func (r *Registry) Update(id, desc string) (domain.Record, error) {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return x, e
	}
	if x.Status == "archived" {
		return x, fmt.Errorf("archived")
	}
	x.Description = desc
	x.Version++
	x.UpdatedAt = storage.Stamp()
	e = r.Store.SaveRecord(x)
	return x, e
}
func (r *Registry) Validate(x domain.Record) error {
	if x.Supplier == "" || x.Warehouse == "" || x.Permission == "" {
		return fmt.Errorf("missing fields")
	}
	return nil
}
func (r *Registry) Duplicate(f domain.Filter) bool { p, _ := r.Find(f); return p.Total > 1 }
