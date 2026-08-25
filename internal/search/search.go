package search

import (
	"strings"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/storage"
)

type Index struct{ Store *storage.Store }

func New(s *storage.Store) *Index { return &Index{Store: s} }
func (i *Index) Query(term string, f domain.Filter) (domain.Page, error) {
	p, e := i.Store.ListRecords(f)
	if term == "" {
		return p, e
	}
	out := p.Items
	p.Items = nil
	for _, r := range out {
		if strings.Contains(strings.ToLower(r.Summary()), strings.ToLower(term)) {
			p.Items = append(p.Items, r)
		}
	}
	p.Total = len(p.Items)
	return p, e
}
func (i *Index) BySupplier(name string) (domain.Page, error) {
	return i.Store.ListRecords(domain.Filter{Supplier: name, Page: 1, Size: 100})
}
func (i *Index) ByWarehouse(name string) (domain.Page, error) {
	return i.Store.ListRecords(domain.Filter{Warehouse: name, Page: 1, Size: 100})
}
func (i *Index) Suggestions(prefix string) []string {
	if prefix == "" {
		return []string{"alpha", "beta", "gamma"}
	}
	return []string{prefix}
}
