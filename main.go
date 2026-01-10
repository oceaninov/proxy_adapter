package main

import (
	"log"
	"proxy-adapter/cmd"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Panicf("Set timezone error: %s", err)
	}
	time.Local = loc

	cmd.Execute()
}
