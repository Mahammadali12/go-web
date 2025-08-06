package main

import (
	// "encoding/json"
	"log"
	"net/http"
	"time"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

)

type App struct{
    config Config
}

type Config struct{
    Addr string
}

var users = []User{}

func (app *App) mount() http.Handler{
    r := chi.NewRouter()

    r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

    r.Get("/register", app.getRegisterPage)
    r.Post("/register", app.handleUserRegister)

    r.Get("/login", app.getLoginPage)

    
	fs := http.StripPrefix("/static/",http.FileServer(http.Dir("static/")))
    r.Handle("/static/*",fs)
    
    return  r
}

func (app *App) run(mux http.Handler) error{


    srv := &http.Server{
        Addr: app.config.Addr,
        Handler: mux,
        WriteTimeout: time.Second*30,
        ReadTimeout: time.Second*10,
        IdleTimeout: time.Minute,
    }

    log.Printf("Server started at %s",app.config.Addr)

    return srv.ListenAndServe()
}


// func (a *App) createUserHandler(w http.ResponseWriter, r *http.Request){

//     var payload User

//     err := json.NewDecoder(r.Body).Decode(&payload)

//     if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//     }

//     u := User{
//         FirstName: payload.FirstName,
//         LastName: payload.LastName,
//     }

//     users = append(users,u)

//     w.WriteHeader(http.StatusCreated)
// }

// func (a *App) getUserHandler(w http.ResponseWriter, r *http.Request){
//     w.Header().Set("Content-Type","application/json")

//     err := json.NewEncoder(w).Encode(users)

//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     w.WriteHeader(http.StatusOK)
// }