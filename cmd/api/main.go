package main

import (
	// "fmt"
	// "fmt"
	"log"
	// "log"
	// "net/http"
    "github.com/Mahammadali12/go-web/internal/env"
)


func main() {

    cfg := Config{
        Addr: env.GetString("PORT", ":8080"), 
    }

    app := &App{
        config: cfg,
    }

    mux := app.mount()
    
    log.Fatal(app.run(mux))
}
