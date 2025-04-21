package user

import (
	"regexp"
	"strings"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (u User) IsValidEmail() bool {
	return emailRegex.MatchString(u.Email)
}

func (u *User) Normalize() {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
}
