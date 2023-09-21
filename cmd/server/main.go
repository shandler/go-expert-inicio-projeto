package main

import "github.com/shandler/go-expert-inicio-projeto/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)

}
