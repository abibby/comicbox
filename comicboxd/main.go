// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// simple does nothing except block while running the service.
package main

import (
	"log"
	"os"

	"bitbucket.org/zwzn/comicbox/comicboxd/app/routes"
	"bitbucket.org/zwzn/comicbox/comicboxd/j"
	"bitbucket.org/zwzn/comicbox/comicboxd/server"
	"github.com/kardianos/service"
)

var logger service.Logger

type program struct {
	server *server.Server
}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	p.server = server.New()
	j.Info("Starting server")

	routes.Web(p.server)

	err := p.server.Start()
	if err != nil {
		j.Errorf("error starting server: %v", err)
		os.Exit(1)
	}
}
func (p *program) Stop(s service.Service) error {
	return p.server.Stop()
}

func main() {
	svcConfig := &service.Config{
		Name:        "comicboxd",
		DisplayName: "ComicBox",
		Description: "A simple comic book server for local files",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	j.SetLogger(logger)

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}