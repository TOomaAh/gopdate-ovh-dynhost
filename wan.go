package main

import (
	"errors"
	"net"

	externalip "github.com/glendc/go-external-ip"
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
	consensus := externalip.DefaultConsensus(nil, nil)
	api := WanIP{}
	ip, err := consensus.ExternalIP()
	if err != nil {
		return api, errors.New("Cannot get wan ip")
	}
	api.IP = ip.To4().String()
	return api, nil
}
