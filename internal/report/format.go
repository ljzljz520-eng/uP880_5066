package report

import (
	"encoding/json"
	"vendor-permission/internal/domain"
)

func JSON(items []domain.Record) (string, error) {
	b, e := json.MarshalIndent(items, "", "  ")
	return string(b), e
}
func Header() []string { return []string{"id", "supplier", "warehouse", "permission", "status"} }
