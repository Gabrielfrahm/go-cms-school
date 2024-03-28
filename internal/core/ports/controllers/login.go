package controllers

import "net/http"

type LoginController interface {
	Login(w http.ResponseWriter, r *http.Request)
}
