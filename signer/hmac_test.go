package signer

import (
	"encoding/json"
	"testing"
)

type S struct {
	DeviceNo string `json:"deviceNo"`
}

func TestHmacWithSha256(t *testing.T) {
	secretKey := "7a4a1d992bf4e98dee11852a48215193"
	path := "/api/v2/dispatch/getDeviceInfo"
	timestamp := "12345678"
	s := &S{
		DeviceNo: "123456白日依山尽",
	}
	sd, _ := json.Marshal(s)
	body := string(sd)

	signstr := path + "\n" + timestamp + "\n" + body + "\n"
	t.Log(signstr)
	signa := HmacWithSha256(signstr, secretKey)
	t.Log(signa)
}
