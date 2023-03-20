package handlers

import (
	"net/http"
	"usersenaryo/userlogic"
)

func ServEt() {
	http.HandleFunc("/create-user", userlogic.HandleUserCreate)
	http.HandleFunc("/get-user/", userlogic.HandleGetUser)
	http.HandleFunc("/delete-user/", userlogic.HandleUserDelete)
	http.HandleFunc("/update-user/", userlogic.HandleUpdateUser)
	http.ListenAndServe(":8000", nil)
}
