package phone_numbers_parser

import (
	"log"
	"os"
	"strconv"
	"time"
	"vk-api"
)

func write(file *os.File, text string) error {
	text += "\n"
	_, err := file.Write([]byte(text))

	return err
}

func mustWrite(file *os.File, text string) {
	err := write(file, text)

	if err != nil {
		panic(err)
	}
}

func Parse(groupID string, file *os.File, api vk_api.VK) error {
	const usersPerRequest = 1000
	members, err := getMembers(groupID, map[string]string{"count": "1"}, api)

	if err != nil {
		return err
	}

	phoneNumbers := 0
	processed := 0
	offset := 0
	count := members.Response.Count

	for offset < count {
		members, err := getMembers(groupID, map[string]string{"offset": strconv.Itoa(offset)}, api)

		if err != nil {
			return err
		}

		users, err := getUsers(members.Response.Items, usersPerRequest, api)

		if err != nil {
			return err
		}

		for _, user := range users.Response {
			if user.MobilePhone != "" && mustFilter(user.MobilePhone) {
				mustWrite(file, user.MobilePhone)

				phoneNumbers++
				log.Println("Got", phoneNumbers, "phone numbers")
			}
		}

		processed += len(members.Response.Items)

		log.Printf("Processed %d/%d users\n", processed, count)

		offset += usersPerRequest
		time.Sleep(time.Millisecond * 500)
	}

	return nil
}
