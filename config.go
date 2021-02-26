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
		config.Server = "0.0.0.0"
		config.Port = 10001
		config.NodeId = 0
	} else {
		err = json.Unmarshal(configJSON, &config)
		if err != nil {
			fmt.Println("Parse json error: ", err)
			config.Server = "0.0.0.0"
			config.Port = 10001
			config.NodeId = 0
		}
	}
	return
}
