package internal

import (
	"path/filepath"
	"testing"
	"vendor-permission/internal/service"
	"vendor-permission/internal/storage"
)

func TestWorkflowSearchUpdatePublish(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := service.New(s)
	if _, e := r.Register("1", "s", "w", "read", "d"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowImportReport(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := service.New(s)
	if _, e := r.Import([]string{"1,s,w,read,d"}); e != nil {
		t.Fatal(e)
	}
}
