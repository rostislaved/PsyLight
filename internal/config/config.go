package config

import (
	"github.com/spf13/viper"
	"log"
)

type OutputConfig struct {
	File string
}

type LEDs struct {
	NumberOfHorizontal int `mapstructure:"numberOfHorizontalLEDs"`
	NumberOfVertical   int `mapstructure:"numberOfVerticalLEDs"`
}

type Ambilight struct {
	HorizontalHeightFraction float32 `mapstructure:"horizontalHeightFraction"`
	VerticalWidthFraction    float32 `mapstructure:"verticalWidthFraction"`
	VerticalOffsetFraction   float32 `mapstructure:"verticalOffsetFraction"`
	HorizontalOffsetFraction float32 `mapstructure:"horizontalOffsetFraction"`
}

type Config struct {
	DesirableFPS float32   `mapstructure:"desirableFPS"`
	LEDs         LEDs      `mapstructure:"LEDS"`
	Ambilight    Ambilight `mapstructure:"Ambilight"`
}

func New(configPath string) Config {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath(configPath)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("couldn't load config: %s", err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("couldn't read config: %s", err)
	}

	return config
}
