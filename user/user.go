package user

import "github.com/google/uuid"

type User struct {
	id       string
	username string
}

func newUser(username string, id string) *User {
	return &User{
		username: username,
		id:       id,
	}
}

func CreateUser(username string) *User {
	return &User{
		username: username,
		id:       uuid.NewString(),
	}
}

//Users will need to be stored in a database and retrievable in the future
