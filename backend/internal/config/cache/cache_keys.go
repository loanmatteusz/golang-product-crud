package config_cache

import "fmt"

var CacheKeys = struct {
	ProductsAll   string
	ProductByID   func(id string) string
	CategoriesAll func(page, limit int, filter string) string
	CategoryByID  func(id string) string
}{
	ProductsAll: "products:all",
	ProductByID: func(id string) string { return "products:" + id },
	CategoriesAll: func(page, limit int, filter string) string {
		return fmt.Sprintf("categories:all:page=%d:limit=%d:filter=%s", page, limit, filter)
	},
	CategoryByID: func(id string) string { return "categories:" + id },
}
