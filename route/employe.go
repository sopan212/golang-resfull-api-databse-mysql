package route

import (
	"encoding/json"
	"fmt"
	"golang_basic_rest_api/config"
	"golang_basic_rest_api/model"
	"log"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halaman Home"))
}
func GetEmployes(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConnectDb()
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

	var result []model.Employe
	for rows.Next() {
		var each = model.Employe{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Address, &each.Position)
		if err != nil {
			log.Println(err.Error())
			return
		}
		result = append(result, each)
	}
	json.NewEncoder(w).Encode(result)
}
func Post_employes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := config.ConnectDb()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()
		var employe = model.Employe{}
		err = json.NewDecoder(r.Body).Decode(&employe)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(employe)
		_, err = db.Exec("INSERT INTO employe(name,address,position)VALUES (?,?,?)",
			employe.Name, employe.Address, employe.Position)
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
func PutEmploye(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		db, err := config.ConnectDb()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()

		var employee model.Employe

		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(employee)

		_, err = db.Exec("UPDATE employe SET name = ?,address = ?,position = ? WHERE id = ?",
			employee.Name, employee.Address, employee.Position, employee.ID)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		fmt.Print("update success")
		json.NewEncoder(w).Encode(employee)
	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

func DeleteEmploye(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		db, err := config.ConnectDb()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer db.Close()

		var employee model.Employe

		err = json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(employee)

		_, err = db.Exec("DELETE employe  WHERE id = ?", employee.ID)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		fmt.Print("Delete Success")
		json.NewEncoder(w).Encode(employee)
	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}
