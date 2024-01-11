package androidcmd

import (
	"testing"
)

func TestExec_Execute(t *testing.T) {
	e := Exec{}
	// test sync
	/*

		res, err := e.Execute("ls -alh", 1)
		if err != nil {
			t.Log("------", string(res))
			t.Fatal(err.Error())
		}
		t.Log("======", string(res))
	*/
	// test async

	res, err := e.Execute("nohup /root/workspace/totmp/fuckyou > f.log 2>&1 &", AsyncCmd, defaultTimeout)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))

	// test recover
	/*
		res, err := e.Execute("service nginx stop", SyncCmd, 10)
		if err != nil {
			t.Log("------", string(res))
			t.Fatal(err.Error())
		}
		t.Log("======", string(res))
	*/
}

func TestExec_ExecSync(t *testing.T) {
	e := Exec{}
	e.ExecSync("pkill -9 fuckyou", 10)
}
