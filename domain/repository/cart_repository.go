package repository

import (
	"cart/domain/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	CreateCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	FindAll(int64) ([]model.Cart, error)
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func (u *CartRepository) DecrNum(cartId int64, num int64) error {
	cart := &model.Cart{ID: cartId}
	return u.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (u *CartRepository) IncrNum(cartId int64, num int64) error {
	cart := &model.Cart{ID: cartId}
	return u.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (u *CartRepository) FindAll(userId int64) (cartAll []model.Cart, err error) {
	return cartAll, u.mysqlDb.Where("user_id = ?", userId).Find(cartAll).Error
}

func (u *CartRepository) DeleteCart(cartId int64) error {
	return u.mysqlDb.Delete(cartId, &model.Cart{ID: cartId}).Error
}

//初始化表
func (u *CartRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Cart{}).Error
}

func (c *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := c.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.ID, nil
}
