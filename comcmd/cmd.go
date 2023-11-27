package funcs

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/golangrustnode/log"
)

//type Exec struct{}

var defaultTimeout int64 = 10

const (
	SyncCmd  int64 = 1
	AsyncCmd int64 = 2
)

/*
执行shell 命令，支持管道
1. 1 普通命令
2. 2 执行daemon进程，长时间运行的命令一般需要使用daemon
3. 其它则默认忽略
*/
func Execute(cmds string, cmdType int64, timeOut int64) (res []byte, err error) {
	res, err = nil, nil
	defer func() {
		if err2 := recover(); err2 != nil {
			log.Error(err2)
			err = errors.New(fmt.Sprintf("recover error: %v", err2))
		}
	}()

	switch cmdType {
	case 1:
		res, err = ExecSync(cmds, timeOut)
	case 2:
		res, err = ExecAsync(cmds)
	}
	return
}

/*
命令异步执行，不需要超时时间
*/
func ExecAsync(cmds string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", cmds)
	sysProcAttr := &syscall.SysProcAttr{}
	// 设置将子进程放入新的进程组
	sysProcAttr.Setpgid = true
	cmd.SysProcAttr = sysProcAttr
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Error("async command wait:", err)
		}
	}()
	return []byte("async executed successfully"), nil
}
func ExecSync(cmds string, timeoutSec int64) ([]byte, error) {
	cmd := exec.Command("bash", "-c", cmds)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	done := make(chan error)
	go func() { done <- cmd.Wait() }()
	timeout := time.After(time.Duration(timeoutSec) * time.Second)
	select {
	case <-timeout:
		cerr := cmd.Process.Kill()
		killError := fmt.Sprintf("timeout and kill error:%v", cerr)
		return nil, errors.New(killError)
	case err := <-done:
		return buf.Bytes(), err
	}
}
