package model

import (
	"freeTranslate/sql"
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

/*
Id         int64     `xorm:"not null pk autoincr comment('主键id') INT" json:"id"`

	From       string    `xorm:"not null  comment('源语言') VARCHAR(255)" json:"from"`
	To         string    `xorm:"comment('目标语言') VARCHAR(255)" json:"to"`
	Src        string    `xorm:"comment('源') VARCHAR(255)" json:"src"`
	Dst        string    `xorm:"comment('目标') VARCHAR(255)" json:"dst"`
	Source     string    `xorm:"comment('来源') VARCHAR(255)" json:"source"`
	ErrorCode  string    `xorm:"comment('错误代码') VARCHAR(255)" json:"error_code"`
	ErrorMsg   string    `xorm:"comment('错误信息') VARCHAR(255)" json:"error_msg"`
	Request    string    `xorm:"comment('请求原文') VARCHAR(255)" json:"request"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
*/
func (h *History) FindOneBySrc() *gorm.DB {
	return sql.GetEngine().First(&h, "src = ?", h.Src)

}
func (h *History) SetOne() *gorm.DB {
	return sql.GetEngine().Create(&h)
}
