package model

import (
	"log/slog"
	"processAVIWithXorm/storage/mysql"
	"processAVIWithXorm/util"
	"time"
)

type YTdlp struct {
	Id         int64     `xorm:"not null pk autoincr comment('主键id') INT(11)" json:"id"`
	TaskId     int64     `xorm:"not null  comment('任务id') INT" json:"TaskId"`
	URL        string    `xorm:"comment('预计文件帧数') VARCHAR(255)" json:"frame"`
	UpdateTime time.Time `xorm:"updated comment('更新时间') DateTime" json:"update_time"`
	CreateTime time.Time `xorm:"created comment('创建时间') DateTime" json:"create_time"`
	DeleteTime time.Time `xorm:"deleted comment('删除时间') DateTime" json:"delete_time"`
}

func SyncYTdlp() {
	if util.GetVal("mysql", "switch") == "on" {
		err := mysql.GetSession().Sync2(new(YTdlp))
		if err != nil {
			slog.Error("同步数据表出错", slog.Any("错误原文", err))
			return
		} else {
			slog.Debug("同步数据表成功")
		}
	}
}

func (y *YTdlp) GetAll() (ytdlps []YTdlp, err error) {
	if util.GetVal("mysql", "switch") == "on" {
		err = mysql.GetSession().Find(ytdlps)
		if err != nil {
			return nil, err
		}
		return ytdlps, nil
	}
	return nil, nil
}
func (y *YTdlp) InsertOne() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Insert(y)
	}
	return 0, nil
}
func (y *YTdlp) FindById() (bool, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(y.Id).Get(y)
	}
	return false, nil
}
func (y *YTdlp) UpdateById() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().ID(y.Id).Update(y)
	}
	return 0, nil
}
func (y *YTdlp) UpdateByTaskId() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Where("task_id = ?", y.TaskId).Update(y)
	}
	return 0, nil
}
func (y *YTdlp) Sum() (int64, error) {
	if util.GetVal("mysql", "switch") == "on" {
		return mysql.GetSession().Count(y)
	}
	return 0, nil
}
