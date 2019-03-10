package pnplib

import (
	"encoding/json"
	"strconv"
	"github.com/Cannon431/go-vk-api"
)

func (api *API) getMembers(groupID string, offset int) (Members, error) {
	var members Members

	data, err := api.vk.Request(vkapi.POST, "groups.getMembers", map[string]string{
		"group_id": groupID,
		"offset":   strconv.Itoa(offset),
	})

	if err != nil {
		return members, err
	}

	err = json.Unmarshal(data, &members)

	return members, err
}

func (api *API) mustGetMembers(groupID string, offset int) Members {
	members, err := api.getMembers(groupID, offset)

	if err != nil {
		panic(err)
	}

	return members
}

func (api *API) getUsers(IDs []int, count int) (Users, error) {
	var users Users

	data, err := api.vk.Request(vkapi.POST, "users.get", map[string]string{
		"user_ids": joinInts(IDs, ","),
		"fields":   "contacts",
		"count":    strconv.Itoa(count),
	})

	if err != nil {
		return users, err
	}

	err = json.Unmarshal(data, &users)

	return users, err
}

func (api *API) mustGetUsers(IDs []int, count int) Users {
	users, err := api.getUsers(IDs, count)

	if err != nil {
		panic(err)
	}

	return users
}
