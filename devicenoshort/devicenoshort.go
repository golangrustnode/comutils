package devicenoshort

import (
	"github.com/golangrustnode/comutils/deviceno"
	"math/rand"
)

func DeviceNoShort(prefix, path, listenAddr string, length int64) {
	d_no := GenerateShortDeviceNoWithPrefix(prefix, length)
	deviceno.Write2file(d_no, path)
	deviceno.StartHttp(path, listenAddr, prefix)
}

func GenerateShortDeviceNoWithPrefix(prefix string, length int64) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return prefix + string(b)
}
