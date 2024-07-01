package netcat_user

import (
	"colors"
	"net"
)

type User struct {
	Connection  net.Conn
	Name, Color string
}

func NewUser(c net.Conn, name, color string) *User {
	usr := new(User)
	usr.Connection = c
	usr.Name = name
	usr.Color = color
	return usr
}

func IsValidUsername(name string) bool {
	return name != ""

	/*
		if name == "" {
			return false
		}

		return true
	*/
}

func (usr *User) ColoredUsername() string {
	if usr.Color == "" {
		return usr.Name
	}
	return usr.Color + usr.Name + colors.ResetColorsTag
}
