package entity

//go:generate easyjson -all test.go

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
