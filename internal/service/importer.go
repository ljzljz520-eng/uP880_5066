package service

import (
	"strings"
	"vendor-permission/internal/domain"
)

func (r *Registry) Import(lines []string) ([]domain.Record, error) {
	out := []domain.Record{}
	for n, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) < 5 {
			continue
		}
		x, e := r.Register(parts[0], parts[1], parts[2], parts[3], parts[4])
		if e != nil {
			return out, e
		}
		out = append(out, x)
		if n > 1000 {
			break
		}
	}
	return out, nil
}
func (r *Registry) BulkValidate(items []domain.Record) []string {
	bad := []string{}
	for _, x := range items {
		if r.Validate(x) != nil {
			bad = append(bad, x.ID)
		}
	}
	return bad
}
