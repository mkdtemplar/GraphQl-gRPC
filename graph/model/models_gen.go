// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cars struct {
	ID      string `json:"id"`
	CarName string `json:"carName"`
	Model   string `json:"model"`
}

type DeleteUser struct {
	ID string `json:"id"`
}

type NewCar struct {
	CarName string `json:"carName"`
	Model   string `json:"model"`
}

type NewUser struct {
	Name string `json:"name"`
}

type UpdateUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Cars []*Cars `json:"cars"`
}
