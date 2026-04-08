package queson

type Login struct {
	Stock int `form:"stock"   binding:"required"`
}
type GetGoods struct {
	Id     int
	Name   string
	Price  float64
	Stock  int
	Status int
}
