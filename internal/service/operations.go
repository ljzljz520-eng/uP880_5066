package service

import (
	"fmt"
	"strings"
	"vendor-permission/internal/domain"
	"vendor-permission/internal/storage"
)

type OperationLog struct{ Messages []string }

func NewOperationLog() *OperationLog { return &OperationLog{Messages: []string{}} }
func (l *OperationLog) Add(m string) {
	if strings.TrimSpace(m) != "" {
		l.Messages = append(l.Messages, m)
	}
}
func (l *OperationLog) Last() string {
	if len(l.Messages) == 0 {
		return ""
	}
	return l.Messages[len(l.Messages)-1]
}
func (l *OperationLog) Count() int { return len(l.Messages) }
func (r *Registry) EnsureStatus(id, status string) error {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if status == "" {
		return fmt.Errorf("status required")
	}
	x.Status = status
	x.Version++
	return r.Store.SaveRecord(x)
}
func (r *Registry) Touch(id string) error {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return e
	}
	x.UpdatedAt = storage.Stamp()
	return r.Store.SaveRecord(x)
}
func (r *Registry) Clone(id, newID string) (domain.Record, error) {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return x, e
	}
	if newID == "" {
		return x, fmt.Errorf("new id required")
	}
	x.ID = newID
	x.Status = "pending"
	x.Version = 1
	return x, r.Store.SaveRecord(x)
}
func (r *Registry) Compare(a, b string) bool {
	x, e := r.Store.GetRecord(a)
	if e != nil {
		return false
	}
	y, e := r.Store.GetRecord(b)
	if e != nil {
		return false
	}
	return x.Supplier == y.Supplier && x.Permission == y.Permission
}
func (r *Registry) SetPermission(id, p string) error {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if p == "" {
		return fmt.Errorf("permission required")
	}
	x.Permission = p
	x.Version++
	return r.Store.SaveRecord(x)
}
func (r *Registry) SetWarehouse(id, w string) error {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if w == "" {
		return fmt.Errorf("warehouse required")
	}
	x.Warehouse = w
	x.Version++
	return r.Store.SaveRecord(x)
}
func (r *Registry) SetSupplier(id, s string) error {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if s == "" {
		return fmt.Errorf("supplier required")
	}
	x.Supplier = s
	x.Version++
	return r.Store.SaveRecord(x)
}
func (r *Registry) SetDescription(id, d string) error {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return e
	}
	x.Description = d
	x.Version++
	return r.Store.SaveRecord(x)
}
func (r *Registry) IsMutable(id string) bool {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return false
	}
	return x.Status != "archived"
}
func (r *Registry) IsReviewable(id string) bool {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return false
	}
	return x.Status == "pending"
}
func (r *Registry) IsPublishable(id string) bool {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return false
	}
	return x.Status == "approved"
}
func (r *Registry) IsArchived(id string) bool {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return false
	}
	return x.Status == "archived"
}
func (r *Registry) Require(id string) (domain.Record, error) {
	x, e := r.Store.GetRecord(id)
	if e != nil {
		return x, fmt.Errorf("record %s: %w", id, e)
	}
	return x, nil
}
func (r *Registry) Summary(id string) string {
	x, e := r.Require(id)
	if e != nil {
		return ""
	}
	return x.Summary()
}
func (r *Registry) Status(id string) string {
	x, e := r.Require(id)
	if e != nil {
		return "unknown"
	}
	return x.Status
}
func (r *Registry) Version(id string) int {
	x, e := r.Require(id)
	if e != nil {
		return 0
	}
	return x.Version
}
func (r *Registry) ValidID(id string) bool         { return strings.TrimSpace(id) != "" && len(id) <= 64 }
func (r *Registry) ValidSupplier(s string) bool    { return strings.TrimSpace(s) != "" && len(s) <= 128 }
func (r *Registry) ValidWarehouse(s string) bool   { return strings.TrimSpace(s) != "" && len(s) <= 128 }
func (r *Registry) ValidPermission(s string) bool  { return s == "read" || s == "write" || s == "admin" }
func (r *Registry) ValidDescription(s string) bool { return len(s) <= 2048 }
func (r *Registry) ValidateAll(x domain.Record) error {
	if !r.ValidID(x.ID) {
		return fmt.Errorf("invalid id")
	}
	if !r.ValidSupplier(x.Supplier) {
		return fmt.Errorf("invalid supplier")
	}
	if !r.ValidWarehouse(x.Warehouse) {
		return fmt.Errorf("invalid warehouse")
	}
	if !r.ValidPermission(x.Permission) {
		return fmt.Errorf("invalid permission")
	}
	if !r.ValidDescription(x.Description) {
		return fmt.Errorf("invalid description")
	}
	return nil
}
func (r *Registry) Prepare(x domain.Record) domain.Record {
	if x.Version < 1 {
		x.Version = 1
	}
	if x.Status == "" {
		x.Status = "pending"
	}
	if x.CreatedAt.IsZero() {
		x.CreatedAt = storage.Stamp()
	}
	x.UpdatedAt = storage.Stamp()
	return x
}
func (r *Registry) SavePrepared(x domain.Record) error {
	y := r.Prepare(x)
	if e := r.ValidateAll(y); e != nil {
		return e
	}
	return r.Store.SaveRecord(y)
}
func (r *Registry) Upsert(x domain.Record) error {
	if _, e := r.Store.GetRecord(x.ID); e == nil {
		x.Version++
	}
	return r.SavePrepared(x)
}
func (r *Registry) Delete(id string) error {
	x, e := r.Require(id)
	if e != nil {
		return e
	}
	x.Status = "archived"
	return r.Store.SaveRecord(x)
}
func (r *Registry) Reopen(id string) error {
	x, e := r.Require(id)
	if e != nil {
		return e
	}
	if x.Status != "archived" {
		return fmt.Errorf("not archived")
	}
	x.Status = "pending"
	return r.Store.SaveRecord(x)
}
func (r *Registry) Count(f domain.Filter) int {
	p, e := r.Find(f)
	if e != nil {
		return 0
	}
	return p.Total
}
func (r *Registry) First(f domain.Filter) (domain.Record, error) {
	f.Page = 1
	f.Size = 1
	p, e := r.Find(f)
	if e != nil || len(p.Items) == 0 {
		return domain.Record{}, fmt.Errorf("empty")
	}
	return p.Items[0], nil
}
func (r *Registry) All(f domain.Filter) ([]domain.Record, error) {
	f.Page = 1
	f.Size = 10000
	p, e := r.Find(f)
	return p.Items, e
}
func (r *Registry) IDs(f domain.Filter) []string {
	xs, _ := r.All(f)
	out := []string{}
	for _, x := range xs {
		out = append(out, x.ID)
	}
	return out
}
func (r *Registry) Suppliers(f domain.Filter) []string {
	xs, _ := r.All(f)
	out := []string{}
	seen := map[string]bool{}
	for _, x := range xs {
		if !seen[x.Supplier] {
			seen[x.Supplier] = true
			out = append(out, x.Supplier)
		}
	}
	return out
}
func (r *Registry) Warehouses(f domain.Filter) []string {
	xs, _ := r.All(f)
	out := []string{}
	seen := map[string]bool{}
	for _, x := range xs {
		if !seen[x.Warehouse] {
			seen[x.Warehouse] = true
			out = append(out, x.Warehouse)
		}
	}
	return out
}
func (r *Registry) Permissions(f domain.Filter) []string {
	xs, _ := r.All(f)
	out := []string{}
	seen := map[string]bool{}
	for _, x := range xs {
		if !seen[x.Permission] {
			seen[x.Permission] = true
			out = append(out, x.Permission)
		}
	}
	return out
}
func (r *Registry) Statuses(f domain.Filter) []string {
	xs, _ := r.All(f)
	out := []string{}
	seen := map[string]bool{}
	for _, x := range xs {
		if !seen[x.Status] {
			seen[x.Status] = true
			out = append(out, x.Status)
		}
	}
	return out
}
