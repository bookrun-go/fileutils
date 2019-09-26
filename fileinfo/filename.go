package fileinfo

import (
	"path"
	"path/filepath"
	"strings"
)

func GetFileSimpleName(fullFileName string) string {
	filenameWithSuffix := filepath.Base(fullFileName)
	fileSuffix := path.Ext(filenameWithSuffix)                         //获取文件后缀
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名
	return filenameOnly
}
