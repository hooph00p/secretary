# The Secretary

- [ ] Repetitive Command Line Operations
  - [x] Command Line
    - [x] String
  - [x] Repeat
  - [x] Interval
    - [x] Seconds
    - [ ] Minutes
    - [ ] Hours
    - [ ] Days

  - [ ] Scheduling Command Line Operations
    - [ ] Start
    - [ ] End

  - [ ] Web GUI

Command:
```
$ secretary -command="go version" -interval=10 -repeat=5
```

Output:
```
[ Secretary ] Command: go version
[ Secretary ] Interval: 10
[ Secretary ] Repeat: 5

[ Secretary ] Iteration: 1 of 5
[ Secretary ] Time: 2016-10-28 18:00:49.114672686 -0400 EDT
go version go1.6 darwin/amd64

[ Secretary ] Iteration: 2 of 5
[ Secretary ] Time: 2016-10-28 18:00:59.114868178 -0400 EDT
go version go1.6 darwin/amd64

[ Secretary ] Iteration: 3 of 5
[ Secretary ] Time: 2016-10-28 18:01:09.116759101 -0400 EDT
go version go1.6 darwin/amd64

[ Secretary ] Iteration: 4 of 5
[ Secretary ] Time: 2016-10-28 18:01:19.119451894 -0400 EDT
go version go1.6 darwin/amd64

[ Secretary ] Iteration: 5 of 5
[ Secretary ] Time: 2016-10-28 18:01:29.114552479 -0400 EDT
go version go1.6 darwin/amd64

[ Secretary ] Completed task. Ending.
```

## Display
