package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func handleCreateEvent(event fsnotify.Event) error {
	fmt.Println("Create file", event)
	fi, err := os.Stat(event.Name)
	if err != nil {
		return err
	}
	fmt.Println(fi.Name())
	return nil
}

func handleDeleteEvent(event fsnotify.Event) error {
	fmt.Println("Delete file", event)
	_, name := filepath.Split(event.Name)
	fmt.Println(name)
	return nil
}

func StartWatch(stopCh <-chan struct{}) error {
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	// 监听两个文件夹
	if err = fsWatcher.Add("E:\\study\\golang\\go-examples"); err != nil {
		return err
	}
	if err = fsWatcher.Add("E:\\study\\golang\\go-examples\\third-party-pkg\\fsnotify-example"); err != nil {
		return err
	}

	go func(fw *fsnotify.Watcher) {
		for {
			select {
			case event := <-fw.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					err = handleCreateEvent(event)
					if err != nil {
						fmt.Println("create event failed", err)
					}
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					err = handleDeleteEvent(event)
					if err != nil {
						fmt.Println("delete event failed", err)
					}
				}
			case err = <-fsWatcher.Errors:
				fmt.Println("error:", err)
			case <-stopCh:
				fsWatcher.Close()
			}
		}
	}(fsWatcher)

	return nil
}

func main() {
	stopCh := make(chan struct{})
	if err := StartWatch(stopCh); err != nil {
		log.Fatal(err)
	}

	select {}
}
