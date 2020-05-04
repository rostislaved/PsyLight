package main

import (
	"github.com/pkg/profile"
	"my-projects/ambilight/config"
	"my-projects/ambilight/measurer"
	"my-projects/ambilight/modes/ambilight"
	"my-projects/ambilight/port"
	_ "net/http/pprof"
	"time"
)

func main() {
	defer profile.Start().Stop()
	usb := port.New()

	cfg := config.New()

	mode := ambilight.New(cfg)

	fps := measurer.New()

	//desirableFPS := 150
	//ns := int64(float32(1e3)/float32(desirableFPS)) * 1e6
	//frameTime := time.Duration(ns)
	//fmt.Println(frameTime)
	//t := time.NewTicker(frameTime)
	//for range t.C {
	for {
		fps.Print()
		//t1 = time.Now()

		mode.UpdateState()
		//passed1 = time.Since(t1)

		//t2 = time.Now()
		time.Sleep(cfg.WaitTime)

		usb.Write(mode.Buffer())
		//passed2 = time.Since(t2)
		//fmt.Printf("t1: %v\tt2: %v\n", passed1.Microseconds(), passed2.Microseconds())

	}

}
