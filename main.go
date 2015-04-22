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
	cmd := exec.Command("go", os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
