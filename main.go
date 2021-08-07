package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var logFile string = "gopdate.log"
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		if err != nil {
			log.Panicln("Cannot read or create logfile")
			os.Exit(-1)
		}
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println("Started")
	for {
		config, err := OpenConfig()

		if err != nil {
			fmt.Println(err)
			WriteDefault()
			config, _ = OpenConfig()
		}
		makeUpdate(config)
		time.Sleep(time.Duration(config.Seconds) * time.Second)
	}
}

func makeUpdate(config Config) {
	ip, err := GetWanIP()

	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(config.Data) == 0 {
		log.Fatalln("No domain specified")
		return
	}

	for i, dynHost := range config.Data {

		if dynHost.Domain == "" {
			log.Printf("No domain specified for the domain %d", i)
			break
		}

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
