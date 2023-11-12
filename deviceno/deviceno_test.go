package deviceno

import "testing"

func TestGenerateUniStringWithPrefix(t *testing.T) {
	for i := 0; i < 10; i++ {
		dno := GenerateUniStringWithPrefix("mengarm")
		t.Log(dno)
	}
	dno := GenerateUniStringWithPrefix("android")
	t.Log(dno)
}

func TestGenerateDeviceNoWithPrefix(t *testing.T) {
	GenerateDeviceNoWithPrefix("mengarm", "/tmp/test.txt")
}
