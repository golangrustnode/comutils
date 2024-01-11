package signer

import (
	hmac2 "crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func HmacWithSha256(data, secret string) string {
	hmac := hmac2.New(sha256.New, []byte(secret))
	hmac.Write([]byte(data))
	dataHmac := hmac.Sum(nil)
	return base64.StdEncoding.EncodeToString(dataHmac)
}
