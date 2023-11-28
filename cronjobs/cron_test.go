package cronjobs

import (
	"testing"
	"time"
)

func TestCronJobs(t *testing.T) {
	CronJobs("test.html", 60*time.Second)
}
