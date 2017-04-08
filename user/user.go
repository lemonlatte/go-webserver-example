package user

import (
	"fmt"
	"net/http"
)

type User struct {
}

func NewUserManager() *User {
	return &User{}
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "login")
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "logout")
}
