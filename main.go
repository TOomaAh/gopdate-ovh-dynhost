package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	for {
		config, err := OpenConfig()

		if err != nil {
			fmt.Println(err)
			WriteDefault()
			config, _ = OpenConfig()
		}
		makeUpdate(config)
		fmt.Println("-------------------------------------------------")
		time.Sleep(time.Duration(config.Seconds) * time.Second)
	}
}

func makeUpdate(config Config) {
	ip, err := GetWanIP()

	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, dynHost := range config.Data {
		if CheckIP(ip, dynHost) {
			log.Printf("Domain %s has already this IP: %s\n", dynHost.Domain, ip.IP)
		} else {
			_, err := UpdateDynHost(ip, dynHost)

			if err != nil {
				log.Printf(err.Error(), dynHost.Domain)
				return
			}
			log.Printf("Domain %s => Update to %s\n", dynHost.Domain, ip.IP)
		}
	}
}
