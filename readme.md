Gopath
======

Warning: Don't run gopath as 'go', it runs `go` internally which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself...

What it does
------------
Gopath give each project its own GOPATH

Why each project needs its own GOPATH
-------------------------------------
Each project needs to keep its dependencies separate form the dependencies of other projects

How to use it
-------------
Just use the `gopath` command anytime you would use `go`.

How it does what it does
------------------------
	- Check for nearest .gopath file
		- If none found, just use $GOPATH
		- If found, set GOPATH to file location
  - Passes your command through to `go`
