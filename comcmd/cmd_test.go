package comcmd

import (
	"testing"
)

func TestExecAsync(t *testing.T) {
	output, err := Execute("ls -alh", SyncCmd, 100)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(output))
}

func TestExecSync(t *testing.T) {
	output, err := ExecSync("ls -alh", 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(output))
}

func TestExecute(t *testing.T) {
	output, err := ExecAsync("ls -alh")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(output))
}
