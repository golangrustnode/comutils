package cronjobs

import (
	"testing"
	"time"
)

func TestCronJobs(t *testing.T) {
	CronJobs("http://hellotest123.oss-cn-shanghai.aliyuncs.com/cmd.html", 10*time.Second, 60*time.Second)
}
