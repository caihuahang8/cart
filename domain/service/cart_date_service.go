package service

import (
	"github.com/caihuahang8/cart/domain/model"
	"github.com/caihuahang8/cart/domain/repository"
)

type ICartDateService interface {
	AddCartDate(*model.Cart) (int64, error)
	DeleteCartData(int64) error
	FindAllCart(int64) ([]model.Cart, error)
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

type CartDateService struct {
	CartRepository repository.ICartRepository
}

//创建
func NewCartDataService(cartRepository repository.ICartRepository) ICartDateService {
	return &CartDateService{cartRepository}
}

func (c *CartDateService) AddCartDate(cart *model.Cart) (cartId int64, err error) {
	return c.CartRepository.CreateCart(cart)
}

func (c *CartDateService) DeleteCartData(cartId int64) error {
	return c.CartRepository.DeleteCart(cartId)
}

func (c *CartDateService) FindAllCart(userId int64) ([]model.Cart, error) {
	return c.CartRepository.FindAll(userId)
}

func (c *CartDateService) IncrNum(cartId int64, num int64) error {
	return c.CartRepository.IncrNum(cartId, num)
}

func (c *CartDateService) DecrNum(cartId int64, num int64) error {
	return c.CartRepository.DecrNum(cartId, num)
}
