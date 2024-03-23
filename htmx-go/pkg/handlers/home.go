package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
    "os"
    "io/ioutil"
    "encoding/json"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    
    tmpl, err := template.ParseFiles(filepath.Join("web", "templates", "index.html"))
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func ClickedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Button clicked!")
}

func HandlingMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go said gin")
}

func PutChannel(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("PUT request processed successfully"))
}

func Search(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")

    // Open the JSON file
    jsonFile, err := os.Open("data.json")
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    defer jsonFile.Close()

    // Read the file into a byte array
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // Initialize a slice to hold the data
    var data []Data

    // Unmarshal the byte array into the data slice
    json.Unmarshal(byteValue, &data)

    // Search the data for the query
    for _, item := range data {
        // Replace 'Field' with the field you want to search
        if item.Name == q {
            // Write the search results to the response
            fmt.Fprintf(w, "Search results for: %s", q)
            return
        }
    }

    // If no results were found
    fmt.Fprintf(w, "No results found for: %s", q)
}

type Data struct {
    Name  string `json:"name"`
    Anime string `json:"anime"`
}