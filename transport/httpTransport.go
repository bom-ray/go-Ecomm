package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bom-ray/goEcomm/store"
	"github.com/bom-ray/goEcomm/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type HttpTransport struct {
	Db store.DataBase
}

func NewHttpTransport(Db store.DataBase) *HttpTransport {
	return &HttpTransport{
		Db: Db,
	}
}

func (h *HttpTransport) StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/createUser", h.CreateUser).Methods("POST")
	router.HandleFunc("/users", h.ListUsers).Methods("GET")
	router.HandleFunc("/showUserCart/{userId}", h.ShowUserCart).Methods("GET")
	router.HandleFunc("/addToCart/{userId}", h.AddToCart).Methods("POST")
	router.HandleFunc("/addOrder/{userId}", h.AddOrder).Methods("POST")
	router.HandleFunc("/orders", h.ListOrders).Methods("GET")
	router.HandleFunc("/addProduct", h.AddProduct).Methods("POST")
	router.HandleFunc("/product/{productId}", h.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/product/{productId}", h.EditProduct).Methods("PUT")
	router.HandleFunc("/products", h.ListProducts).Methods("GET")
	router.HandleFunc("/addCategory", h.AddCategory).Methods("POST")
	router.HandleFunc("/category", h.ListCategory).Methods("GET")
	router.HandleFunc("/listProductsByCat/{catName}", h.ListProductsByCat).Methods("GET")
	http.ListenAndServe(":3000", router)
}

func (h *HttpTransport) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var user *types.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = uuid.NewString()
	//json.Unmarshal([]byte(r.Body.Close().Error()), user)
	createdUser := h.Db.CreateUser(*user)
	fmt.Println(createdUser)
	json.NewEncoder(rw).Encode(createdUser)
}

func (h *HttpTransport) ListUsers(rw http.ResponseWriter, r *http.Request) {
	users := h.Db.ListUsers()
	json.NewEncoder(rw).Encode(users)
}

func (h *HttpTransport) ShowUserCart(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userCart := h.Db.ShowUserCart(userId)
	json.NewEncoder(rw).Encode(userCart)
}

func (h *HttpTransport) AddToCart(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	var cartValue *types.AddCart
	json.NewDecoder(r.Body).Decode(&cartValue)
	cartValue.UserId = userId
	h.Db.AddToCart(*cartValue)
}

func (h *HttpTransport) AddOrder(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	h.Db.AddOrder(userId)
	//need to implement individual order
}

func (h *HttpTransport) AddProduct(rw http.ResponseWriter, r *http.Request) {
	var products *types.Product
	json.NewDecoder(r.Body).Decode(&products)
	products.Id = uuid.NewString()
	h.Db.AddProduct(*products)
}

func (h *HttpTransport) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	h.Db.DeleteProduct(productId)
}

func (h *HttpTransport) EditProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	var products *types.Product
	json.NewDecoder(r.Body).Decode(&products)
	products.Id = productId
	h.Db.EditProduct(*products)
}

func (h *HttpTransport) ListProducts(rw http.ResponseWriter, r *http.Request) {
	products := h.Db.ListProducts()
	json.NewEncoder(rw).Encode(products)
}

func (h *HttpTransport) AddCategory(rw http.ResponseWriter, r *http.Request) {
	catId := uuid.NewString()
	var cat *types.Category
	json.NewDecoder(r.Body).Decode(&cat)
	cat.Id = catId
	h.Db.AddCategory(*cat)
}

func (h *HttpTransport) ListProductsByCat(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	catId := vars["catName"]
	products := h.Db.ListProductsByCat(catId)
	json.NewEncoder(rw).Encode(products)
}

func (h *HttpTransport) ListCategory(rw http.ResponseWriter, r *http.Request) {
	cat := h.Db.ListCategory()
	json.NewEncoder(rw).Encode(cat)
}

func (h *HttpTransport) ListOrders(rw http.ResponseWriter, r *http.Request) {
	orders := h.Db.ListOrders()
	json.NewEncoder(rw).Encode(orders)
}
