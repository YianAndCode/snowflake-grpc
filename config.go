package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ServerConfig 服务配置
type ServerConfig struct {
	Server string
	Port   int
	NodeId int
}

func getConfig() (config ServerConfig) {
	configJSON, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Read config.json error: ", err)
		config.Server = "127.0.0.1"
		config.Port = 6666
		config.NodeId = 1
	} else {
		err = json.Unmarshal(configJSON, &config)
		if err != nil {
			fmt.Println("Parse json error: ", err)
			config.Server = "127.0.0.1"
			config.Port = 6666
			config.NodeId = 1
		}
	}
	return
}
