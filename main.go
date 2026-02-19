package main

import (
	"kautsarhasby/ewallet-ums/cmd"
	"kautsarhasby/ewallet-ums/helpers"
)

func main() {

	helpers.SetupEnv()
	helpers.SetupLogger()
	helpers.SetupDatabase()

	go cmd.ServeGRPC()
	cmd.ServeHTTP()

}
