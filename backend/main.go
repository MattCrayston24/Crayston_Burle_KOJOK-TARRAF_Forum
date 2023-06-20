package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var (
	db    *sql.DB
	store = sessions.NewCookieStore([]byte("secret-key"))
)

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
	session, _ := store.Get(r, "session")
	user := getUser(session.Values["user"])

	tmpl, err := template.ParseFiles("../front/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("../front/connexion.html")
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password")

		var utilisateur Utilisateur
		err := db.QueryRow("SELECT * FROM utilisateur WHERE adresse_mail = ? AND mot_de_passe = ?", email, password).Scan(&utilisateur.ID, &utilisateur.NomUtilisateur, &utilisateur.AdresseMail, &utilisateur.MotDePasse, &utilisateur.RoleID, &utilisateur.DateInscription)
		if err != nil {
			http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		}

		if utilisateur.AdresseMail != "" {
			session.Values["user"] = utilisateur.NomUtilisateur
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		}

		username := r.FormValue("username")
		signupEmail := r.FormValue("signup-email")
		signupPassword := r.FormValue("signup-password")
		roleID := 2

		if username != "" && signupEmail != "" && signupPassword != "" {
			_, err := db.Exec("INSERT INTO utilisateur (nom_utilisateur, adresse_mail, mot_de_passe, role_id) VALUES (?, ?, ?, ?)", username, signupEmail, signupPassword, roleID)
			if err != nil {
				http.Redirect(w, r, "/connexion", http.StatusSeeOther)
			} else {
				session.Values["user"] = username
				session.Save(r, w)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}
	}
}

func getUser(username interface{}) *Utilisateur {
	if username == nil {
		return &Utilisateur{}
	}
	var utilisateur Utilisateur
	db.QueryRow("SELECT * FROM utilisateur WHERE nom_utilisateur = ?", username).Scan(&utilisateur.ID, &utilisateur.NomUtilisateur, &utilisateur.AdresseMail, &utilisateur.MotDePasse, &utilisateur.RoleID, &utilisateur.DateInscription)

	return &utilisateur
}
