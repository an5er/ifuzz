package payloads

import "strings"

func LoadList(value string) []string {
	return strings.Split(value, "-")
}
