package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//Representation des tables en struct

type Role struct {
	ID_ROLE int
	GRADE   string
}

type Categorie struct {
	ID_CATEGORIE int
	TITRE        string
}

type Utilisateur struct {
	ID_UTILISATEUR  int
	NOM_UTILISATEUR string
	ADRESSE_MAIL    string
	MOT_DE_PASSE    string
	ID_ROLE         int
}

type Topic struct {
	ID_TOPIC       int
	TITRE          string
	ID_UTILISATEUR int
}

type Message struct {
	ID_MESSAGE     int
	CONTENU        string
	ID_MESSAGE_1   int
	ID_TOPIC       int
	ID_UTILISATEUR int
}

type Definir struct {
	ID_TOPIC     int
	ID_CATEGORIE int
}

//fonction pour toutes les pages

func main() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/newforum")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	assets := http.FileServer(http.Dir("../front/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", assets))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/connexion", loginHandler)
	http.HandleFunc("/programme", programHandler)
	http.HandleFunc("/alimentation", alimentationHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/produits", produitsHandler)

	fmt.Println("(http://localhost:8080/) - Server is running on port 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func alimentationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front/alimentation.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front/contact.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func produitsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front/produits.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*func programHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front/programme.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}*/

func programHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front/programme.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT topic.ID_TOPIC, topic.TITRE FROM topic JOIN definir ON topic.ID_TOPIC = definir.ID_TOPIC WHERE definir.ID_CATEGORIE = 1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var topics []Topic
	for rows.Next() {
		var t Topic
		if err := rows.Scan(&t.ID_TOPIC, &t.TITRE); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		topics = append(topics, t)
	}

	err = tmpl.Execute(w, topics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "../front/connexion.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password") // Pensez à hacher ce mot de passe avant de le comparer à celui dans la base de données

		var utilisateur Utilisateur
		err := db.QueryRow("SELECT * FROM UTILISATEUR WHERE ADRESSE_MAIL = ? AND MOT_DE_PASSE = ?", email, password).Scan(&utilisateur.ID_UTILISATEUR, &utilisateur.NOM_UTILISATEUR, &utilisateur.ADRESSE_MAIL, &utilisateur.MOT_DE_PASSE, &utilisateur.ID_ROLE)
		if err != nil {
			fmt.Fprintln(w, "Connexion échouée")
		}

		if utilisateur.ADRESSE_MAIL != "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Fprintln(w, "Connexion échouée")
		}

		username := r.FormValue("username")
		signupEmail := r.FormValue("signup-email")
		signupPassword := r.FormValue("signup-password") // Vous devriez hacher ce mot de passe avant de le stocker
		roleID := 3                                      // Selon votre exemple d'insertion, le rôle d'utilisateur normal a un ID_ROLE de 3

		if username != "" && signupEmail != "" && signupPassword != "" {
			_, err := db.Exec("INSERT INTO UTILISATEUR (NOM_UTILISATEUR, ADRESSE_MAIL, MOT_DE_PASSE, ID_ROLE) VALUES (?, ?, ?, ?)", username, signupEmail, signupPassword, roleID)
			if err != nil {
				fmt.Fprintln(w, "Inscription échouée")
			} else {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}
	}
}
