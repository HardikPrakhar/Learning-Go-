package main

import (
	"echo-server/config"
	"echo-server/server"
	"flag"
	"log"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "")
	flag.IntVar(&config.Port, "port", 7379, "")
	flag.Parse()
}

func main() {
	setupFlags()

	log.Println("Starting TCP Echo Server...")
	server.Run()
}
