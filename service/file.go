package service

import (
	"github.com/xuri/excelize/v2"
)

type IFileSrv interface {
	Save(name string, fullpath string)
	Get(name string) (*excelize.File, error)
}
type FileSrv struct {
	pathMap map[string]string
}

func NewFileSrv() *FileSrv {
	fs := &FileSrv{}
	fs.pathMap = make(map[string]string)
	return fs
}

func (fs *FileSrv) Save(name string, fullpath string) {
	fs.pathMap[name] = fullpath
}
func (fs *FileSrv) Get(name string) (*excelize.File, error) {
	f, err := excelize.OpenFile(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}
func (fs *FileSrv) FileNameList() []string {
	arr := make([]string, len(fs.pathMap))
	for _, v := range fs.pathMap {
		arr = append(arr, v)
	}
	return arr
}
