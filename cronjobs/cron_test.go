package cronjobs

import "testing"

func TestCronJobs(t *testing.T) {
	CronJobs("http://upgrade.njjinke.cn/update/armbian/cmd.html")
}
