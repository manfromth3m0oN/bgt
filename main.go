package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Elec ElecConfig
}

func main() {

	log.Println("Welcome to bgt")
	var config Config
	dat, err := ioutil.ReadFile("bgt.json")
	if err != nil {
		log.Fatalf("Failed to read budget definition: %v", err)
	}
	if err = json.Unmarshal(dat, &config); err != nil {
		log.Fatalf("Failed to marshal budget file: %v", err)
	}
	log.Println(config)
	log.Printf("Total Electricity Cost per Month: %v", config.Elec.calculate())
}
