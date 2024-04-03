package controllers

import "net/http"

type LogoutController interface {
	Logout(w http.ResponseWriter, r *http.Request)
}
