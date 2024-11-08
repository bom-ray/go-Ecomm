package types

// type CreateUser struct {
// 	FirstName string
// 	LastName  string
// 	Email     string
// 	Image     string
// }

type Category struct {
	Id   string
	Name string
}

type Review struct {
	Id        string
	Comment   string
	SenderId  string
	SenderUrl string
	Rating    int64
}

type Product struct {
	Id          string
	Name        string
	Category    []string
	Size        []string
	Img         []string
	Rating      int64
	Review      []Review
	InStock     bool
	Price       float64
	Description string
	Brand       string
	Videos      []string
	Sold        bool
}

type Orders struct {
	UserId     string
	OrdersList []CartItems
	Total      float64
}

type AddCart struct {
	UserId    string
	ProductId string
	Qty       int64
	Amount    float64
}

type CartItems struct {
	ProductId string
	Qty       int64
	Amount    float64
}

type Cart struct {
	UserId string
	Items  map[string]CartItems
	Total  float64
}

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Image     string
}
