package cmds

import (
	"fmt"
	"os"
	"strings"
)

func Fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func merge(cn string, sans []string) []string {
	var found bool
	for _, name := range sans {
		if strings.EqualFold(name, cn) {
			found = true
			break
		}
	}
	if found {
		return sans
	}
	return append([]string{cn}, sans...)
}
