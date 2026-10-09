package core

import "time"

type Config struct {
	Visual   VisualConfig
	AI       AIConfig
	Hardware HardwareConfig
	Network  NetworkConfig
}

type VisualConfig struct {
	//color palette here
	// think borders, background, accent colors, etc whatever else i think up later
	//ascii window's art(defualt will be momonga)
}

type AIConfig struct {
	ApiKey                 string
	MachineLearningEnabled bool

	PythonServiceURL string
	InferenceTimeout time.Duration
}

type HardwareConfig struct {
	HasESPSense      bool // Toggles listening/cam usage!
	RoomSensorsLocal bool // tru if wired false if lan/over network
}

type NetworkConfig struct {
	PCCommPort int
}

func NewConfig() *Config { // default*
	return &Config{
		// add more stuff later add default config here, this should only run IF not already one
		// Maybe it makes the file for user as well?
	}
}

func LoadConfig(path string) (*Config, error) {
	//if config detected or pointed at idk
	return nil, nil // temp
}
