package main

import (
	// `"fmt"
	"fmt"
	// "html/template"

	"log"
    "database/sql"
    "os"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
    // "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"

)

type todo struct{
    Title string
    Done bool
}

type TodoPageData struct{
    Title string
    Todos []todo
}

type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

var(
    key = securecookie.GenerateRandomKey(8)
    store = sessions.NewCookieStore(key)

)

func login(w http.ResponseWriter, r* http.Request)  {

    session, err := store.Get(r, "cookie-name")

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    session.Values["authenticated"] = true

    err = session.Save(r,w)
    fmt.Print("Authenticated\n")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func secret(w http.ResponseWriter, r* http.Request)  {
    session, _ := store.Get(r, "cookie-name")

    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    fmt.Printf("DAYUUUUUM\n")
}


func main() {

    hest()
    r := mux.NewRouter()
    fmt.Print("THE START\n")

    // tmpl := template.Must(template.ParseFiles("form.html"))

    // r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     tmpl.Execute(w, nil)
    // }).Methods("GET")

    // r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

    //     details := ContactDetails{
    //         Email: r.FormValue("email"),
    //         Subject: r.FormValue("subject"),
    //         Message: r.FormValue("message"),
    //     }

    //     fmt.Printf("%v\n",details)

    //     tmpl.Execute(w, struct{Success bool}{true})

    // }).Methods("POST")

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

        vars := mux.Vars(r)
        fmt.Printf("title - %s\n",vars["title"])
        fmt.Printf("page - %s\n",vars["page"])
        fmt.Printf("method - %s\n",r.Method)
    })

    r.HandleFunc("/login",login)

    r.HandleFunc("/secret",secret)

    http.ListenAndServe(":8080", r)
}

func hest()  {
    // Get database connection details from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")

    // Create connection string
    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)

    // Open database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Failed to open database connection:", err)
    }
    defer db.Close()

    // Test connection
    if err := db.Ping(); err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    log.Println("Successfully connected to MySQL database!")
    fmt.Print("THE END\n")
    

    // Set up Gin router
    // r := gin.Default()
    // r.GET("/ping", func(c *gin.Context) {
    //     c.JSON(200, gin.H{"message": "pong"})
    // })

    // Add routes for your workflow automation tool here
    // Example: r.GET("/workflows", workflowsHandler(db))

    // Start server
    // r.Run(":8080") 
}