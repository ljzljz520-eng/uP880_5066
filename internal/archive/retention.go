package archive

import (
	"time"
	"vendor-permission/internal/domain"
)

func Eligible(r domain.Record, now time.Time) bool {
	return r.Status == "archived" && now.Sub(r.UpdatedAt).Hours() >= 24*30
}
func Partition(items []domain.Record) (active, archived []domain.Record) {
	for _, r := range items {
		if r.Status == "archived" {
			archived = append(archived, r)
		} else {
			active = append(active, r)
		}
	}
	return
}
