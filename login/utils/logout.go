package handlers

import "net/http"

// function to handle session termination by loging out users
func logoutUser(w http.ResponseWriter, ) {
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // Deletes the cookie
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("Logged out successfully"))
}
