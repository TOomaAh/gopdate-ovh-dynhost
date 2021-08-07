package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type DynHost struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Seconds int       `json:"seconds"`
	Data    []DynHost `json:"config"`
}

func WriteDefault() error {
	file, err := os.Create("config.json")

	var config Config = Config{
		Seconds: 300,
		Data:    []DynHost{},
	}

	if err != nil {
		fmt.Println("Cannot open config file")
		return errors.New("Cannot open config file")
	}

	b, jsonError := json.Marshal(&config)

	if jsonError != nil {
		fmt.Println("Cannot format config")
		return errors.New("Cannot format config")
	}

	file.WriteString(string(b))

	defer file.Close()

	return nil

}

func OpenConfig() (Config, error) {
	file, err := os.Open("config.json")
	var config Config

	if err != nil {
		return config, errors.New("File not found")
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return config, errors.New("Cannot read file")
	}

	parseError := json.Unmarshal(byteValue, &config)

	if parseError != nil {
		return config, errors.New("Cannot parse config")
	}

	return config, nil
}
