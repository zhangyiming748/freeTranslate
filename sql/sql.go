package sql

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

var db *gorm.DB

func SetEngine() {
	db, _ = gorm.Open(sqlite.Open("trans.db"), &gorm.Config{})
	// 迁移 schema
	err := db.AutoMigrate(History{})
	err = db.AutoMigrate(Sensitive{})
	if err != nil {
		slog.Error("数据库初始化失败", slog.Any("错误原文", err))
		os.Exit(-1)
	}
	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// Delete - 删除 product
	//db.Delete(&product, 1)
	fmt.Println(db)
}
func GetEngine() *gorm.DB {
	return db
}
