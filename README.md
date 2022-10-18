# inotify study playground

Implementation list:

- [x] Watch a file
- [x] Watch a directory
- [x] Demo
- [ ] Implement via [syscall](https://pkg.go.dev/syscall) (Ref.: [inotify](https://man7.org/linux/man-pages/man7/inotify.7.html)) :D

# Run it
```sh
$ make run
```

on another terminal
```sh
$ make run-demo
```

the output should be something like
```
2022/10/18 08:46:12 started DIR watcher
2022/10/18 08:46:12 started FILE watcher
2022/10/18 08:46:17 [FILE_WATCHER] File modified: tmp/file.log
2022/10/18 08:46:17 [DIR_WATCHER] File created:  tmp/file2.log
2022/10/18 08:46:24 [DIR_WATCHER] File deleted:  tmp/file2.log
2022/10/18 08:46:24 [DIR_WATCHER] File deleted:  tmp/file.log
2022/10/18 08:46:24 [DIR_WATCHER] File deleted:  tmp
```

# Clean
```sh
$ make clean
```
