package service

import (
	"github.com/bom-ray/goEcomm/store"
	"github.com/bom-ray/goEcomm/types"
)

type Logic struct {
	Db store.DataBase
}

func NewLogic(dB store.DataBase) *Logic {
	return &Logic{
		Db: dB,
	}
}

func (l *Logic) CreateUser(u types.User) types.User {
	user := l.Db.CreateUser(u)
	return user
}

func (l *Logic) ListUsers() []types.User {
	return l.Db.ListUsers()
}

func (l *Logic) AddToCart(addToCart types.AddCart) {
	//l.AddToCart(addToCart)
	l.Db.AddToCart(addToCart)
}
