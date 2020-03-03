package streamer

import (
	"study/oo-patterns/observer/models"
	"study/oo-patterns/observer/watcher"
)

type WeatherStreamer struct {
	data               models.WeatherData
	registeredWatchers []watcher.Watcher
}

func NewWeatherStreamer() *WeatherStreamer {
	return &WeatherStreamer{
		registeredWatchers: []watcher.Watcher{},
	}
}

func (ws *WeatherStreamer) AddWatcher(w watcher.Watcher) {
	ws.registeredWatchers = append(ws.registeredWatchers, w)
}

func (ws *WeatherStreamer) RemoveWatcher(w watcher.Watcher) {
	for i := 0; i < len(ws.registeredWatchers); i++ {
		if ws.registeredWatchers[i] == w {
			ws.registeredWatchers = append(ws.registeredWatchers[0:i], ws.registeredWatchers[i+1:]...)
		}
	}
}

func (ws *WeatherStreamer) NotifyWatchers() {
	for _, watcher := range ws.registeredWatchers {
		watcher.Update(ws.data)
	}
}

func (ws *WeatherStreamer) UpdateData(data models.WeatherData) {
	ws.data = data
	ws.NotifyWatchers()
}
