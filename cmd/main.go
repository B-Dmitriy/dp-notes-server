package main

import (
	"fmt"
	"log"
	"webservice/internal/config"
	"webservice/internal/server"
)

func main() {
	c, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = server.NewServer(c.Host, c.Port, c.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v \n", c)
}
