package main

import (
	"github.com/react-go/config"
	"github.com/react-go/server"
)

func main() {
	//fmt.Println("hello")
	flags := config.ParseFlags()
	server.Run(flags)
}
