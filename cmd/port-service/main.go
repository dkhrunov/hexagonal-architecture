package main

import (
	"fmt"
	"log"
	"os"
)

func init() {

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	//  read config from env
	// _ := config.Read()
	fmt.Println("Hello world")
	return nil
}
