package core

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type LifeCycleManager struct {
	wg   sync.WaitGroup
	Quit chan struct{}
}

func NewLifeCycleManager() *LifeCycleManager {
	return &LifeCycleManager{Quit: make(chan struct{})}
}

func (m *LifeCycleManager) Go(fn func(quit <-chan struct{})) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		fn(m.Quit)
	}()
}

func (m *LifeCycleManager) Run() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigChan)

	sig := <-sigChan
	fmt.Printf("Received signal: %v\n", sig)
	close(m.Quit) //says stop

	done := make(chan struct{})
	go func() {
		m.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("everything finshed closed clean")
	case <-time.After(10 * time.Second):
		fmt.Println("Timed out force Quit later")
	}
}
