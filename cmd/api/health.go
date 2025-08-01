package main

import "net/http"

func (app *App)  healthCheckHandler (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Health check coming from daldan ....."))
}