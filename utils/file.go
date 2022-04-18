package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Read(path string) []string {
	var result []string
	FileIsExist(path)
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	} else {
		contents := string(lines)
		lines := strings.Split(contents, "\r\n")
		lines = strings.Split(contents, "\n")
		result = make([]string, 0, len(lines))
		for _, line := range lines {
			result = append(result, line)
		}
	}
	return result
}

func FileIsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息

	if err != nil {

		if os.IsExist(err) {

			return true

		}

		return false

	}

	return true
}
