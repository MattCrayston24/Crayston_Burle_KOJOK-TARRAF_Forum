package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

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
	SESSION_TOKEN   sql.NullString
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
	http.HandleFunc("/programme", programHandler)
	http.HandleFunc("/alimentation", alimentationHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/produits", produitsHandler)
	http.HandleFunc("/connexion", loginHandler)
	http.HandleFunc("/create_topic", createTopicHandler)

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

	// Recuperer le cookie "session_token"
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Recuperer le nom d'utilisateur à partir du token de session
	username, err := getUsernameFromSessionToken(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"username": username,
	}

	err = tmpl.Execute(w, data)
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

	cookie, err := r.Cookie("session_token")
	if err == nil {
		username, err := getUsernameFromSessionToken(cookie.Value)
		if err == nil {
			tmpl.Execute(w, map[string]string{
				"username": username,
			})
			return
		}
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

	rows, err := db.Query("SELECT topic.ID_TOPIC, topic.TITRE FROM topic JOIN definir ON topic.ID_TOPIC = definir.ID_TOPIC WHERE definir.ID_CATEGORIE = 2")
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

	cookie, err := r.Cookie("session_token")
	if err == nil {
		username, err := getUsernameFromSessionToken(cookie.Value)
		if err == nil {
			tmpl.Execute(w, struct {
				Topics   []Topic
				Username string
			}{
				Topics:   topics,
				Username: username,
			})
			return
		}
	}

	err = tmpl.Execute(w, topics)
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

	rows, err := db.Query("SELECT topic.ID_TOPIC, topic.TITRE FROM topic JOIN definir ON topic.ID_TOPIC = definir.ID_TOPIC WHERE definir.ID_CATEGORIE = 3")
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

	cookie, err := r.Cookie("session_token")
	if err == nil {
		username, err := getUsernameFromSessionToken(cookie.Value)
		if err == nil {
			tmpl.Execute(w, struct {
				Topics   []Topic
				Username string
			}{
				Topics:   topics,
				Username: username,
			})
			return
		}
	}

	err = tmpl.Execute(w, topics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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

	cookie, err := r.Cookie("session_token")
	if err == nil {
		username, err := getUsernameFromSessionToken(cookie.Value)
		if err == nil {
			tmpl.Execute(w, struct {
				Topics   []Topic
				Username string
			}{
				Topics:   topics,
				Username: username,
			})
			return
		}
	}

	err = tmpl.Execute(w, topics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createTopicHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front/create_topic.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "GET":
		rows, err := db.Query("SELECT * FROM categorie")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var categories []Categorie
		for rows.Next() {
			var c Categorie
			if err := rows.Scan(&c.ID_CATEGORIE, &c.TITRE); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			categories = append(categories, c)
		}

		cookie, err := r.Cookie("session_token")
		if err == nil {
			username, err := getUsernameFromSessionToken(cookie.Value)
			if err == nil {
				tmpl.Execute(w, struct {
					Categories []Categorie
					Username   string
				}{
					Categories: categories,
					Username:   username,
				})
				return
			}
		}

		err = tmpl.Execute(w, categories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "POST":
		// Crée un nouveau topic
		titre := r.FormValue("titre")
		categorie, err := strconv.Atoi(r.FormValue("categorie"))
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		// Ici, vous devrez remplacer "1" par l'ID de l'utilisateur connecté
		_, err = db.Exec("INSERT INTO topic (TITRE, ID_UTILISATEUR) VALUES (?, 1)", titre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Récupère l'ID du dernier topic créé
		var idTopic int
		err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&idTopic)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Associe le topic à la catégorie choisie
		_, err = db.Exec("INSERT INTO definir (ID_TOPIC, ID_CATEGORIE) VALUES (?, ?)", idTopic, categorie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirige l'utilisateur vers la page d'accueil
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "../front/connexion.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password")
		username := r.FormValue("username")
		signupEmail := r.FormValue("signup-email")
		signupPassword := r.FormValue("signup-password")
		roleID := 3

		if username != "" && signupEmail != "" && signupPassword != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupPassword), bcrypt.DefaultCost)
			if err != nil {
				fmt.Fprintln(w, "Failed to hash password")
				return
			}

			_, err = db.Exec("INSERT INTO UTILISATEUR (NOM_UTILISATEUR, ADRESSE_MAIL, MOT_DE_PASSE, ID_ROLE) VALUES (?, ?, ?, ?)", username, signupEmail, string(hashedPassword), roleID)
			if err != nil {
				fmt.Fprintln(w, "Inscription échouée")
				return
			} else {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				fmt.Fprintln(w, "Inscription réussie")
				return
			}
		}

		if email != "" && password != "" {
			var utilisateur Utilisateur

			err := db.QueryRow("SELECT * FROM UTILISATEUR WHERE ADRESSE_MAIL = ?", email).Scan(&utilisateur.ID_UTILISATEUR, &utilisateur.NOM_UTILISATEUR, &utilisateur.ADRESSE_MAIL, &utilisateur.MOT_DE_PASSE, &utilisateur.ID_ROLE, &utilisateur.SESSION_TOKEN)
			if err != nil {
				fmt.Printf("Erreur SQL : %v\n", err)
				http.Error(w, "Connexion échouée", http.StatusUnauthorized)
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(utilisateur.MOT_DE_PASSE), []byte(password))
			if err != nil {
				http.Error(w, "Connexion échouée", http.StatusUnauthorized)
				return
			}

			if utilisateur.ADRESSE_MAIL != "" {
				// Generate a new random session token
				sessionToken := generateSessionToken()

				// Set the token in the database
				_, err = db.Exec("UPDATE UTILISATEUR SET SESSION_TOKEN = ? WHERE ADRESSE_MAIL = ?", sessionToken, email)
				if err != nil {
					http.Error(w, "Failed to update session token", http.StatusInternalServerError)
					return
				}

				// Set the token as a cookie
				http.SetCookie(w, &http.Cookie{
					Name:     "session_token",
					Value:    sessionToken,
					Expires:  time.Now().Add(24 * time.Hour), // The cookie will expire in 24 hours
					HttpOnly: true,
					Secure:   true, // Use this if your site uses https
				})

				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			} else {
				http.Error(w, "Connexion échouée", http.StatusUnauthorized)
				return
			}
		}
	}
}

func generateSessionToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

/*recuperer le nom de l'utilisateur*/

func getUsernameFromSessionToken(token string) (string, error) {
	var username string
	err := db.QueryRow("SELECT NOM_UTILISATEUR FROM UTILISATEUR WHERE SESSION_TOKEN = ?", token).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			// Pas d'erreur si aucune ligne correspondante n'est trouvée
			return "", nil
		}
		// Pour toute autre erreur, la renvoyer
		return "", err
	}
	return username, nil
}
