package main

import (
	"github.com/kinxyo/knowledge-box/htmx-go/pkg/handlers"
	"net/http"
)

func main() {
	// Serve static files
	http.Handle("/web/static/css/", http.StripPrefix("/web/static/css/", http.FileServer(http.Dir("./web/static/css"))))
	http.Handle("/web/static/js/", http.StripPrefix("/web/static/js/", http.FileServer(http.Dir("./web/static/js"))))

	// Handle routes
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/clicked", handlers.ClickedHandler)
	http.HandleFunc("/messages", handlers.HandlingMessage)
	http.HandleFunc("/channel", handlers.PutChannel)

	// Start the server
	http.ListenAndServe(":3000", nil)
}
