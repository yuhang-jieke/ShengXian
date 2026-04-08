package model

import (
	"ShengXian/handler/queson"
	"ShengXian/handler/request"

	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	Name   string  `gorm:"type:varchar(30);comment:商品名称"`
	Price  float64 `gorm:"type:decimal(10,2);comment:商品价格"`
	Stock  int     `gorm:"type:int(11);comment:商品库存"`
	Status uint    `gorm:"type:int(1);comment:商品状态"`
}

func (g *Goods) FindGoodsCount(db *gorm.DB, name string) error {
	return db.Where("name=?", name).First(&g).Error
}

func (g *Goods) GetGoods(db *gorm.DB, form request.GetGoods) ([]queson.GetGoods, error) {
	var list []queson.GetGoods
	tx := db.Model(&Goods{})
	if form.Page <= 0 || form.Page > 3 {
		form.Page = 1
	}
	if form.Size <= 0 || form.Size > 3 {
		form.Size = 1
	}
	if form.Name != "" {
		tx = tx.Where("name=?", form.Name)
	}
	offset := (form.Page - 1) * form.Size
	err := tx.Offset(offset).Limit(form.Size).Find(&list).Error
	return list, err
}

func (g *Goods) ItemGoods(db *gorm.DB, form request.Item) error {
	return db.Where("id=?", form.Id).First(&g).Error
}

func (g *Goods) CountJian(db *gorm.DB, form request.GoodsCar) error {
	return db.Model(&Goods{}).Where("name=?", form.Name).Update("stock", gorm.Expr("stock-?", form.Num)).Error
}

type Shop struct {
	gorm.Model
	Name    string `gorm:"type:varchar(30);comment:店铺名称"`
	Address string `gorm:"type:varchar(30);comment:店铺地址"`
}
type Types struct {
	gorm.Model
	Name string `gorm:"type:varchar(30);comment:分类名称"`
}
type ShopCar struct {
	gorm.Model
	GoodsName string  `gorm:"type:varchar(30);comment:商品名称"`
	Price     float64 `gorm:"type:decimal(10,2);comment:单价金额"`
	Num       int     `gorm:"type:int(11);comment:数量"`
	Total     float64 `gorm:"type:decimal(10,2);comment:总金额"`
}

func (c *ShopCar) AddShopCar(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c *ShopCar) ShopList(db *gorm.DB, id int) error {
	return db.Where("id=?", id).Find(&c).Error
}

type Address struct {
	gorm.Model
	Name string `gorm:"type:varchar(30);comment:地址名称"`
}

func (a *Address) CreateAdd(db *gorm.DB) error {
	return db.Create(&a).Error
}

type Order struct {
	gorm.Model
	OrderSn string  `gorm:"type:varchar(32);comment:订单号"`
	Price   float64 `gorm:"type:decimal(10,2);comment:总金额"`
	Status  string  `gorm:"type:varchar(30);comment:状态"`
}
