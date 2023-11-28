package model

import (
	"log/slog"
	"processAVIWithXorm/storage/mysql"
	"processAVIWithXorm/util"
	"time"
)

type Save struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	SrcSize    uint64    `xorm:"comment('源文件大小 字节') Double" json:"src_size"`
	SrcName    string    `xorm:"comment('源文件名') Varchar(255)" json:"src_name"`
	DstSize    uint64    `xorm:"comment('目标文件大小 字节') Double" json:"dst_size"`
	DstName    string    `xorm:"comment('目标文件名') Varchar(255)" json:"dst_name"`
	Size       float64   `xorm:"comment('节省的空间 GB') Double" json:"save"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

// select sum(size)from Save;
func SyncSave() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(Save))
		if err != nil {
			slog.Error("同步节省空间数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步节省空间数据表成功")
		}
	}
}

func (s *Save) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(s)
	}
	return 0, nil
}
func (s *Save) InsertAll(ss []*Save) (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(&ss)
	}
	return 0, nil
}
