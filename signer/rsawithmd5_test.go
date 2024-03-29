package signer

import "testing"

func TestSign(t *testing.T) {
	data := "242335324中文测试{}:{}[]"
	pvstr := "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMlzuoOHEAaE/y3BZEdfq15PplvgOXp4fLMSU+bsJd6KyCSxZCdr/Tymcf8dnsQIxW33zRFm4cTWg30xTvqus0O5sCnswOvTLYRStRQ4N7Q5hlseUyTZXLesN5v8FHQhlfPxW5a8dSJYDfDSRcukXeWOoPKyVgYeQYYlWmEgF2SlAgMBAAECgYBHjbGb8idBCZBRwjVKHLMTxprCW3xPAejY+hOPHCW/S/AXCGMYlYTXPA+VnAG+g5cMnk3adn6OmdsOZqs040SMkaYIU07maUDWoZa4S7Is6O7eTAvYVLX8Z61bdztqnuaNW9QKAN8GbWesdFfZOB016aycWYcpx0jLUkXyUWajAQJBAPmdp9VWJHFgNMXtjVOGU2HnF7XouFFsCMLcGkIALylQBKHj2cNuu+2FdMNeIVwh95dxdY80JIoJ1ZS3dPayGqECQQDOmrk9f75Anqkb4tiUCC+a8Lgluar4AFTYfWkWAjaPQY4Q9upLYpEgMFokR+BNZfCSwCC2cDaKcgkr0b5PLC+FAkEAmqPiKxd/4kFA0HIXwOwR0jBRjx/KuJ6eWuHy3nqwvP5WsFixOxLqF/861CIHtsVQVyEyPHKIRIBNMiHpsz+poQJAKNmS0sQVO/Wi/i6GIA7WSs5ZuqvRUCPQotV5F67sVrrheh04Chu9Eh49VHsZsHP9cyal0SyyvAD8KzvXYxQbeQJAXEKLbGwCftNrS3zFAvAk2eFY6Up5y8u9cYHpXN1bFf12wHMMNOf+bTI7px8lksDusMmBSr1QGgYV0JWLk4XRaA=="
	sign, err := Sign(data, pvstr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(sign)
}
