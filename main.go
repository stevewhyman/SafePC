package main

// env CGO_ENABLED=0 go build -ldflags "-s -w" -o nats-api

import (
	"flag"
	"fmt"

	"github.com/wh1te909/tacticalrmm/natsapi"
)

var version = "2.3.0"

func main() {
	ver := flag.Bool("version", false, "Prints version")
	mode := flag.String("m", "", "Mode")
	config := flag.String("c", "", "config file")
	flag.Parse()

	if *ver {
		fmt.Println(version)
		return
	}

	switch *mode {
	case "wmi":
		api.GetWMI(*config)
	case "checkin":
		api.CheckIn(*config)
	case "agentinfo":
		api.AgentInfo(*config)
	default:
		fmt.Println(version)
	}
}
