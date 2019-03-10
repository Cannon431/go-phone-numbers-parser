package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"phone-numbers-parser/phone-numbers-parser"
	"vk-api"
)

func main() {
	iniFile, err := ini.Load("config.ini")

	if err != nil {
		log.Fatal(err)
	}

	dumpsDir := iniFile.Section("Dumps").Key("dumps_dir").String()

	token := iniFile.Section("VK_API").Key("access_token").String()
	v := iniFile.Section("VK_API").Key("v").String()
	lang := iniFile.Section("VK_API").Key("lang").String()

	ignored := iniFile.Section("Filters").Key("ignored").String()
	minDigitsCount := iniFile.Section("Filters").Key("min_digits_count").String()

	err = os.Setenv("ignored", ignored)

	if err != nil {
		log.Fatal(err)
	}

	err = os.Setenv("min_digits_count", minDigitsCount)

	if err != nil {
		log.Fatal(err)
	}

	var groupID string
	fmt.Print("Enter group ID: ")
	_, err = fmt.Scanln(&groupID)

	if err != nil {
		log.Fatal(err)
	}

	if _, err = os.Stat(dumpsDir); err != nil {
		err = os.Mkdir(dumpsDir, 0777)

		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create(dumpsDir + groupID + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start parsing")

	err = phone_numbers_parser.Parse(groupID, file, vk_api.VK{
		Token: token,
		Version: v,
		Lang: lang,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}

