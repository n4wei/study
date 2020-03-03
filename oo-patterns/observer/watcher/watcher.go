package watcher

import "study/oo-patterns/observer/models"

type Watcher interface {
	Update(d models.WeatherData)
}
