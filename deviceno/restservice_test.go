package deviceno

import "testing"

func TestStartHttp(t *testing.T) {
	StartHttp("/tmp/test.txt", "127.0.0.1:60001")
}
