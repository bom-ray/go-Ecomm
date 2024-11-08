package main

import (
	"github.com/bom-ray/goEcomm/store"
	transport "github.com/bom-ray/goEcomm/transport"
	"github.com/bom-ray/goEcomm/types"
)

func main() {
	// usersDb := make(map[string]types.User)
	// cart := make(map[string]types.Cart)
	// orders := make([]types.Orders, 0)
	// products := make(map[string]types.Product)
	// category := make([]types.Category, 0)
	// catToProds := make(map[string][]types.Product)
	Db := store.NewStorage(
		make(map[string]types.User),
		make(map[string]types.Cart),
		make([]types.Orders, 0),
		make(map[string]types.Product),
		make([]types.Category, 0),
		make(map[string][]types.Product),
	)
	// Db := store.NewStorage(
	// 	usersDb, cart, orders, products, category, catToProds,
	// )
	server := transport.NewHttpTransport(Db)
	server.StartServer()
}
