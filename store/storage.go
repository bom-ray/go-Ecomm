package store

import (
	"sort"

	"github.com/bom-ray/goEcomm/types"
)

type Storage struct {
	UserIdToUser map[string]types.User
	UserCart     map[string]types.Cart
	Orders       []types.Orders
	Products     map[string]types.Product
	Category     []types.Category
	CatToProds   map[string][]types.Product
}

func NewStorage(u map[string]types.User,
	cart map[string]types.Cart,
	orders []types.Orders,
	product map[string]types.Product,
	cat []types.Category,
	catToProds map[string][]types.Product,
) DataBase {
	return &Storage{
		UserIdToUser: u,
		UserCart:     cart,
		Orders:       orders,
		Products:     product,
		Category:     cat,
		CatToProds:   catToProds,
	}
}

// func (s *Storage) String() string {
// 	return fmt.Sprintf("[Users: %]\n", s.Users)
// }

func (s *Storage) ListOrders() []types.Orders {
	orders := s.Orders
	return orders
}

func (s *Storage) ListCategory() []types.Category {
	cat := s.Category
	return cat
}

func (s *Storage) ListProductsByCat(catId string) []types.Product {
	return s.CatToProds[catId]
}

func (s *Storage) AddCategory(cat types.Category) {
	s.Category = append(s.Category, cat)
}

func (s *Storage) ListProducts() []types.Product {
	productMap := s.Products
	productsArray := make([]types.Product, 0)

	for k, _ := range productMap {
		productsArray = append(productsArray, s.Products[k])
	}

	sort.Slice(productsArray, func(i, j int) bool {
		return productsArray[i].Name > productsArray[j].Name
	})

	return productsArray
}

func (s *Storage) EditProduct(p types.Product) {
	s.Products[p.Id] = p

	//loop through category of product
	//use catToProds to edit cat to product array
	for k, _ := range p.Category {
		catName := p.Category[k]
		productList := s.CatToProds[catName]
		for i, _ := range productList {
			if productList[i].Id == p.Id {
				//productList[i] = p
				s.CatToProds[catName][i] = p
				//not so sure about break
				break
			}
		}
	}
}

func (s *Storage) DeleteProduct(id string) {
	delete(s.Products, id)
}

func (s *Storage) AddProduct(p types.Product) {
	s.Products[p.Id] = p

	//loop through category of product
	//use catToProds to add cat to product array
	for k, _ := range p.Category {
		catName := p.Category[k]
		s.CatToProds[catName] = append(s.CatToProds[catName], p)
	}

}

func (s *Storage) AddOrder(UserId string) {

	total := s.UserCart[UserId].Total

	var orderList []types.CartItems
	items := s.UserCart[UserId].Items
	for k, _ := range items {
		list := types.CartItems{
			ProductId: items[k].ProductId,
			Qty:       items[k].Qty,
			Amount:    items[k].Amount,
		}
		orderList = append(orderList, list)
		//total += float64(items[k].Qty) * items[k].Amount
	}

	order := types.Orders{
		UserId:     UserId,
		OrdersList: orderList,
		Total:      total,
	}

	s.Orders = append(s.Orders, order)

	//fmt.Println(s.Orders)
}

func (s *Storage) AddToCart(addCart types.AddCart) {
	cartDetails := s.UserCart[addCart.UserId]
	cartItems := cartDetails.Items[addCart.ProductId]
	cartItems.Qty += addCart.Qty
	cartItems.Amount = addCart.Amount
	cartItems.ProductId = addCart.ProductId

	s.UserCart[addCart.UserId].Items[addCart.ProductId] = cartItems

	var total float64
	//loop through the map and add total
	items := s.UserCart[addCart.UserId].Items

	for k, _ := range items {
		total += float64(items[k].Qty) * items[k].Amount
	}

	//fmt.Println(total)
	//s.UserCart[addCart.UserId].Total = total
	//s.UserCart[addCart.UserId].Total
	//carts := make(map[string]types.CartItems)

	s.UserCart[addCart.UserId] = types.Cart{
		UserId: addCart.UserId,
		Items:  items,
		Total:  total,
	}

	//fmt.Println(s.UserCart[addCart.UserId])

}

func (s *Storage) ShowUserCart(uid string) types.Cart {
	cart := s.UserCart[uid]
	return cart
}

func (s *Storage) CreateUser(u types.User) types.User {
	s.UserIdToUser[u.Id] = u
	//s.Users = append(s.Users, u)
	s.UserCart[u.Id] = types.Cart{
		UserId: u.Id,
		Items:  make(map[string]types.CartItems),
		Total:  0.0,
	}
	return u
}

func (s *Storage) ListUsers() []types.User {
	userMap := s.UserIdToUser
	usersArray := make([]types.User, 0)

	for k, _ := range userMap {
		usersArray = append(usersArray, userMap[k])
	}

	return usersArray
}
