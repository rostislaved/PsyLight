package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func New(configPath string, verboseMod bool) Config {

	v := viper.New()

	dir, file := getDirAndFile(configPath)

	if verboseMod {
		fmt.Println("Use config: " + filepath.Join(dir, file))
	}

	v.SetConfigName(file)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("couldn't load config: %s", err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("couldn't read config: %s", err)
	}

	return config
}

func getDirAndFile(configPath string) (string, string) {
	var dir, file string

	if configPath == "" {
		exePath, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}

		dir = strings.TrimSuffix(exePath, filepath.Base(exePath))
		file = "config"
	} else {
		dir, file = filepath.Split(configPath)

		fileRaw := strings.Split(file, ".")
		file = fileRaw[0]
	}

	return dir, file
}
