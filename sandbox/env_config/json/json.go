package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

type Config struct {
	LineBotConfig LineBotConfig `json:"line_bot_config"`
}

type LineBotConfig struct {
	ChannelSecret string `json:"channel_secret"`
	ChannelToken string `json:"channel_token"`
}

func main() {
	// read config file(json)
	file, err := ioutil.ReadFile("/Users/zyyx-kubo/GoProjects/learning-go/sandbox/env_config/json/config.json")
	if err != nil {
		log.Fatal(err)
	}

	// json unmarshal
	var config Config
	json.Unmarshal(file, &config)

	// outputs
	fmt.Printf("channel_secret is %s\n", config.LineBotConfig.ChannelSecret)
	fmt.Printf("channel_token is %s\n", config.LineBotConfig.ChannelToken)
}
