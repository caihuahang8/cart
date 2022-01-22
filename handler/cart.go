package handler

import (
	"cart/domain/model"
	"cart/domain/service"
	cart "cart/proto"
	"context"
	common "github.com/caihuahang8/common"
)

type Cart struct {
	CartDataService service.ICartDateService
}

//添加购物车
func (h *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request, cart)
	response.CartId, err = h.CartDataService.AddCartDate(cart)
	return err
}

//清空购物车
func (h *Cart) CleanCart(ctx context.Context, request *cart.Clean, response *cart.Response) error {
	//implement method
	return nil
}

//添加购物车数量
func (h *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := h.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "添加购物车成功"
	return nil
}

//购物车减少商品数量
func (h *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := h.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "减少购物车成功"
	return nil
}

//删除购物车
func (h *Cart) DeleteItemByID(ctx context.Context, request *cart.CartID, response *cart.Response) error {
	err := h.CartDataService.DeleteCartData(request.Id)
	if err != nil {
		return err
	}
	response.Meg = "删除成功"
	return nil
}

//查询用户所有的购物车信息
func (h *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error {
	//1.根据用户id查找出所有购物车
	cartAll, err := h.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}
	//2.映射到CartInfo集合返回给前端
	for _, v := range cartAll {
		cartInfo := &cart.CartInfo{}
		if err := common.SwapTo(v, cartInfo); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cartInfo)
	}
	return nil
}
