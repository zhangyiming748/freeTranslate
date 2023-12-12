package model

import (
	"freeTranslate/storage/mysql"
	"freeTranslate/util"
	"log/slog"
	"time"
)

type History struct {
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
}

func SyncHistory() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(History))
		if err != nil {
			slog.Error("同步历史数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步历史数据表成功")
		}
	}
}

func (h *History) GetAll() (Historys []History, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(Historys)
		if err != nil {
			return nil, err
		}
		return Historys, nil
	}
	return nil, nil
}

func (h *History) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(h)
	}
	return 0, nil
}

func (h *History) FindById() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(h.Id).Get(h)
	}
	return false, nil
}
func (h *History) FindBySrc() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Where("src = ?", h.Src).Get(h)
	}
	return false, nil
}

func (h *History) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(h.Id).Update(h)
	}
	return 0, nil
}
func (h *History) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(h)
	}
	return 0, nil
}
