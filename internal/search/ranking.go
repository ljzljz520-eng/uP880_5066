package search

import (
	"sort"
	"strings"
	"vendor-permission/internal/domain"
)

func Rank(items []domain.Record, term string) []domain.Record {
	out := append([]domain.Record{}, items...)
	sort.SliceStable(out, func(i, j int) bool { return strings.Contains(strings.ToLower(out[i].Summary()), strings.ToLower(term)) })
	return out
}
func Tokens(term string) []string { return strings.Fields(strings.ToLower(term)) }
