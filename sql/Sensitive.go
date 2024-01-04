package sql

import (
	"gorm.io/gorm"
	"time"
)

type Sensitive struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Src       string `gorm:"sensitive"`
	Dst       string `gorm:"safety"`
	CreatedAt time.Time
}

func (s *Sensitive) FindBySrc() *gorm.DB {
	return GetEngine().Where("src = ?", s.Src).First(&s)
}
func (s *Sensitive) GetAll() []Sensitive {
	all := []Sensitive{}
	GetEngine().Table("sensitives").Find(&all)
	return all
}
func (s *Sensitive) AddOne() *gorm.DB {
	return GetEngine().Create(&s)
}
