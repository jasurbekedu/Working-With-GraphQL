// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Username string `json:"username"`
}

type Profession struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID         int         `json:"id"`
	Username   string      `json:"username"`
	Profession *Profession `json:"profession"`
}
