package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type DatabaseConfig struct {
	Host string `mapstructure:"hostname"`
	Port string
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}

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
}

type Config struct {
	WaitTime  time.Duration `mapstructure:"WaitTimeMS"`
	LEDs      LEDs          `mapstructure:"LEDS"`
	Ambilight Ambilight     `mapstructure:"Ambilight"`
}

func New() Config {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("couldn't load config: %s", err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("couldn't read config: %s", err)
	}

	config.WaitTime = config.WaitTime * 1e6

	return config
}
