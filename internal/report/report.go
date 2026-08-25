package report

import (
	"fmt"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/storage"
)

type Reporter struct{ Store *storage.Store }

func New(s *storage.Store) *Reporter { return &Reporter{Store: s} }
func (r *Reporter) Export(f domain.Filter) (string, error) {
	p, e := r.Store.ListRecords(f)
	if e != nil {
		return "", e
	}
	out := "id,supplier,warehouse,permission,status\n"
	for _, x := range p.Items {
		out += fmt.Sprintf("%s,%s,%s,%s,%s\n", x.ID, x.Supplier, x.Warehouse, x.Permission, x.Status)
	}
	return out, nil
}
func (r *Reporter) Stats() map[string]int {
	return map[string]int{"records": r.Store.Count("records"), "audits": r.Store.Count("audits"), "workflows": r.Store.Count("workflows")}
}
func (r *Reporter) Detail(id string) (domain.Record, error) { return r.Store.GetRecord(id) }
