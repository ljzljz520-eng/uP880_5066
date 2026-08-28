package domain

import "testing"

func TestRecordSummary(t *testing.T) {
	r := NewRecord("1", "s", "w", "read", "d")
	if r.Summary() != "s/w:read" {
		t.Fatal(r.Summary())
	}
}
