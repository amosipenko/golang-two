// Package main implements functions to demonstrate how to work with panics.
//
// The CreateUser creates new User and returns his name:
//
// CreateUser() (string, error)
//
package main

import (
	"fmt"
)

// user is used to store new user's information
type user struct {
	name string
	age  int
}

func (u *user) getName() string {
	return u.name
}

func main() {
	name, err := CreateUser()
	if err != nil {
		err = fmt.Errorf("Couldn`t get user`s name. %s", err)
		fmt.Println(err)
	} else {
		fmt.Println("New user: ", name)
	}
}

// CreateUser creates new User and returns his name
func CreateUser() (name string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	user := &user{"Sam", 30}
	user = nil
	return user.getName(), nil
}
