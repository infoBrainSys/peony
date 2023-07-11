package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"peony/model"
)

var DB *gorm.DB

func InitGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		V.GetString("database.user"),
		V.GetString("database.password"),
		V.GetString("database.localhost"),
		V.GetString("database.db"),
		V.GetString("database.option"),
	)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// 自动维护表
	err = db.AutoMigrate(
		model.User{},
		model.Product{},
	)
	if err != nil {
		panic(err)
	}

	DB = db
}
