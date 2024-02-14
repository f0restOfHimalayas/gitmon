package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kardianos/service"

	"github.com/f0restOfHimalayas/gitmon"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	for {

		<-time.After(time.Second * 3)
		config := gitmon.LoadConfig()
		fmt.Printf("hello...%v", config)
	}
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "gitmon",
		DisplayName: "gitmon",
		Description: "Monitors git repos for upstream changes",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
