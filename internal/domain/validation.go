package domain

import "fmt"

func ValidateFilter(f Filter) error {
	if f.Page < 0 {
		return fmt.Errorf("negative page")
	}
	if f.Size < 0 {
		return fmt.Errorf("negative size")
	}
	return nil
}
func (f Filter) Offset() int {
	if f.Page < 1 {
		return 0
	}
	if f.Size < 1 {
		return 0
	}
	return (f.Page - 1) * f.Size
}
func (p Page) HasNext() bool { return p.Page*p.Size < p.Total }
func (p Page) Empty() bool   { return len(p.Items) == 0 }
