package controllers

import "net/http"

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	ListAllUser(w http.ResponseWriter, r *http.Request)
	ListById(w http.ResponseWriter, r *http.Request)
}
