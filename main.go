package main

import (
	"os"
	"os/exec"
	"runtime"
)

func getNixShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}
	return shell
}

func getWinShell() string {
	shell := os.Getenv("ComSpec")
	if shell == "" {
		shell = "cmd"
	}
	return shell
}

func getShell() string {
	switch runtime.GOOS {
	case "linux":
		return getNixShell()
	case "openbsd":
		return getNixShell()
	case "freebsd":
		return getNixShell()
	case "darwin":
		return getNixShell()
	case "windows":
		return getWinShell()
	}

	panic(("Undefined GOOS: " + runtime.GOOS))
}

func main() {
	shell := getShell()
	command := exec.Command(shell)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		panic(err)
	}
}
