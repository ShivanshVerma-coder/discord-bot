package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Token string
var Prefix string

type configStruct struct {
	Token  string
	Prefix string
}

func ReadConfig() error {
	var config configStruct
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file")
		return err
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Println("Error parsing config file")
		return err
	}
	Token = config.Token
	Prefix = config.Prefix
	return nil
}
