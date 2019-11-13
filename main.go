package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	os.Setenv("CHERE_INVOKING", "1")
	os.Setenv("MSYSTEM", "MSYS")
	cmd := exec.Command(`c:\msys64\usr\bin\bash.exe`, `--login`, `-i`)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
