package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BraedynL0530/Pinata/internal/core"
)

func main() {
	cfg, err := core.LoadConfig("temp/file/path.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", cfg) // future will be used to start up hardware and such IF its allowed in config

	manager := core.NewLifeCycleManager()
	handleShutdown(manager)

}

func handleShutdown(manager *core.LifeCycleManager) {
	manager.Run()
	manager.Go(func(quit <-chan struct{}) {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-quit:
				//quit in setup gui
				//save workspace location if possible
				//save config
				return // polite kill to each important/all apps
			case <-ticker.C:
				//force kill
			}
		}
	})
}

func InitWeather(endpoint string) (results string, err *error) {
	return "", nil // temp
}
func UpdateWeather(endpoint string) (results string, err *error) {
	return "", nil // use Open-Meteo API
}
