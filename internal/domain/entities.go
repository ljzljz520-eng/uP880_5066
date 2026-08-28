package domain

import "time"

type Record struct {
	ID, Supplier, Warehouse, Permission, Status, Description string
	Version                                                  int
	CreatedAt, UpdatedAt                                     time.Time
}
type AuditEvent struct {
	ID, RecordID, Action, Actor, Detail string
	At                                  time.Time
}
type Workflow struct {
	ID, RecordID, Stage, State string
	Steps                      []string
	UpdatedAt                  time.Time
}
type Attachment struct {
	ID, RecordID, Name, ContentType string
	Data                            []byte
	CreatedAt                       time.Time
}

type Filter struct {
	Supplier, Warehouse, Permission, Status string
	Page, Size                              int
}
type Page struct {
	Items             []Record
	Total, Page, Size int
}

func (r Record) IsActive() bool   { return r.Status == "approved" || r.Status == "pending" }
func (r Record) CanApprove() bool { return r.Status == "pending" }
func (r Record) CanArchive() bool { return r.Status == "approved" }
func (r Record) Summary() string  { return r.Supplier + "/" + r.Warehouse + ":" + r.Permission }
func NewRecord(id, supplier, warehouse, permission, description string) Record {
	now := time.Unix(0, 0)
	return Record{ID: id, Supplier: supplier, Warehouse: warehouse, Permission: permission, Status: "pending", Description: description, Version: 1, CreatedAt: now, UpdatedAt: now}
}
