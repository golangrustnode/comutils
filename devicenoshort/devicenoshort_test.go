package devicenoshort

import (
	"testing"
)

func TestGenerateShortDeviceNoWithPrefix(t *testing.T) {
	deviceNO := GenerateShortDeviceNoWithPrefix("android10", 10)
	t.Log(deviceNO)
}

func TestDeviceNoShort(t *testing.T) {
	DeviceNoShort("android10ty", "/Users/developer/pcdn/tmp/pcdnuuid", "127.0.0.1:12345", 10)
}
