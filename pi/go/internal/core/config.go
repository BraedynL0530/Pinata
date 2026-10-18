package core

import "time"

type Config struct {
	Visual   VisualConfig
	AI       AIConfig
	Hardware HardwareConfig
	Network  NetworkConfig
}

type FileExist bool

type VisualConfig struct {
	PrimaryColor string // temp may need to change value
	AccentColor  string
	TriteryColor string //bs for now
	imagePath    string //"assets/BG.webp" for now!

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

func LoadConfig(path string) *Config {

	return &Config{
		Visual: VisualConfig{ // current real values
			imagePath: "TEMP!/assets/BG.webp",
		},
		AI: AIConfig{ // ALL TEMP VALUES
			ApiKey:                 "",
			MachineLearningEnabled: false,
			PythonServiceURL:       "",
			InferenceTimeout:       60 * time.Second,
		},
		Hardware: HardwareConfig{ // again temp
			HasESPSense:      false,
			RoomSensorsLocal: true,
		},
		Network: NetworkConfig{
			PCCommPort: 0, // who wouldve guessed temp
		},
	}
}

func (c *Config) SearchConfig(path string) (*Config, error) { return nil, nil } // search for USER config
