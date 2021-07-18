package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var (
	key = []byte ("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	store = sessions.NewCookieStore(key)
)

func index(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth{
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprint(w, "<h1>Welcome to page ...</h1>")
}

func login(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)

	fmt.Fprint(w, "Login success!")
}

func logout(w http.ResponseWriter, r *http.Request)  {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = false
	session.Save(r, w)

	fmt.Fprint(w, "Logout success!")
}

func main(){
	http.HandleFunc("/home", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
