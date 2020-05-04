package main

import (
	"flag"
	"fmt"
	"my-projects/ambilight/config"
	"my-projects/ambilight/measurer"
	"my-projects/ambilight/modes/ambilight"
	"my-projects/ambilight/port"
	_ "net/http/pprof"
	"time"
)

func main() {
	var verboseMod = flag.Bool("v", false, "Show FPS information")
	var configPath = flag.String("c", ".", "Set config path")

	flag.Parse()

	cfg := config.New(*configPath)

	usb := port.New()

	mode := ambilight.New(cfg)

	fps := measurer.New()

	ns := int64(1e9 / cfg.DesirableFPS)
	frameTime := time.Duration(ns)

	fmt.Println("Desired FPS:", cfg.DesirableFPS, "  =>  Frame time:", frameTime)
	t := time.NewTicker(frameTime)
	for range t.C {
		if *verboseMod {
			fps.Print()

		}

		mode.UpdateState()

		usb.Write(mode.Buffer())
	}

}
