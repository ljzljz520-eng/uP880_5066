package main

import (
	"fmt"
	"os"
	"vendor-permission/internal/service"
	"vendor-permission/internal/storage"
)

func main() {
	path := "permissions.db"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	s, e := storage.Open(path)
	if e != nil {
		panic(e)
	}
	defer s.Close()
	r := service.New(s)
	if s.Count("records") == 0 {
		r.Register("demo-1", "Northwind", "WH-1", "read", "demo permission")
	}
	fmt.Println("仓储供应商权限台 ready", s.Count("records"))
}
