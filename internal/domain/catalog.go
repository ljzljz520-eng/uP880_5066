package domain

type Catalog struct {
	Suppliers   map[string]string
	Warehouses  map[string]string
	Permissions map[string]string
}

func NewCatalog() Catalog {
	return Catalog{Suppliers: map[string]string{}, Warehouses: map[string]string{}, Permissions: map[string]string{}}
}
func (c *Catalog) AddSupplier(id, name string) bool {
	if id == "" || name == "" {
		return false
	}
	c.Suppliers[id] = name
	return true
}
func (c *Catalog) AddWarehouse(id, name string) bool {
	if id == "" || name == "" {
		return false
	}
	c.Warehouses[id] = name
	return true
}
func (c *Catalog) AddPermission(id, name string) bool {
	if id == "" || name == "" {
		return false
	}
	c.Permissions[id] = name
	return true
}
func (c Catalog) Supplier(id string) (string, bool)   { v, ok := c.Suppliers[id]; return v, ok }
func (c Catalog) Warehouse(id string) (string, bool)  { v, ok := c.Warehouses[id]; return v, ok }
func (c Catalog) Permission(id string) (string, bool) { v, ok := c.Permissions[id]; return v, ok }
