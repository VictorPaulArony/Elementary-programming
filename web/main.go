package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	// Dummy user data
	username = "admin"
	password = "password"
)

func main() {
	// Serve the login page
	http.HandleFunc("/", serveLoginPage)

	// Handle login form submission
	http.HandleFunc("/login", handleLogin)

	// Start the web server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// serveLoginPage serves the login HTML page
func serveLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("login.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// handleLogin processes the login form submission
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get username and password from form
	inputUsername := r.FormValue("username")
	inputPassword := r.FormValue("password")

	// Check credentials
	if inputUsername == username && inputPassword == password {
		fmt.Fprintf(w, "Login successful!")
	} else {
		fmt.Fprintf(w, "Invalid username or password")
	}
}
