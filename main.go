package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/selfharm-enj/otus_m2_hw8/internal/model"
	"github.com/selfharm-enj/otus_m2_hw8/internal/repository"
	"github.com/selfharm-enj/otus_m2_hw8/internal/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	dataCh := make(chan model.IDReader)

	service.StartService(dataCh)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dataCh)
				return
			case data, ok := <-dataCh:
				if !ok {
					return
				}
				repository.AddData(data)
			}
		}
	}()

	go func() {
		service.LogChanges()
	}()

	<-sigCh
	cancel()

	time.Sleep(500 * time.Millisecond)
}
