package lib

import (
	"path/filepath"
)

type FileWatcher struct {
	watcher
	Files []string
}

func (w *FileWatcher) GetAllFiles() int {
	size := 0
	for _, name := range w.Files {
		size++
		path := filepath.Join(w.Root, w.Dir, name)
		w.processFile(path, 0)
	}
	return size
}

func (w *FileWatcher) matchesFile(name string) bool {
	for _, f := range w.Files {
		if name == filepath.Join(w.Root, w.Dir, f) {
			return true
		}
	}
	return false
}

func (w *FileWatcher) addWatchDirs() {
	for _, f := range w.Files {
		name := filepath.Join(w.Root, w.Dir, f)

		Plog.PrintC("watching", name)
		w.addWatchDir(filepath.Dir(name))
	}
}

func (w *FileWatcher) handleNewDir(name string) {
	return
}
