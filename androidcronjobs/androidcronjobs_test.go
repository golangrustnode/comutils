package androidcronjobs

import (
	"testing"
	"time"
)

func TestAndroidCronJobs(t *testing.T) {
	AndroidCronJobs("http://upgrade.njjinke.cn/test/androidtest.sh", 1*time.Second, 6*time.Second)
}
