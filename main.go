package main

import (
	"os"
	"os/exec"
	"path"
)

// Gotcha 1: Evaluating environment variable in command
// Environment variables written in your command will be evaluated by the shell before `gopath` has a change to update your environment
// Example: gopath echo $GOPATH
// Will print your standard GOPATH, ignoring the closest .gopath file
// Solution: gopath printenv GOPATH
// If all you want is the resulting GOPATH, you just just run gopath with no arguments

func main() {
	cwd, _ := os.Getwd()
	dir := cwd
	gopath := os.Getenv("GOPATH")
	for {
		if _, err := os.Stat(path.Join(dir, ".gopath")); err == nil {
			gopath = dir
			break
		}
		if dir == "/" {
			break
		}
		dir = path.Dir(dir)
	}
	os.Setenv("GOPATH", gopath)
	if len(os.Args) < 2 {
		os.Stdout.WriteString(gopath + "\n")
		os.Exit(0)
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
