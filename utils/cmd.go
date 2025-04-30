package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Cmd(base string, args ...string) error {
	cmd := exec.Command(base, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s %s failed: %v\n", base, strings.Join(args, " "), err)
	}
	return err
}
