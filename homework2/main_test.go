package main

import "fmt"

func ExampleCreateUser() {
	name, err := CreateUser()
	if err != nil {
		err = fmt.Errorf("Couldn`t get user`s name. %s", err)
		fmt.Println(err)
	} else {
		fmt.Println("New user: ", name)
	}
}
