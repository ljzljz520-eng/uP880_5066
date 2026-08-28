package approval

import (
	"path/filepath"
	"testing"
	"vendor-permission/internal/service"
	"vendor-permission/internal/storage"
)

func TestApprove(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := service.New(s)
	r.Register("1", "s", "w", "read", "d")
	a := New(s)
	x, e := a.Approve("1", "u")
	if e != nil || x.Status != "approved" {
		t.Fatal(e)
	}
}
