# Walk through a git repo with ease.

gwalk is a utility written in go that can navigate through the evolution of git repo. 

Start with the inital commit in master, and move forward one by one to learn all about the code you are reading.

## Usage
```
$ gwalk -b master
$ gwalk --help :
  -b string
        branch to use (default "master")

Available actions - init,next,prev,exit
```

## TODO
- Improved error handling
- Better logging to stdout
- Support for other versioning systems