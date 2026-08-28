package search

import (
	"path/filepath"
	"testing"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/storage"
)

func TestQuery(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.SaveRecord(domain.NewRecord("1", "Acme", "w", "read", "d"))
	p, e := New(s).Query("acme", domain.Filter{Page: 1, Size: 10})
	if e != nil || len(p.Items) != 1 {
		t.Fatal(e)
	}
}
