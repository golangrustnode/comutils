package comfile

import (
	"errors"
	"github.com/golangrustnode/log"
	"os"
)

func Write2file(content, path string) {
	//判断文件是否存在
	//如果存在就什么也不做
	//如果不存在就创建文件,并把content写入文件
	if IsExist(path) {
		log.Info("文件已经存在")
		return
	}
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	n, err := f.WriteString(content)
	log.Info("write length:", n)
	if err != nil {
		log.Fatal(err)
	}
}

func IsExist(path string) bool {
	//判断文件是否存在
	//如果存在就返回true
	//如果不存在就返回false
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
