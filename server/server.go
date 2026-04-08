package server

import (
	"ShengXian/basic/config"
	"ShengXian/handler/request"
	"ShengXian/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGoodsCount(c *gin.Context) {
	var form request.GoodsCount
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	var goods model.Goods
	err := goods.FindGoodsCount(config.DB, form.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "支付成功",
		"date": "GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.[GIN-debug] [WARNING] Running in debug mode. Switch to release mode in production.using env:   export GIN_MODE=releaseusing code:  gin.SetMode(gin.ReleaseMode)",
	})
	return
}
func GetGoods(c *gin.Context) {
	var form request.GetGoods
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	var goods model.Goods
	list, err := goods.GetGoods(config.DB, form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": list,
	})
	return
}
func ItemGoods(c *gin.Context) {
	var form request.Item
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	var goods model.Goods
	err := goods.ItemGoods(config.DB, form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": model.Goods{
			Name:   goods.Name,
			Price:  goods.Price,
			Stock:  goods.Stock,
			Status: goods.Status,
		},
	})
	return
}
func GoodsCar(c *gin.Context) {
	var form request.GoodsCar
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	var goods model.Goods
	err := goods.FindGoodsCount(config.DB, form.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	err = goods.CountJian(config.DB, form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "修改失败",
		})
		return
	}
	total := float64(form.Num) * goods.Price
	if total > 30 {
		total = total - 10
	}
	shopcar := model.ShopCar{
		GoodsName: form.Name,
		Price:     goods.Price,
		Num:       form.Num,
		Total:     total,
	}
	err = shopcar.AddShopCar(config.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "购物车添加成功",
	})
	return
}
func ShopCarList(c *gin.Context) {
	var form request.ShopCarList
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	var shopcar model.ShopCar
	err := shopcar.ShopList(config.DB, form.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": model.ShopCar{
			GoodsName: shopcar.GoodsName,
			Price:     shopcar.Price,
			Num:       shopcar.Num,
			Total:     shopcar.Total,
		},
	})
	return
}
func Address(c *gin.Context) {
	var form request.Address
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	address := model.Address{
		Name: form.Name,
	}
	err := address.CreateAdd(config.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
	return
}
