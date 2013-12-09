package main

import (
	"pilosa/core"
	"pilosa/db"
	"flag"
	"log"
)

var tcpLoc string
var httpLoc string

func init() {
	flag.StringVar(&tcpLoc, "l", "127.0.0.1:1300", "ip:port to listen on (tcp)")
	flag.StringVar(&httpLoc, "h", "127.0.0.1:1400", "ip:port to listen on (http)")
	flag.Parse()
}

func main() {
	tcp, err := db.NewLocation(tcpLoc)
	if err != nil {
		log.Fatal("Location not valid:", tcpLoc)
	}
	http, err := db.NewLocation(httpLoc)
	if err != nil {
		log.Fatal("Location not valid:", httpLoc)
	}

	service := core.NewService(tcp, http)
	service.Run()
}
