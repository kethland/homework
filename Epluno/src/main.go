package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "html/template"
    "log"
    "time"
)

var tmpl = template.Must(template.ParseGlob("../templates/*.html"))

type serverConfig struct  {
		host string      
		readTimeout time.Duration  
		writeTimeout time.Duration
}

type Page struct {
    NavigationBar string
}

func main() { 
	
	ServerConfig("localhost:8080")
	
    router := mux.NewRouter()
    router.HandleFunc("/", HomeHandler)
    router.HandleFunc("/second", SecondHandler)
    //router.HandleFunc("/third/{number}", ThirdHandler)
    //router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    fmt.Println("HTTP Server starting on port 8080")

    err := http.ListenAndServe(":8080", router);
    if err != nil {
        log.Fatal(err)
    }
    
}

func ServerConfig(host string) *serverConfig{

	config := serverConfig{host: host}
	config.readTimeout = 5 * time.Second
	config.writeTimeout = 5 * time.Second
	
	return &config
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Home Page Invoked.")
    p := Page{NavigationBar: `Home`}
    err := tmpl.ExecuteTemplate(w, "HomeHandler.html", p)
    if err != nil {
        log.Fatal("Cannot Get View ", err)
    }
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Second Page Invoked.")
    p := Page{NavigationBar: `Second Page`}
    err := tmpl.ExecuteTemplate(w, "SecondView.html", p)
    if err != nil {
        log.Fatal("Cannot Get View ", err)
    }
    
}


