package deviceno

import (
	"fmt"
	"testing"
)

func TestGetNetInfo(t *testing.T) {
	nInfo := GetNetInfo()

	t.Log(nInfo)
	s1 := "字符串"
	s2 := "拼接"
	s3 := s1 + s2
	fmt.Print(s3)
}
