package fileop

import (
	"github.com/golangrustnode/log"
	"io/ioutil"
	"strings"
)

/*
如果文件中包含new，则不替换，否则替换old为new
*/
func ReplaceStringInFile(filename, old, new string, addedto bool) (err error) {
	notExist := true
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(err)
		return err
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		tmpLine := strings.TrimSpace(line)
		if strings.Contains(line, new) {
			log.Info("new " + new + " already exist in " + filename)
			return nil
		}
		if strings.Contains(tmpLine, old) && tmpLine[0] != '#' {
			log.Info("replace " + old + " with " + new)
			lines[i] = new
			notExist = false
		}
	}
	if notExist && addedto {
		log.Info("add " + new + " to " + filename)
		lines = append(lines, new)
	}
	output := strings.Join(lines, "\n")
	log.Info(output)
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Error(err)
	}
	return err
}
