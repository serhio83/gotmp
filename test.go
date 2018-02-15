package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DrHost string
	DrPort string
	DrUser string
	DrPass string
	DrRegs []string
}

func ReadConfig() *Config  {
	var c Config
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.DrHost, c.DrPort, c.DrUser, c.DrPass, c.DrRegs)

	return &c
}

func main() {
	ReadConfig()
}
