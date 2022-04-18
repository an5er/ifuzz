package payloads

import (
	"bufio"
	"os"
)

func LoadStdin() []string {
	payloads := make([]string, 0)
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		payloads = append(payloads, scan.Text())
	}

	return payloads
}
