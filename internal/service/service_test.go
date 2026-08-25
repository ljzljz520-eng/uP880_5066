package service

import (
	"path/filepath"
	"testing"
	"vendor-permission/internal/storage"
)

func TestRegisterAndUpdate(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := New(s)
	r.Register("1", "s", "w", "read", "d")
	x, e := r.Update("1", "new")
	if e != nil || x.Description != "new" {
		t.Fatal(e, x)
	}
}
func TestImport(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	r := New(s)
	xs, e := r.Import([]string{"1,s,w,read,d"})
	if e != nil || len(xs) != 1 {
		t.Fatal(e)
	}
}
