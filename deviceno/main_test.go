package deviceno

import "testing"

func TestCustomDeviceNo(t *testing.T) {
	CustomDeviceNo("justatest", "/tmp/test.txt", "127.0.0.1:34567")
}
