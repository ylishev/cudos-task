package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"cudos-task/contract"
)

type Shutdown struct {
	ready      atomic.Int32
	inProgress atomic.Bool
	readyChan  chan bool
}

// NewShutdown capture shutdown signals from OS and stop the main application
func NewShutdown(cancel context.CancelFunc) *Shutdown {
	sh := &Shutdown{readyChan: make(chan bool)}
	c := make(chan os.Signal, 1)
	go func() {
		s := <-c
		log.Printf("Shutdown in progress, please wait (%v)\n", s)
		sh.shutdown()
		cancel()
	}()

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	return sh
}

// SetReady allows to define a safety block between SetReady(false) and SetReady(true)
// between which the application can shut down gracefully
func (s *Shutdown) SetReady(ready bool) bool {
	if ready {
		if s.ready.CompareAndSwap(1, 0) {
			if s.inProgress.CompareAndSwap(true, true) {
				if s.ready.CompareAndSwap(0, 2) {
					close(s.readyChan)
				}
			}
			return true
		}
	} else {
		s.ready.CompareAndSwap(0, 1)
	}
	return false
}

func (s *Shutdown) shutdown() {
	if !s.inProgress.CompareAndSwap(false, true) {
		return
	}

	if s.ready.CompareAndSwap(0, 2) {
		close(s.readyChan)
		return
	}

	select {
	case <-time.After(contract.ShutdownMaxWaitTime):
	case <-s.readyChan:
	}
}
