package payloads

import "ifuzz/utils"

func LoadFile(value string) []string {
	payloads := make([]string, 0)
	if utils.FileIsExist(value) {
		payloads = utils.SliceAdd(payloads, utils.Read(value))
	}
	return payloads
}
