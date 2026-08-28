package archive

import (
	"path/filepath"
	"testing"
	"vendor-permission/internal/approval"
	"vendor-permission/internal/service"
	"vendor-permission/internal/storage"
)

func TestArchiveRestore(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := service.New(s)
	r.Register("1", "s", "w", "read", "d")
	approval.New(s).Approve("1", "u")
	m := New(s)
	m.Archive("1", "u")
	x, e := m.Restore("1", "u")
	if e != nil || x.Status != "approved" {
		t.Fatal(e)
	}
}
