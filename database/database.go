package database

import (
	"gorm.io/gorm"
	_ "gorm.io/driver/postgres"
)

var (
	DBConn *gorm.DB
)
