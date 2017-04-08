package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lemonlatte/go-webserver-basestack/user"
	"github.com/lemonlatte/go-webserver-basestack/ws"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am home.")
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", index)

	userManager := user.NewUserManager()

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/login", userManager.Login)
	userRouter.HandleFunc("/logout", userManager.Logout)

	wsManager := ws.NewWebsocketManager()

	wsRouter := r.PathPrefix("/ws").Subrouter()
	wsRouter.HandleFunc("/echo", wsManager.Echo)

	log.Fatal(http.ListenAndServe(":8000", r))
}
