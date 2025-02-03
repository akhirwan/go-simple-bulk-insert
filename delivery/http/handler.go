package http

import (
	"go-simple-bulk-insert/delivery/container"
	"go-simple-bulk-insert/domain/counter"
)

type handler struct {
	counterHandler counter.CounterHandler
}

func SetupHandler(container container.Container) handler {
	return handler{
		counterHandler: counter.NewCounterHandler(container.CounterFeature),
	}
}
