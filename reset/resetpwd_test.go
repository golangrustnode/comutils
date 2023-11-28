package reset

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
	"time"
)

func TestResetPasswd(t *testing.T) {
	nowstr := time.Now().String()
	hash := md5.Sum([]byte(nowstr))
	pwd := hex.EncodeToString(hash[:])
	t.Log(pwd)
}
