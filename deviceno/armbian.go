package deviceno

import (
	"github.com/golangrustnode/log"
	"github.com/skip2/go-qrcode"
)

func ArmbianDeviceNoForImg(prefix, uidPath, qrPath string) error {
	if !IsExist(uidPath) {
		//生成uuid
		deviceno := GenerateUniStringWithPrefix(prefix)
		//写入文件
		Write2file(deviceno, uidPath)
	}
	//生成二维码
	dco, err := ReadFile(uidPath)
	if err != nil {
		log.Error(err)
		return err
	}
	qrcontent, err := GenerateQR(dco)
	if err != nil {
		return err
	}
	return OverrideWrite2File(qrcontent, qrPath)
}

func GenerateQR(deviceno string) (string, error) {
	//把deviceno生成二维码
	//1.生成二维码
	//2.把二维码写入文件
	qrcode, err := qrcode.New(deviceno, qrcode.Medium)
	if err != nil {
		log.Error(err)
		return "", err
	}
	qrstr := qrcode.ToString(true)
	return qrstr, nil
}
