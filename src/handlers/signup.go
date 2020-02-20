package handlers

import (
	"fmt"
	"net/http"
)

type SignUpForm struct {
	FieldNames []string
	Fileds     map[string]string
	Errors     map[string]string
}

func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	fmt.Println("get")
	RenderTemplate(w, "./templates/signupform.html", s)
}

func ValidateSignUpForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	s := SignUpForm{}
	s.FieldNames = []string{"username", "email"}
	s.Fileds = make(map[string]string)
	s.Errors = make(map[string]string)

	switch r.Method {

	case "GET":
		DisplaySignUpForm(w, r, &s)
	case "POST":
		ValidateSignUpForm(w, r)
	default:
		DisplaySignUpForm(w, r, &s)
	}

}
