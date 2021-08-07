package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

type WanIP struct {
	IP string `json:"ip"`
}

func CheckIP(ip WanIP, dynHost DynHost) bool {
	ips, err := net.LookupIP(dynHost.Domain)

	if err != nil {
		return false
	}

	for i := 0; i < len(ips); i++ {
		ipv4 := ips[i].To4()
		if ipv4 != nil && ipv4.String() == ip.IP {
			return true
		}
	}

	return false
}

func GetWanIP() (WanIP, error) {
	resp, err := http.Get("https://api.ipify.org?format=json")

	var api WanIP
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	parserError := json.Unmarshal(body, &api)

	if parserError != nil {
		return api, errors.New("Cannot get ip wan")
	}
	return api, nil
}
