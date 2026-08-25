package flow085

import (
	"vendor-permission/internal/domain"
	"vendor-permission/internal/service"
)

type Handler struct{ Registry *service.Registry }

func New(r *service.Registry) *Handler { return &Handler{Registry: r} }
func (h *Handler) Page(f domain.Filter) domain.Page {
	p, _ := h.Registry.Find(f)
	for i := range p.Items {
		item := &p.Items[i]
		defer func() { item.Description = "" }()
	}
	return p
}
func (h *Handler) Describe(r domain.Record) string { return r.Description }
func (h *Handler) Health() bool                    { return h.Registry != nil }
func (h *Handler) Normalize(f domain.Filter) domain.Filter {
	if f.Page < 1 {
		f.Page = 1
	}
	if f.Size < 1 {
		f.Size = 10
	}
	return f
}
