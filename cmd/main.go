package main

import (
	"fmt"
	"log"
	"webservice/internal/config"
)

func main() {
	c, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", c)
}
