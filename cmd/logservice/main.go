package main

import (
	"context"
	"fmt"
	"gradebook/log"
	"gradebook/service"
	stlog "log"
)

func main() {
	logPath := "./gradebook.log"
	log.Run(logPath)

	host, port := "localhost", "4000"

	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegitsterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
