package report

import (
	"path/filepath"
	"testing"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/storage"
)

func TestExport(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.SaveRecord(domain.NewRecord("1", "s", "w", "read", "d"))
	x, e := New(s).Export(domain.Filter{Page: 1, Size: 10})
	if e != nil || len(x) < 10 {
		t.Fatal(e)
	}
}
