package panics

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func (u *user) getName() string {
	return u.name
}

func main() {
	name, err := createUser()
	if err != nil {
		err = fmt.Errorf("Couldn`t get user`s name. %s", err)
		fmt.Println(err)
	} else {
		fmt.Println("New user: ", name)
	}
}

func createUser() (name string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	user := &user{"Sam", 30}
	user = nil
	return user.getName(), nil
}
