package pnplib

import "fmt"

type Users struct {
	UsersResponse []UsersResponse `json:"response"`
	URL           string
}

type UsersResponse struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	IsClosed        bool   `json:"is_closed"`
	CanAccessClosed bool   `json:"can_access_closed"`
	MobilePhone     string `json:"mobile_phone"`
}

func (user *UsersResponse) GetLink() string {
	return fmt.Sprintf("https://vk.com/id%d", user.ID)
}

type Members struct {
	MembersResponse `json:"response"`
	URL             string
}

type MembersResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}
