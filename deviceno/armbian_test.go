package deviceno

import (
	"fmt"
	"testing"
)

func TestGenerateQR(t *testing.T) {
	qr, err := GenerateQR("hello world")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(qr)
}

func TestArmbianDeviceNoForImg(t *testing.T) {
	if err := ArmbianDeviceNoForImg("android", "/tmp/uidtest.txt", "/tmp/qrpathtest.txt"); err != nil {
		t.Error(err)
	}
}
