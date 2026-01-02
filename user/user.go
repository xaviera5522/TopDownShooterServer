package user

//TODO: Refactor to make all of this one big package

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

func (u *User) getId() string {
	return u.id
}

//Users will need to be stored in a database and retrievable in the future
