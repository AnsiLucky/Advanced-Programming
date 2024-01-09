package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path"
	"text/template"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var tpl *template.Template
var db *gorm.DB

func init() {
	tpl = template.Must(template.ParseGlob(path.Join("web", "*.html")))
}

func main() {
	// Initialize Gorm with PostgreSQL
	dsn := "user=postgres dbname=assik2advanced port=7777 password=15691804 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto Migrate will create the users table based on the User struct
	db.AutoMigrate(&User{})

	// Create
	user := User{Name: "John Doe", Email: "john@example.com", Age: 30}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Read
	var fetchedUser User
	db.First(&fetchedUser, 1)
	log.Println(fetchedUser)

	// Update
	db.Model(&fetchedUser).Update("Name", "Doe John")

	// Delete
	db.Delete(&fetchedUser, 1)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/createUser", createUser)

	http.ListenAndServe("localhost:8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Parse incoming JSON data
	var user User
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the user to the database (using Gorm)
	result := db.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request parameters
	params := mux.Vars(r)
	userID := params["id"]

	// Fetch the user from the database by ID
	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	// Send the user details as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
