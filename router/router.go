package router

import (
	"ShengXian/server"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("goods/find", server.GetGoodsCount)
	r.GET("get/goods", server.GetGoods)
	r.GET("item/goods", server.ItemGoods)
	r.POST("shop/add", server.GoodsCar)
	r.GET("shop/list", server.ShopCarList)
	r.POST("address", server.Address)
	return r
}
