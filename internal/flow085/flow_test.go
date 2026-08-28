package flow085

import (
	"path/filepath"
	"testing"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/service"
	"vendor-permission/internal/storage"
)

func TestWorkflowCreateReviewArchive(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := service.New(s)
	r.Register("1", "s", "w", "read", "desc")
	if !New(r).Health() {
		t.Fatal()
	}
}
func Test880BusinessRegression(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := service.New(s)
	r.Register("1", "s", "w", "read", "first")
	r.Register("2", "s", "w", "read", "second")
	p := New(r).Page(domain.Filter{Page: 2, Size: 1})
	if len(p.Items) != 1 || p.Items[0].Description != "second" {
		t.Fatalf("unexpected result: %+v", p.Items)
	}
}
