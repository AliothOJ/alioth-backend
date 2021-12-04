package main

import (
	"flag"
	"fmt"
	"github.com/AliothOJ/backend/internal/config"
	"log"
)

type LaunchOption struct {
	ConfPath string
}

func ParseFlags() (*LaunchOption, error) {
	option := new(LaunchOption)
	flag.StringVar(&(option.ConfPath), "config", "../../config.yaml", "path to yaml file")
	flag.Parse()
	if err:=config.CheckExists(option.ConfPath); err != nil {
		return nil, err
	}
	return option, nil
}

func main() {
	option, err :=  ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}
	Config, err := config.LoadConfig(option.ConfPath)
	if err != nil {
		log.Fatalln(err)
	}
	// test print config
	fmt.Println(Config.Server.Port)
	fmt.Println(Config.Microservices[0].Slaves[0].GRPCPort)

}
