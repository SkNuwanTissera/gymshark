package main

import (
	"github.com/SkNuwanTissera/gymshark/internal/packer"
	"github.com/SkNuwanTissera/gymshark/internal/server"
	"golang.org/x/exp/slog"
)

const restAPIPort = "8080"

func main() {
	bootstrap()
}

func bootstrap() {
	newSizerSrvc := packer.NewSizerService(packer.SortedSizes)
	newPackerSrvc := packer.NewPacketsService()

	newServer := server.NewServer(newSizerSrvc, newPackerSrvc)
	err := newServer.Serve(restAPIPort)
	if err != nil {
		slog.Error(err.Error())
	}
}
