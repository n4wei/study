package watcher

import (
	"fmt"
	"study/oo-patterns/observer/models"
)

type WeatherReport struct {
	data models.WeatherData
}

func NewWeatherReport() Watcher {
	return &WeatherReport{}
}

func (wr *WeatherReport) Update(data models.WeatherData) {
	wr.data = data
	wr.Display()
}

func (wr *WeatherReport) Display() {
	fmt.Printf("Weather Report: Temperature %vF, Humidity %v\n", wr.data.Temperature, wr.data.Humidity)
}
