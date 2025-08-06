package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func (a *App)  getLoginPage(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("./templates/login.html")
	
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	tmpl.Execute(w,nil)
}

func (a *App)  handleLogin(w http.ResponseWriter, r *http.Request){
	key := securecookie.GenerateRandomKey(32)
	store := sessions.NewCookieStore(key)

	session, _ := store.Get(r,"cookie-name")

	// Authentication


	session.Values["authenticated"] = true
	session.Save(r,w)
}