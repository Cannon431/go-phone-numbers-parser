package phone_numbers_parser

import (
	"encoding/json"
	"strconv"
	"vk-api"
)

func getMembers(groupID string, params map[string]string, api vk_api.VK) (Members, error) {
	var members Members
	params["group_id"] = groupID

	data, err := api.Request("groups.getMembers", params)

	if err != nil {
		return members, err
	}

	err = json.Unmarshal(data, &members)

	return members, err
}

func getUsers(IDs []int, count int, api vk_api.VK) (Users, error) {
	var users Users

	data, err := api.Request("users.get", map[string]string{
		"user_ids": joinInts(IDs, ","),
		"fields": "contacts",
		"count": strconv.Itoa(count),
	})

	if err != nil {
		return users, err
	}

	err = json.Unmarshal(data, &users)

	return users, err
}
