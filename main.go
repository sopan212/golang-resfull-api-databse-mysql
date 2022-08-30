package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Employe struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Position string `json:"position"`
}
type Employes []Employe

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/getEmployes", getEmployes)
	http.HandleFunc("/post-employe", post_employe)

	log.Printf("Server Run in localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func connectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_basic_sql")
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return db, nil

}
func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halaman Home"))
}
func getEmployes(w http.ResponseWriter, r *http.Request) {
	db, err := connectDb()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,name,address,position FROM employe")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Employe
	for rows.Next() {
		var each = Employe{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Address, &each.Position)
		if err != nil {
			log.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	json.NewEncoder(w).Encode(result)
}
func post_employe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := connectDb()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()
		var employe = Employe{}
		err = json.NewDecoder(r.Body).Decode(&employe)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(employe)
		_, err = db.Exec("INSERT INTO employed(name,address,position)VALUES ()")
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("insert success")
		json.NewEncoder(w).Encode(employe)
	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}
