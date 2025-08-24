package config_cache

import "fmt"

var CacheKeys = struct {
	ProductByID   func(id string) string
	CategoryByID  func(id string) string
	ProductsAll   func(page, limit int, filter string) string
	CategoriesAll func(page, limit int, filter string) string
}{
	ProductByID:  func(id string) string { return "products:" + id },
	CategoryByID: func(id string) string { return "categories:" + id },
	ProductsAll: func(page, limit int, filter string) string {
		return fmt.Sprintf("products:all:page=%d:limit=%d:filter=%s", page, limit, filter)
	},
	CategoriesAll: func(page, limit int, filter string) string {
		return fmt.Sprintf("categories:all:page=%d:limit=%d:filter=%s", page, limit, filter)
	},
}
