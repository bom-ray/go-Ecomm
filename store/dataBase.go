package store

import "github.com/bom-ray/goEcomm/types"

type DataBase interface {
	CreateUser(types.User) types.User
	//show single user
	ListUsers() []types.User
	ShowUserCart(string) types.Cart
	AddToCart(types.AddCart)
	AddOrder(string)
	ListOrders() []types.Orders
	AddProduct(types.Product)
	DeleteProduct(string)
	EditProduct(types.Product)
	ListProducts() []types.Product
	//add category
	AddCategory(types.Category)
	ListProductsByCat(string) []types.Product
	ListCategory() []types.Category
}
