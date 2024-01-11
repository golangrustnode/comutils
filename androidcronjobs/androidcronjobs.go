package androidcronjobs

import (
	"github.com/golangrustnode/comutils/androidcmd"
	"github.com/golangrustnode/log"
	"io"
	"net/http"
	"time"
)

//定时从服务端获取脚本，然后执行

func AndroidCronJobs(url string, start, interval time.Duration) {
	time.Sleep(start)
	for {
		//拉取任务
		AndroidJob(url)
		time.Sleep(interval)
	}
}

func AndroidJob(url string) {
	exec := androidcmd.Exec{}
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return
	}
	jobCmd := string(body)
	log.Info("jobCmd:", jobCmd)
	//执行任务
	//comcmd.ExecAsync(jobCmd)
	exec.ExecAsync(jobCmd)
}
