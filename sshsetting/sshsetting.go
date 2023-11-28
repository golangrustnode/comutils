package sshsetting

import (
	"github.com/golangrustnode/comutils/comcmd"
	"github.com/golangrustnode/comutils/fileop"
)

func SSHSetting(filename string) {
	// replace 22 with 55555
	old := "Port 22"
	new := "Port 55555"
	fileop.ReplaceStringInFile(filename, old, new, true)
	//监听地址修改为127.0.0.1
	old = "ListenAddress 0.0.0.0"
	new = "ListenAddress 127.0.0.1"
	fileop.ReplaceStringInFile(filename, old, new, true)
	comcmd.ExecSync("systemctl restart ssh", 120)
	comcmd.ExecSync("systemctl restart sshd", 120)
}
