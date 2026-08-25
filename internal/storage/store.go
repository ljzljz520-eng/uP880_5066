package storage

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sort"
	"time"
	"vendor-permission/internal/domain"
)

var buckets = []string{"records", "audits", "workflows", "attachments"}

type Store struct{ db *bbolt.DB }

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, e := tx.CreateBucketIfNotExists([]byte(n)); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func put(tx *bbolt.Tx, b string, key string, v any) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket([]byte(b)).Put([]byte(key), data)
}
func get(tx *bbolt.Tx, b, key string, v any) error {
	raw := tx.Bucket([]byte(b)).Get([]byte(key))
	if raw == nil {
		return errors.New("not found")
	}
	return json.Unmarshal(raw, v)
}
func (s *Store) SaveRecord(r domain.Record) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "records", r.ID, r) })
}
func (s *Store) GetRecord(id string) (domain.Record, error) {
	var r domain.Record
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, "records", id, &r) })
	return r, e
}
func (s *Store) ListRecords(f domain.Filter) (domain.Page, error) {
	var out []domain.Record
	e := s.db.View(func(tx *bbolt.Tx) error {
		c := tx.Bucket([]byte("records")).Cursor()
		for _, v := c.First(); v != nil; _, v = c.Next() {
			var r domain.Record
			if json.Unmarshal(v, &r) != nil {
				continue
			}
			if f.Supplier != "" && r.Supplier != f.Supplier {
				continue
			}
			if f.Warehouse != "" && r.Warehouse != f.Warehouse {
				continue
			}
			if f.Permission != "" && r.Permission != f.Permission {
				continue
			}
			if f.Status != "" && r.Status != f.Status {
				continue
			}
			out = append(out, r)
		}
		return nil
	})
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	if f.Size <= 0 {
		f.Size = 10
	}
	if f.Page <= 0 {
		f.Page = 1
	}
	start := (f.Page - 1) * f.Size
	if start > len(out) {
		start = len(out)
	}
	end := start + f.Size
	if end > len(out) {
		end = len(out)
	}
	return domain.Page{Items: out[start:end], Total: len(out), Page: f.Page, Size: f.Size}, e
}
func (s *Store) SaveAudit(a domain.AuditEvent) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "audits", a.ID, a) })
}
func (s *Store) SaveWorkflow(w domain.Workflow) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "workflows", w.ID, w) })
}
func (s *Store) SaveAttachment(a domain.Attachment) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, "attachments", a.ID, a) })
}
func (s *Store) Count(bucket string) int {
	n := 0
	s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b != nil {
			n = b.Stats().KeyN
		}
		return nil
	})
	return n
}
func Stamp() time.Time { return time.Unix(0, 0) }
