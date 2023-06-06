package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/auth", authHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "connexion.html") // Assurez-vous que le chemin vers votre fichier HTML est correct
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("mot_de_passe")

		if r.FormValue("action") == "login" {
			var user User
			row := db.QueryRow("SELECT * FROM users WHERE email = ? AND password = ?", email, password)
			err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
			if err != nil {
				http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther) // redirige vers la page d'accueil après une connexion réussie
		} else if r.FormValue("action") == "signup" {
			pseudo := r.FormValue("pseudo")

			stmt, err := db.Prepare("INSERT INTO users(name, email, password) VALUES(?, ?, ?)")
			if err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}
			defer stmt.Close()

			_, err = stmt.Exec(pseudo, email, password)
			if err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/auth", http.StatusSeeOther) // redirige vers la page de connexion après une inscription réussie
		}
	}
}
