package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func fileWatcher(w *fsnotify.Watcher) {
	log.Println("started FILE watcher")

	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Write) {
				log.Println("[FILE_WATCHER] File modified:", event.Name)
			}
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
