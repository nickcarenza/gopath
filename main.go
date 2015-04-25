package main

import (
	"os"
	"os/exec"
	"path"
)

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
