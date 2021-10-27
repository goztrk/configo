package main

import (
	"github.com/Quanthir/configo"
)

type config struct {
	addr string `json:"server-address"`
	port string `json:"server-port"`
}

var conf *config

func main() {
	// These are default values
	conf = &config{
		addr: "127.0.0.1",
		port: "80",
	}

	_ = configo.Add("server", conf, false)

	// Loads config from file and if file does not exists,
	// it saves the file with default values.
	_ = configo.Load("server")

	// now conf has updated values from `server.json` file
}
