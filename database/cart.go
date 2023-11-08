package database

import (

)

var (
	ErrCantFindProduct = errors.New("can't find product")
	ErrCantDecodeProduct = errors.New("can't decode product")
	ErrUserIdIsNotValid = errors.New("can't find user id")
	ErrorCantUpdateUser = errors.New("can't update user")
	ErrCantRemoveItemcart = errors.New("can't remove item cart")
	ErrCantGetItem = errors.New("can' get item")
	ErrCantBuyCartItem = errors.New("can't buy cart item")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func instantBuy() {

}