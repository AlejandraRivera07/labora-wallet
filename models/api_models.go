package models

type API_Request struct {
	National_id     string `json:"national_id"`
	Country         string `json:"country"`
	Person_type     string `json:"type"`
	User_authorized bool   `json:"user_authorized"`
}
