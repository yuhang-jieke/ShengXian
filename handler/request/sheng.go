package request

type Item struct {
	Id string `form:"id"   binding:"required"`
}
type GoodsCount struct {
	Name string `form:"name"   binding:"required"`
}
type GetGoods struct {
	Name string `form:"name"   binding:""`
	Page int    `form:"page"   binding:""`
	Size int    `form:"size"   binding:""`
}
type GoodsCar struct {
	Name string `form:"name"   binding:""`
	Num  int    `form:"num"   binding:""`
}
type ShopCarList struct {
	Id int `form:"id"   binding:""`
}
type Address struct {
	Name string `form:"name"   binding:""`
}
