package comfile

import "testing"

func TestWrite2file(t *testing.T) {
	content := "hello world"
	path := "/tmp/test.txt"
	Write2file(content, path)
}
