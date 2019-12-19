//package main
//
//import (
//	"github.com/abice/go-enum/generator/assets"
//	"os"
//	"net/http"	
//	"fmt"
//	"github.com/gorilla/mux"
//	"html/template"
//	"time"
//	"os/signal"
//	"context"
//)
//
////var HTMLServer server
//var navigationBarHTML string
//var homepageTpl *template.Template
//var secondViewTpl *template.Template
//var thirdViewTpl *template.Template
////var Config
//
//func main() {
//	
//	serverCfg := Config{
//		Host:         "localhost:8080",
//		ReadTimeout:  5 * time.Second,
//		WriteTimeout: 5 * time.Second,
//	}
//	
//	htmlServer := Start(serverCfg)
//	defer htmlServer.Stop()
//
//	sigChan := make(chan os.Signal, 1)
//	signal.Notify(sigChan, os.Interrupt)
//	<-sigChan
//
//	fmt.Println("main : shutting down")
//}
//
//func init() {
//	navigationBarHTML = assets.MustAssetString("templates/navigation_bar.html")
//
//	homepageHTML := assets.MustAssetString("templates/index.html")
//	homepageTpl = template.Must(template.New("homepage_view").Parse(homepageHTML))
//
////	secondViewHTML := assets.MustAssetString("templates/second_view.html")
////	secondViewTpl = template.Must(template.New("second_view").Parse(secondViewHTML))
////	
////	thirdViewFuncMap := ThirdViewFormattingFuncMap()
////	thirdViewHTML := assets.MustAssetString("templates/third_view.html")
////	thirdViewTpl = template.Must(template.New("third_view").Funcs(thirdViewFuncMap).Parse(thirdViewHTML))
//
//}
//
//// Start launches the HTML Server
//func Start(cfg Config) *HTMLServer {
//	// Setup Context
//	_, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	// Setup Handlers
//	router := mux.NewRouter()
//	router.HandleFunc("/", HomeHandler)
//	router.HandleFunc("/second", SecondHandler)
//	router.HandleFunc("/third/{number}", ThirdHandler)
//	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
//
//	// Create the HTML Server
//	htmlServer := HTMLServer{
//		server: &http.Server{
//			Addr:           cfg.Host,
//			Handler:        router,
//			ReadTimeout:    cfg.ReadTimeout,
//			WriteTimeout:   cfg.WriteTimeout,
//			MaxHeaderBytes: 1 << 20,
//		},
//	}
//
//	// Add to the WaitGroup for the listener goroutine
//	htmlServer.wg.Add(1)
//
//	// Start the listener
//	go func() {
//		fmt.Printf("\nHTMLServer : Service started : Host=%v\n", cfg.Host)
//		htmlServer.server.ListenAndServe()
//		htmlServer.wg.Done()
//	}()
//
//	return &htmlServer
//}
//
//func (htmlServer *HTMLServer) Stop() error {
//	// Create a context to attempt a graceful 5 second shutdown.
//	const timeout = 5 * time.Second
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	defer cancel()
//
//	fmt.Printf("\nHTMLServer : Service stopping\n")
//
//	// Attempt the graceful shutdown by closing the listener
//	// and completing all inflight requests
//	if err := htmlServer.server.Shutdown(ctx); err != nil {
//		// Looks like we timed out on the graceful shutdown. Force close.
//		if err := htmlServer.server.Close(); err != nil {
//			fmt.Printf("\nHTMLServer : Service stopping : Error=%v\n", err)
//			return err
//		}
//	}
//
//	// Wait for the listener to report that it is closed.
//	htmlServer.wg.Wait()
//	fmt.Printf("\nHTMLServer : Stopped\n")
//	return nil
//}
//
//func HomeHandler(w http.ResponseWriter, r *http.Request) {
//	push(w, "/static/style.css", "style")
//	push(w, "/static/navigation_bar.css", "style")
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//
//	fullData := map[string]interface{}{
//		"NavigationBar": template.HTML(navigationBarHTML),
//	}
//	render(w, r, homepageTpl, "homepage_view", fullData)
//}
//
// SecondHandler renders the second view template
//func SecondHandler(w http.ResponseWriter, r *http.Request) {
//	push(w, "/static/style.css", "style")
//	push(w, "/static/navigation_bar.css", "style")
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//
//	fullData := map[string]interface{}{
//		"NavigationBar": template.HTML(navigationBarHTML),
//	}
//	render(w, r, secondViewTpl, "second_view", fullData)
//}
//
//// ThirdHandler renders the third view template
//func ThirdHandler(w http.ResponseWriter, r *http.Request) {
//	push(w, "/static/style.css", "style")
//	push(w, "/static/navigation_bar.css", "style")
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//
//	var queryString string
//	pathVariables := mux.Vars(r)
//	queryNumber, err := strconv.Atoi(pathVariables["number"])
//	if err != nil {
//		queryString = pathVariables["number"]
//	}
//	fullData := map[string]interface{}{
//		"NavigationBar": template.HTML(navigationBarHTML),
//		"Number":        queryNumber,
//		"StringQuery":   queryString,
//	}
//	render(w, r, thirdViewTpl, "third_view", fullData)
//}
//
//// Push the given resource to the client.
//func push(w http.ResponseWriter, resource string) {
//	pusher, ok := w.(http.Pusher)
//	if ok {
//		if err := pusher.Push(resource, nil); err == nil {
//			return
//		}
//	}
//}




package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "html/template"
    "log"
)

var tmpl = template.Must(template.ParseGlob("../templates/*.html"))

type serverConfig struct  {
		host string      
		readTimeout int  
		writeTimeout int 
}

type Page struct {
    NavigationBar string
}

func main() { 
	
	var config = ServerConfig("localhost:8080")
	
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

func ServerConfig(string host) *serverConfig{

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


