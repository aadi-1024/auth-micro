package models

type User struct {
	Uid      int    `json:"uid,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
