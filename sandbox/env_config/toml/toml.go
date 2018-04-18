package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"path/filepath"
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
	abs_path, _ := filepath.Abs(".")
	_, err := toml.DecodeFile(filepath.Join(abs_path, "sandbox/env_config/toml/config.toml"), &config)
	if err != nil {
		log.Fatal(err)
	}

	// outputs
	fmt.Printf("channel_secret is %s\n", config.LineBotConfig.ChannelSecret)
	fmt.Printf("channel_token is %s\n", config.LineBotConfig.ChannelToken)
}
