package sql

import (
	"fmt"
	"freeTranslate/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type History struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	From      string `gorm:"form"`
	To        string `gorm:"to"`
	Src       string `gorm:"src"`
	Dst       string `gorm:"dst"`
	Source    string `gorm:"source"`
	ErrorMsg  string `gorm:"error_msg"`
	ErrorCode string `gorm:"error_code"`
	Request   string `gorm:"request"`
	CreatedAt time.Time
}

func SetEngine() {
	db, _ = gorm.Open(sqlite.Open("trans.db"), &gorm.Config{})

	// 迁移 schema
	err := db.AutoMigrate(model.History{})
	if err != nil {
		return
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
func (h *History) FindOneBySrc(string) *gorm.DB {
	return GetEngine().First(&h, "src = ?", h.Src)

}
func (h *History) SetOne() *gorm.DB {
	return GetEngine().Create(&h)
}
