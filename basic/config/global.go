package config

import "gorm.io/gorm"

var (
	GlobalConf *AppConfig
	DB         *gorm.DB
)
