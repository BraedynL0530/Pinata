package main

import (
	"sync"

	"github.com/BraedynL0530/Pinata/internal/core"
)

type initialize struct {
	wg     sync.WaitGroup
	config *core.Con
}

func main() {

}

func (i *initialize) startup() {
	i.wg.Add(1)
	go func() {
		defer i.wg.Done()
	}()
}
