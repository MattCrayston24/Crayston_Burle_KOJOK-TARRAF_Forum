package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Utilisateur struct {
	ID              int
	NomUtilisateur  string
	AdresseMail     string
	MotDePasse      string
	RoleID          int
	DateInscription string
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/forum")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	assets := http.FileServer(http.Dir("../front/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", assets))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/connexion", loginHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../front/index.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "../front/connexion.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password")

		var utilisateur Utilisateur
		err := db.QueryRow("SELECT * FROM utilisateur WHERE adresse_mail = ? AND mot_de_passe = ?", email, password).Scan(&utilisateur.ID, &utilisateur.NomUtilisateur, &utilisateur.AdresseMail, &utilisateur.MotDePasse, &utilisateur.RoleID, &utilisateur.DateInscription)
		if err != nil {
			fmt.Fprintln(w, "Connexion échouée")
		}

		if utilisateur.AdresseMail != "" {
			fmt.Fprintln(w, "Connexion réussie")
		} else {
			fmt.Fprintln(w, "Connexion échouée")
		}

		username := r.FormValue("username")
		signupEmail := r.FormValue("signup-email")
		signupPassword := r.FormValue("signup-password")
		roleID := 2

		if username != "" && signupEmail != "" && signupPassword != "" {
			_, err := db.Exec("INSERT INTO utilisateur (nom_utilisateur, adresse_mail, mot_de_passe, role_id) VALUES (?, ?, ?, ?)", username, signupEmail, signupPassword, roleID)
			if err != nil {
				fmt.Fprintln(w, "Inscription échouée")
			} else {
				fmt.Fprintln(w, "Inscription réussie")
			}
		}
	}
}
