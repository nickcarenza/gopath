Gopath
======

What it does
------------
Gopath give each project its own GOPATH

Why each project needs its own GOPATH
-------------------------------------
Each project needs to keep its dependencies separate form the dependencies of other projects

How it does what it does
------------------------
	- Check for nearest .gopath file
		- If none found, just use $GOPATH
		- If found, set GOPATH to file location
