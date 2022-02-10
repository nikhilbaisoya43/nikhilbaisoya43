package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID            int
	Name_products string
	Image_url     string
	Description   string
	Price         int
	Name_user     string
	Email_user    string
	Category      string
}

// Declaring the mysql path
var db_url = "root:rootroot@tcp(localhost:3306)/mydb?parseTime=true"
var err error
var DB *gorm.DB

func main() {
	Conn_Est()
	Handler_Routing()
}

/*
connection established in mysql using
*/
func Conn_Est() {
	DB, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	DB.AutoMigrate(&Product{})
}

/*
api set to create delete and view all product list
*/
func Handler_Routing() {
	r := mux.NewRouter()
	r.HandleFunc("/allproduct", AllProduct).Methods("GET")
	r.HandleFunc("/Cproduct", CreateProduct).Methods("POST")
	r.HandleFunc("/Dproduct", DeleteProductbyId).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

/*
established handlers
*/
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newproduct Product
	json.NewDecoder(r.Body).Decode(&newproduct)
	DB.Create(&newproduct)
	json.NewEncoder(w).Encode(newproduct)
}

func AllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var product []Product
	DB.Find(&product)
	json.NewEncoder(w).Encode(product)
}
func DeleteProductbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var delproduct Product
	DB.Delete(&delproduct, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("product is deleted")
}
