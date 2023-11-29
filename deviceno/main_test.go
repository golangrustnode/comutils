package deviceno

import "testing"

func TestCustomDeviceNo(t *testing.T) {
	CustomDeviceNo("justatest", "/tmp/test.txt", "0.0.0.0:34568")
}
