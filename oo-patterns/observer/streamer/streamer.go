package streamer

import "study/oo-patterns/observer/watcher"

type Streamer interface {
	AddWatcher(w watcher.Watcher)
	RemoveWatcher(w watcher.Watcher)
	NotifyWatchers()
}
