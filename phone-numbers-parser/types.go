package phone_numbers_parser

type Users struct {
	Response []struct {
		ID              int    `json:"id"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		IsClosed        bool   `json:"is_closed"`
		CanAccessClosed bool   `json:"can_access_closed"`
		MobilePhone     string `json:"mobile_phone"`
	} `json:"response"`
	URL string
}

type Members struct {
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
	URL string
}
