package fileop

import (
	"github.com/golangrustnode/log"
	"testing"
)

func TestReplaceStringInFile(t *testing.T) {
	filename := "/Users/developer/workspace/go/src/tutorial/sshd_config"
	old := "ListenAddress 0.0.0.0"
	new := "ListenAddress 127.0.0.1"
	if err := ReplaceStringInFile(filename, old, new, true); err != nil {
		log.Error(err)
	}
}
