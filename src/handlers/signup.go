package handlers

import (
	"fmt"
	"net/http"
)

func DisplaySignUpForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
}

func ValidateSignUpForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		DisplaySignUpForm(w, r)
	case "POST":
		ValidateSignUpForm(w, r)
	default:
		DisplaySignUpForm(w, r)
	}

}
