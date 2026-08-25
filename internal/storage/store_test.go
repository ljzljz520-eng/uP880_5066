package storage

import (
	"path/filepath"
	"testing"
	"vendor-permission/internal/domain"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "x.db")
	s, _ := Open(p)
	s.SaveRecord(domain.NewRecord("r1", "s", "w", "read", "d"))
	s.Close()
	s, _ = Open(p)
	defer s.Close()
	if _, e := s.GetRecord("r1"); e != nil {
		t.Fatal(e)
	}
}
func TestListPaging(t *testing.T) {
	s, _ := Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	for i := 0; i < 3; i++ {
		s.SaveRecord(domain.NewRecord(string(rune('a'+i)), "s", "w", "read", "d"))
	}
	p, _ := s.ListRecords(domain.Filter{Page: 2, Size: 2})
	if len(p.Items) != 1 {
		t.Fatal(len(p.Items))
	}
}
