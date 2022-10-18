package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func dirWatcher(w *fsnotify.Watcher) {
	log.Println("started DIR watcher")

	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Create) {
				log.Println("[DIR_WATCHER] File created: ", event.Name)
			} else if event.Has(fsnotify.Remove) {
				log.Println("[DIR_WATCHER] File deleted: ", event.Name)
			}
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
