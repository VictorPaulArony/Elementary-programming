package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// function to let the user register in the system
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		RenderTemplates(w, "login.html", nil)
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			ErrorHandler(w, http.StatusBadRequest)
			return
		}

		// get the user details from the login form
		userName := r.FormValue("userName")
		userEmail := r.FormValue("userEmail")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirmPassword")

		// check if password matches
		if password != confirmPassword {
			http.Error(w, "Password does not match", http.StatusUnauthorized)
			return
		}

		// Validate that no fields are empty
		if userName == "" || userEmail == "" || password == "" || confirmPassword == "" {
			log.Printf("Empty fields detected")
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		// Begin transaction
		tx, err := DB.Begin()
		if err != nil {
			log.Printf("Error starting transaction: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		// check if the user email is already in the db
		var emailExist string
		err = DB.QueryRow("SELECT useremail FROM users WHERE useremail = ?", userEmail).Scan(&emailExist)
		if err == nil {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		} else if err != sql.ErrNoRows {
			http.Error(w, "Error checking email", http.StatusInternalServerError)
			return
		}

		// Check if the username already exists in the database
		var existingUsername string
		err = DB.QueryRow("SELECT username FROM users WHERE username = ?", userName).Scan(&existingUsername)
		if err == nil {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		} else if err != sql.ErrNoRows {
			http.Error(w, "Error checking username", http.StatusInternalServerError)
			return
		}

		// hash the user password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error encrypting password", http.StatusInternalServerError)
			return
		}

		// generate the uuid for a user
		userId := uuid.New().String()

		// inserting the user into the user database
		if _, err = DB.Exec("INSERT INTO users (id, useremail, username, password) VALUES (?, ?, ?, ?)", userId, userEmail, userName, string(hashedPassword)); err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// Commit the transaction
		if err = tx.Commit(); err != nil {
			log.Printf("Error committing transaction: %v", err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}
		
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
