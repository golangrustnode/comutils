package reset

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/golangrustnode/log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ResetPasswd() {
	resetPasswdFile := "/etc/resetPasswd2.txt"
	if fileExists(resetPasswdFile) {
		log.Info("密码已经设置，无需修改")
		return
	}
	nowstr := time.Now().String()
	hash := md5.Sum([]byte(nowstr))
	accountPassword := hex.EncodeToString(hash[:])
	cmd := exec.Command("passwd", "root")
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%s\n%s\n", accountPassword, accountPassword))
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(err)
		return
	}
	file, err := os.Create(resetPasswdFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	log.Info("修改密码成功")
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
