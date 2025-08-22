package config_cache

// CacheKeys mantém todas as chaves de cache da aplicação
var CacheKeys = struct {
	ProductsAll   string
	ProductByID   func(id string) string
	CategoriesAll string
	CategoryByID  func(id string) string
}{
	ProductsAll:   "products:all",
	ProductByID:   func(id string) string { return "products:" + id },
	CategoriesAll: "categories:all",
	CategoryByID:  func(id string) string { return "categories:" + id },
}
