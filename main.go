package main

import (
	"embed"
	"saintspace/app/logger"
	"saintspace/app/runners/server"
)

//go:embed dist/*
var dist embed.FS

func main() {

	// Set up the logger
	loggerService, err := logger.New(false)
	if err != nil {
		panic(err)
	}

	server := server.New(loggerService, dist)
	err = server.Run()
	if err != nil {
		panic(err)
	}
}
