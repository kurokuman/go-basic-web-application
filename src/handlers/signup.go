package handlers

import (
	"fmt"
	"net/http"
	"unicode/utf8"
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

func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	fmt.Println("post")
	if utf8.RuneCountInString(r.FormValue("username")) <= 8 {
		s.Errors["usernameError"] = "The name must be at least eight characters long."
	}

	if r.FormValue("email") == "" {
		s.Errors["emailError"] = "The e-mail address field is required"
	}

	if r.FormValue("pass") != r.FormValue("c_pass") {
		s.Errors["confirmPasswordError"] = "The password and confirm password don't match"
	}

	if len(s.Errors) > 0 {
		DisplaySignUpForm(w, r, s)
	} else {
		ProcessSignUpForm(w, r, s)
	}

}

func ProcessSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	DisplaySignUpForm(w, r, s)

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
		ValidateSignUpForm(w, r, &s)
	default:
		DisplaySignUpForm(w, r, &s)
	}

}
