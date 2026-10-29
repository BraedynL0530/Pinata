package core

import (
	"image/color"
	"time"
)

type Config struct {
	Visual   VisualConfig
	AI       AIConfig
	Hardware HardwareConfig
	Network  NetworkConfig
	Extra    Additional
}

type FileExist bool

type VisualConfig struct {
	PrimaryColor   *color.Color //default light borderline violet blue
	SecondaryColor *color.Color //default pastel pink
	TertiaryColor  *color.Color //default black(or dark grey but basically black
	AccentColor    *color.Color // light grey just to contrast bgs and such
	ImagePath      string       //"assets/BG.webp" for now!

	//color palette here
	// think borders, background, accent colors, etc whatever else i think up later
	//ascii window's art(default will be momonga)
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

type Additional struct {
	Location     string // optional or else na
	MetricSystem bool   //true = metric false is freedom units
}

func NewConfig() *Config { // default*
	return &Config{
		Visual: VisualConfig{ // current real values
			ImagePath: "TEMP!/assets/BG.webp",
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
			PCCommPort: 0, // who would've guessed temp
		},
		Extra: Additional{ // am i stupid why is this a bug?
			Location:     "na",
			MetricSystem: true,
		},
	}
}

func LoadConfig(path string) (config *Config, err error) {

	if err != nil {
		return nil, err
	}
	return &Config{
		//temp
	}, nil
}

func (c *Config) SearchConfig(path string) (*Config, error) { return nil, nil } // search for USER config
