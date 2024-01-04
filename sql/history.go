package sql

import (
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

func (h *History) FindOneBySrc() *gorm.DB {
	return GetEngine().First(&h, "src = ?", h.Src)
}
func (h *History) SetOne() *gorm.DB {
	return GetEngine().Create(&h)
}
func (h *History) ExistSrc() bool {
	result := GetEngine().First(&h, "src = ?", h.Src)
	if result.Error.Error() == "record not found" {
		return false
	} else {
		return true
	}
}
