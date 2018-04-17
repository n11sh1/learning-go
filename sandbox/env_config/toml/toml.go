package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	LineBotConfig LineBotConfig `toml:"line_bot_config"`
}

type LineBotConfig struct {
	ChannelSecret string `toml:"channel_secret"`
	ChannelToken string `toml:"channel_token"`
}

func main() {
	// read config file(toml)
	var config Config
	_, err := toml.DecodeFile("/Users/zyyx-kubo/GoProjects/learning-go/sandbox/env_config/toml/config.tml", &config)
	if err != nil {
		log.Fatal(err)
	}

	// outputs
	fmt.Printf("channel_secret is %s\n", config.LineBotConfig.ChannelSecret)
	fmt.Printf("channel_token is %s\n", config.LineBotConfig.ChannelToken)
}
