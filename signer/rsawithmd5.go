package signer

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/golangrustnode/log"
)

/*
参考文档https://www.kancloud.cn/tonyyu/go_work_cases/1797471
*/
func Sign(data, pvstr string) (string, error) {
	privbytes, err := base64.StdEncoding.DecodeString(pvstr)
	if err != nil {
		log.Error(err)
		return "", err
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(privbytes)
	if err != nil {
		log.Error(err)
		return "", err
	}
	h := md5.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.MD5, hash[:])
	if err != nil {
		fmt.Println("SignPKCS1v15:", err)
		return "", err
	}
	output := base64.StdEncoding.EncodeToString(signature)
	return output, nil
}

func Verify(originalData, signData, pubKey string) {}
