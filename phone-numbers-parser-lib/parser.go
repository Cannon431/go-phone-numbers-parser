package pnplib

import (
	"fmt"
	"github.com/Cannon431/go-vk-api"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
	"time"
)

type Parser struct {
	file    *os.File
	api     API
	request Request
	filters Filters
}

type API struct {
	vk vkapi.VK
}

type Request struct {
	timeout         int
	usersPerRequest int
}

type Filters struct {
	minDigitsCount int
	ignored        []string
}

func New(fileName string, iniFile *ini.File) *Parser {
	dumpsDir := iniFile.Section("Dumps").Key("dumps_dir").String()
	if _, err := os.Stat(dumpsDir); err != nil {
		if err = os.Mkdir(dumpsDir, 0777); err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create(dumpsDir + fileName + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	api := API{
		vk: vkapi.VK{
			Token:   iniFile.Section("VK_API").Key("access_token").String(),
			Version: iniFile.Section("VK_API").Key("v").String(),
			Lang:    iniFile.Section("VK_API").Key("lang").String(),
		},
	}

	request := Request{
		timeout:         iniFile.Section("Request").Key("timeout").MustInt(),
		usersPerRequest: iniFile.Section("Request").Key("users_per_request").MustInt(),
	}

	filters := Filters{
		minDigitsCount: iniFile.Section("Filters").Key("min_digits_count").MustInt(),
		ignored:        iniFile.Section("Filters").Key("ignored").Strings(","),
	}

	return &Parser{
		file:    file,
		api:     api,
		request: request,
		filters: filters,
	}
}

func (p *Parser) write(user UsersResponse) error {
	text := fmt.Sprintf(
		"%s %s (%s): %s\n",
		user.FirstName,
		user.LastName,
		user.GetLink(),
		user.MobilePhone,
	)

	_, err := p.file.Write([]byte(text))

	return err
}

func (p *Parser) mustWrite(text UsersResponse) {
	if err := p.write(text); err != nil {
		panic(err)
	}
}

func (p *Parser) filter(number string) (bool, error) {
	if number == "" {
		return false, nil
	}

	number = strings.TrimSpace(number)

	if countOfDigits(number) < p.filters.minDigitsCount {
		return false, nil
	}

	for _, filter := range p.filters.ignored {
		if number == filter {
			return false, nil
		}
	}

	return true, nil
}

func (p *Parser) mustFilter(number string) bool {
	correct, err := p.filter(number)

	if err != nil {
		panic(err)
	}

	return correct
}

func (p *Parser) Parse(groupID string) error {
	log.Println("Start parsing")

	members := p.api.mustGetMembers(groupID, 0)

	phoneNumbers := 0
	processed := 0
	offset := 0
	count := members.MembersResponse.Count

	log.Printf("Processed 0/%d users\n", count)

	for offset < count {
		members = p.api.mustGetMembers(groupID, offset)
		users := p.api.mustGetUsers(members.MembersResponse.Items, p.request.usersPerRequest)

		for _, user := range users.UsersResponse {
			if p.mustFilter(user.MobilePhone) {
				p.mustWrite(user)
				phoneNumbers++

				log.Println("Got", phoneNumbers, "phone numbers")
			}
		}

		processed += p.request.usersPerRequest

		if processed > count {
			processed = count
		}

		log.Printf("Processed %d/%d users\n", processed, count)

		offset += p.request.usersPerRequest
		time.Sleep(time.Millisecond * time.Duration(p.request.timeout))
	}

	log.Println("Done!")

	return nil
}
