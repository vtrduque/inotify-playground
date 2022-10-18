package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

const (
	watchFile = "./tmp/file.log"
	watchDir  = "./tmp"
)

func main() {
	watcherFile, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	watcherDir, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcherFile.Close()
	defer watcherDir.Close()

	go fileWatcher(watcherFile)
	go dirWatcher(watcherDir)

	err = watcherFile.Add(watchFile)
	if err != nil {
		log.Fatal(err)
	}

	err = watcherDir.Add(watchDir)
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}
