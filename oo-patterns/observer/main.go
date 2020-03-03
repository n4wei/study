package main

import (
	"study/oo-patterns/observer/models"
	"study/oo-patterns/observer/streamer"
	"study/oo-patterns/observer/watcher"
)

func main() {
	ws := streamer.NewWeatherStreamer()
	wr := watcher.NewWeatherReport()
	ws.AddWatcher(wr)

	ws.UpdateData(models.WeatherData{
		Temperature: 32.0,
		Humidity:    0.66,
	})
	ws.UpdateData(models.WeatherData{
		Temperature: 40.0,
		Humidity:    0.80,
	})

	ws.RemoveWatcher(wr)

	ws.UpdateData(models.WeatherData{
		Temperature: 100.0,
		Humidity:    1.00,
	})

	ws.AddWatcher(wr)
	ws.AddWatcher(wr)

	ws.UpdateData(models.WeatherData{
		Temperature: 101.0,
		Humidity:    0.99,
	})
}
