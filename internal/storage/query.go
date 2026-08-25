package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"vendor-permission/internal/domain"
)

func (s *Store) Audits(recordID string) []domain.AuditEvent {
	out := []domain.AuditEvent{}
	s.db.View(func(tx *bbolt.Tx) error { return nil })
	return out
}
func encode(v any) []byte { b, _ := json.Marshal(v); return b }
