package main

import (
	"context"
	"fmt"
	"gradebook/log"
	"gradebook/registry"
	"gradebook/service"
	stlog "log"
)

func main() {
	logPath := "./gradebook.log"
	log.Run(logPath)

	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	reg := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  serviceAddress,
	}

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		reg,
		log.RegitsterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
