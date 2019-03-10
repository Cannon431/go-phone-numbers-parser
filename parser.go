package main

import (
	"./phone-numbers-parser-lib"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

const configFileName = "config.ini"

func main() {
	iniFile, err := ini.Load(configFileName)

	if err != nil {
		log.Fatal(err)
	}

	var groupID string
	fmt.Print("Enter group ID: ")
	if _, err = fmt.Scanln(&groupID); err != nil {
		log.Fatal(err)
	}

	parser := pnplib.New(groupID+".txt", iniFile)

	if err = parser.Parse(groupID); err != nil {
		log.Fatal(err)
	}
}
