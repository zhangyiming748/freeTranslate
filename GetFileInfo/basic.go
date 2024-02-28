package GetFileInfo

// todo 重新改写为 一个-一个文件夹-全部文件夹
import (
	"freeTranslate/GetAllFolder"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// note 重新改写为具有全部类型文件的通用结构体

type BasicInfo struct {
	FullPath  string      `json:"full_path,omitempty"`  // 文件的绝对路径
	FullName  string      `json:"full_name,omitempty"`  // 文件名
	PurgeName string      `json:"purge_name,omitempty"` // 单纯文件名
	PurgeExt  string      `json:"purge_ext,omitempty"`  // 单纯扩展名
	PurgePath string      `json:"purge_path,omitempty"` // 文件所在路径 包含最后一个路径分隔符
	MD5       string      `json:"md_5,omitempty"`       // 文件MD5
	MediaInfo interface{} `json:"media_info"`           // 文件类型对应的mediainfo结构体
}

/*
获取单个文件基础信息
*/

func GetFileInfo(absPath string) BasicInfo {
	ext := path.Ext(absPath)
	dir, file := filepath.Split(absPath)
	i := BasicInfo{
		FullPath:  absPath,
		FullName:  file,
		PurgeName: strings.Replace(file, ext, "", 1),
		PurgeExt:  strings.Replace(ext, ".", "", 1),
		PurgePath: dir,
		MD5:       GetMD5(absPath),
	}
	// todo 测试使用文件头判断文件类型
	//t := SelectType(strings.Replace(ext, ".", "", 1))

	return i
}

/*
获取目录下符合条件的所有文件基础信息
*/
func GetAllFileInfo(dir, pattern string) []BasicInfo {
	var aim []BasicInfo
	files, err := os.ReadDir(dir)
	if err != nil {
		slog.Warn("出错", slog.Any("读取文件夹下内容", err))
		return nil
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			slog.Debug("获取文件信息", slog.String("跳过隐藏文件", file.Name()))
			continue
		}
		if file.IsDir() {
			slog.Debug("获取文件信息", slog.String("跳过文件夹", file.Name()))
			continue
		}
		absPath := strings.Join([]string{dir, file.Name()}, string(os.PathSeparator))
		ext := strings.Replace(path.Ext(absPath), ".", "", 1)
		if In(ext, strings.Split(pattern, ";")) {
			bi := GetFileInfo(absPath)
			aim = append(aim, bi)
			slog.Debug("获取到的单个文件全部信息", slog.Any("", bi))
		} else {
			slog.Info("跳过非目标文件")
		}
	}
	return aim
}

/*
获取全部目录下符合条件的所有文件基础信息
*/
func GetAllFilesInfo(dir, pattern string) []BasicInfo {
	var aims []BasicInfo
	folders := GetAllFolder.List(dir)
	//包括根目录
	folders = append(folders, dir)
	for _, folder := range folders {
		aim := GetAllFileInfo(folder, pattern)
		aims = append(aims, aim...)
	}
	return aims
}
