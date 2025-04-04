package main

import (
	"art/ctools" // import ctools package for encoding and decoding functions
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// PageData structures data to be used in the server's HTML templates
type PageData struct {
	Input      string
	Result     string
	Error      string
	HttpStatus int
}

func main() {
	http.HandleFunc("/", indexHandler)                                                          // Route for the main page
	http.HandleFunc("/decoder", decodeHandler)                                                  // Route for decoding requests
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static")))) // Serve static files

	fmt.Println("Server is running on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err) // Fatal log in case the server fails to start
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // Ensure method is GET
		return
	}
	data := PageData{
		HttpStatus: http.StatusOK,
		Result:     "",
	}
	renderResult(w, data) // Render the main page without data

}

// decodeHandler processes decoding and encoding requests
func decodeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // Ensure method is POST
		return
	}

	input := r.FormValue("input")   // get value from the template (user input from text box)
	action := r.FormValue("action") //get value the action (encode or decode) from buttons

	if input == "" {
		data := PageData{
			Error:      "HTTP 400: Input is empty",
			HttpStatus: http.StatusBadRequest,
		}
		renderResult(w, data) // Handle empty input scenario
		return
	}

	var result string
	var err error
	// Execute the appropriate action based on user selection

	if action == "decode" {
		result, err = ctools.DecodeInput(input)
	} else if action == "encode" {
		result, err = ctools.EncodeFromArt(input)
	} else {
		data := PageData{
			Error:      "HTTP 400: Input malformed",
			HttpStatus: http.StatusBadRequest,
		}
		renderResult(w, data)
		return
	}

	if err != nil {
		data := PageData{
			Error:      "HTTP 400: " + err.Error(),
			HttpStatus: http.StatusBadRequest,
		}
		renderResult(w, data) // Handle errors from encoding or decoding
		return
	}
	data := PageData{
		Input:      input,
		Result:     result,
		HttpStatus: http.StatusAccepted,
	}
	renderResult(w, data)
}

// renderResult renders the HTML page using template data
func renderResult(w http.ResponseWriter, data PageData) {
	tmplPath := filepath.Join("static", "index.html")
	templ, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("Error parsing result template %s: %v", tmplPath, err)
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(data.HttpStatus)

	err = templ.Execute(w, data)
	if err != nil {
		log.Printf("Error executing result template %s: %v", tmplPath, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
