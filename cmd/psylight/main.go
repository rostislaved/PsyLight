package main

import (
	"flag"
	"fmt"
	"log"
	"my-projects/ambilight/internal/config"
	"my-projects/ambilight/internal/measurer"
	"my-projects/ambilight/internal/modes/ambilight"
	"my-projects/ambilight/internal/port"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	currentDir,err  := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CURRENT PATH:"+currentDir.Name())

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
