Gopath
======

Warning: Don't alias gopath as 'go', it runs `go` internally which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself which would then call itself...

What it does
------------
Gopath gives each project its own GOPATH

Why each project needs its own GOPATH
-------------------------------------
Each project needs to keep its dependencies separate from the dependencies of other projects

How to use it
-------------
Just use the `gopath` command anytime you would use `go`.

How it does what it does
------------------------
  - Check for nearest .gopath file
    - If none found, just use $GOPATH
    - If found
      - Read file contents
        - If not empty, use path in file
        - If empty, use file location
  - Passes your command through to `go`

Gotchas
-------

 - Evaluating environment variable in command

 > Environment variables written in your command will be evaluated by the shell before `gopath` has a change to update your environment
 > 
 > Example: gopath echo $GOPATH
 > Will print your standard GOPATH, ignoring the closest .gopath file.
 > 
 > Solution: gopath printenv GOPATH
 > If all you want is the resulting GOPATH, you just just run gopath with no arguments.

Tips n Tricks
-------------

 - Persisting gopath for multiple commands without changing your environment

 > This will print the correct gopath three times without changing your environment
 > (GOPATH=`gopath`; printenv GOPATH && echo $GOPATH; printenv GOPATH)
