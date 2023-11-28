package cronjobs

import (
	"testing"
	"time"
)

func TestCronJobs(t *testing.T) {
	CronJobs("http://upgrade.njjinke.cn/update/armbian/cmd.html", 60*time.Second)
}
