package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/kardianos/service"

	"github.com/f0restOfHimalayas/gitmon"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	history := make(map[string]string)
	for {
		<-time.After(time.Second * 10)
		projectPathsToMonitor, err := gitmon.LoadConfig()
		if err != nil {
			continue
		}
		fmt.Printf("\nstaring new cycle for all repos: %v", projectPathsToMonitor)

		for _, repo := range projectPathsToMonitor {
			go func(r string) {
				commitId, commits, err := gitmon.FetchLatestCommits(r)
				if err != nil {
					fmt.Errorf("%v", err)
					return
				}
				if val, ok := history[r]; ok {
					if commitId == val {
						fmt.Printf("found in history...")
						return
					}
				}
				err = beeep.Notify(
					fmt.Sprintf("GitMon: New updates: %s", r),
					fmt.Sprintf("\n%s\n%s\n", r, commits),
					"",
				)
				if err != nil {
					fmt.Printf("notification failed with err %v", err)
				} else {
					fmt.Printf("no error while sending notification")
				}
				history[r] = commitId
			}(repo)
		}
	}
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	gitmon.Log("startig gitmon....")
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
