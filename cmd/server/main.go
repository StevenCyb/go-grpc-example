package main

import (
	"gogrpc/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const listen = "0.0.0.0:50051"

func main() {
	s, err := server.New(listen)
	AssertNoErr("failed to create server: %w", err)

	defer s.Close()

	go func() {
		AssertNoErr("failed to start server: %w", s.Start())
	}()

	log.Printf("Listen on: %s", listen)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}

func AssertNoErr(format string, err error) {
	if err != nil {
		log.Fatalf(format, err)
	}
}
