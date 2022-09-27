// Package internal contains internal
package internal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/run"
)

func NewRunGroup() *RunGroup {
	return &RunGroup{g: &run.Group{}}
}

type RunGroup struct {
	g *run.Group
}

func (o *RunGroup) Add(execute func() error, interrupt func(error)) {
	o.g.Add(execute, interrupt)
}

func (o *RunGroup) Run(onShutdown func(error)) error {
	sigs := make(chan os.Signal, 1)
	allDone := make(chan struct{})
	o.g.Add(func() error {
		signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		for {
			select {
			case <-sigs:
				return nil
			case <-allDone:
				return nil
			}
		}
	}, func(err error) {
		onShutdown(err)
		select {
		case allDone <- struct{}{}:
		default:
		}
	})

	return o.g.Run()
}
