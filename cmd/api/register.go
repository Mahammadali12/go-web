package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (a *App)getRegisterPage(w http.ResponseWriter, r *http.Request)  {
	tmpl, err := template.ParseFiles("./templates/register.html")


	
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	tmpl.Execute(w,nil)
}

func (a *App) handleUserRegister(w http.ResponseWriter, r *http.Request)  {
	
	U := User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	// fmt.Print(U)
	users = append(users, U)
	fmt.Print(users)
}