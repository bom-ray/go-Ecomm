package testMain

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	// vara := mux.Vars(r)
// 	// id = vara["id"]
// 	fmt.Printf("got / request\n")
// 	io.WriteString(w, "This is my website!\n")
// }
// func getHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got /hello request\n")
// 	io.WriteString(w, "Hello, HTTP!\n")
// }

// func testMain() {
// 	// id := uuid.NewString()
// 	//fmt.Println(id)
// 	cart := make(map[string]types.Cart)
// 	usersDb := make(map[string]types.User)
// 	orders := make([]types.Orders, 0)
// 	products := make(map[string]types.Product)
// 	category := make([]types.Category, 0)
// 	catToProds := make(map[string][]types.Product)
// 	store := store.NewStorage(usersDb, cart, orders, products, category, catToProds)
// 	logic := service.NewLogic(store)
// 	//fmt.Println(logic.Db)

// 	product1 := types.Product{
// 		Id:          uuid.NewString(),
// 		Name:        "product1",
// 		Category:    []string{"Electronics", "computer"},
// 		Size:        []string{"XL"},
// 		Img:         []string{"https://bookinshop.com"},
// 		Rating:      int64(3),
// 		Price:       30.0,
// 		InStock:     true,
// 		Description: "first Product",
// 		Videos:      []string{"https://video.com"},
// 	}

// 	logic.Db.AddProduct(product1)

// 	user1 := types.User{
// 		Id:        uuid.NewString(),
// 		FirstName: "james",
// 		LastName:  "mathew",
// 		Email:     "bomray@yahoo.com",
// 		Image:     "https://image",
// 	}

// 	user2 := types.User{
// 		Id:        uuid.NewString(),
// 		FirstName: "andrew",
// 		LastName:  "brad",
// 		Email:     "andrew@yahoo.com",
// 		Image:     "https://image",
// 	}

// 	logic.CreateUser(user1)
// 	logic.CreateUser(user2)
// 	//user_cart := logic.Db.ShowUserCart("2")

// 	add_cart1 := types.AddCart{
// 		UserId:    user1.Id,
// 		ProductId: product1.Id,
// 		Qty:       5,
// 		Amount:    30.0,
// 	}
// 	add_cart2 := types.AddCart{
// 		UserId:    user2.Id,
// 		ProductId: product1.Id,
// 		Qty:       5,
// 		Amount:    30.0,
// 	}
// 	logic.AddToCart(add_cart1)
// 	logic.AddToCart(add_cart2)

// 	//add order
// 	logic.Db.AddOrder(user1.Id)
// 	//userList := logic.ListUsers()
// 	//fmt.Println(userList)
// 	//fmt.Println(logic.Db.ShowUserCart("2"))
// 	prod := logic.Db.ListProducts()

// 	fmt.Println(prod)
// 	phone_cat := types.Category{
// 		Id:   uuid.NewString(),
// 		Name: "phone",
// 	}
// 	logic.Db.AddCategory(phone_cat)
// 	//fmt.Println(logic.Db.ListCategory())

// 	computer_cat := types.Category{
// 		Id:   uuid.NewString(),
// 		Name: "computer",
// 	}
// 	logic.Db.AddCategory(computer_cat)
// 	//fmt.Println(logic.Db.ListCategory())

// 	fmt.Println(logic.Db.ListProductsByCat("phone"))

// 	mux := mux.NewRouter()
// 	mux.HandleFunc("/", getRoot).Methods("POST")
// 	mux.HandleFunc("/hello", getHello)

// 	http.ListenAndServe(":3333", mux)

// 	// mux := http.NewServeMux()
// 	// mux.HandleFunc("GET /path/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Fprint(w, "got path\n")

// 	// })

// 	// mux.HandleFunc("GET /task/{id}/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	//id := r.PathValue("id")
// 	// 	id := r.URL.Path
// 	// 	fmt.Fprintf(w, "handling task with id=%v\n", id)
// 	// })

// 	// http.ListenAndServe("localhost:8090", mux)

// }
