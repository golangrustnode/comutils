package deviceno

import "testing"

func TestWrite2file(t *testing.T) {
	content := "hello world"
	path := "/tmp/test.txt"
	Write2file(content, path)
}

func TestReadFile(t *testing.T) {
	if content, err := ReadFile("/tmp/test.txt"); err != nil {
		t.Error(err)
	} else {
		t.Log(content)
	}
}

func TestOverrideWrite2File(t *testing.T) {
	content := "hello world2"
	path := "/tmp/ovtest.txt"
	OverrideWrite2File(content, path)
}
