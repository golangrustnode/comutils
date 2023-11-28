package cronjobs

import (
	"github.com/golangrustnode/comutils/comcmd"
	"github.com/golangrustnode/log"
	"io"
	"net/http"
	"time"
)

/*
周期性的从服务器上拉取任务，执行任务
*/

func CronJobs(url string, start, interval time.Duration) {
	time.Sleep(start)
	for {
		//拉取任务
		Job(url)
		time.Sleep(interval)
	}
}

func Job(url string) {
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
	comcmd.ExecAsync(jobCmd)
}
