package models

type User struct {	
	User_id        int     `json:"user_id"`
	User_name      string  `json:"user_name"`
	user_pass      string  `json:"user_pass"`
	User_type      string  `json:"user_type"`
	User_mobile    string  `json:"user_mobile"`
	Add_time  	   string  `json:"add_time"`
}

type  Users []User