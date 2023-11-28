package sshsetting

import (
	"github.com/golangrustnode/comutils/fileop"
)

func SSHSetting() {
	// replace 22 with 55555
	filename := "/etc/ssh/sshd_config"
	old := "Port 22"
	new := "Port 55555"
	fileop.ReplaceStringInFile(filename, old, new, true)
	//监听地址修改为127.0.0.1
	old = "ListenAddress 0.0.0.0"
	new = "ListenAddress 127.0.0.1"
	fileop.ReplaceStringInFile(filename, old, new, true)
}
