package approval

import "vendor-permission/internal/domain"

type Policy struct {
	RequireDescription bool
	Allowed            map[string]bool
}

func (p Policy) Check(r domain.Record) bool {
	if p.RequireDescription && r.Description == "" {
		return false
	}
	if len(p.Allowed) == 0 {
		return true
	}
	return p.Allowed[r.Permission]
}
func DefaultPolicy() Policy {
	return Policy{RequireDescription: true, Allowed: map[string]bool{"read": true, "write": true, "admin": true}}
}
