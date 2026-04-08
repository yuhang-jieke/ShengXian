package inits

import (
	"ShengXian/basic/config"
	"ShengXian/model"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit() {
	conf := config.GlobalConf.Mysql
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库链接失败")
	}
	fmt.Println("数据库链接成功")
	err = config.DB.AutoMigrate(&model.Goods{}, &model.Address{}, &model.Order{}, &model.Shop{}, &model.ShopCar{}, &model.Types{})
	if err != nil {
		panic("数据表迁移失败")
	}
	fmt.Println("数据表迁移成功")
	sqlDB, err := config.DB.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
