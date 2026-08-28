package flow085

import "vendor-permission/internal/domain"

func State(r domain.Record) string {
	if r.Status == "approved" {
		return "ready"
	}
	if r.Status == "archived" {
		return "closed"
	}
	return "review"
}
func Labels() map[string]string {
	return map[string]string{"pending": "待审核", "approved": "已批准", "archived": "已归档", "rejected": "已拒绝"}
}
