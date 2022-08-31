package route

import (
	"encoding/json"
	"fmt"
	"golang_basic_rest_api/config"
	"golang_basic_rest_api/model"
	"log"
	"net/http"
)

func GtDepartement(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConnectDb()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer db.Close()
	row, err := db.Query("SELECT id,name,code FROM departments")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer row.Close()
	var ress []model.Departement
	for row.Next() {
		var item = model.Departement{}
		err := row.Scan(&item.ID, &item.Name, &item.Code)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		ress = append(ress, item)
	}
	json.NewEncoder(w).Encode(ress)
}
func PostDepartement(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := config.ConnectDb()
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		defer db.Close()
		var department = model.Departement{}
		err = json.NewDecoder(r.Body).Decode(&department)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		log.Print(department)

		_, err = db.Exec("INSERT INTO departments(id,name,code) VALUES (?,?,?)",
			department.ID, department.Name, department.Code)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		log.Println("Insert Success")
		json.NewEncoder(w).Encode(department)

	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}
func PutDepartement(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		db, err := config.ConnectDb()
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		defer db.Close()
		var department = model.Departement{}
		err = json.NewDecoder(r.Body).Decode(&department)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		log.Print(department)

		_, err = db.Exec("UPDATE departments SET name = ?,code = ? WHERE id = ?",
			department.ID, department.Name, department.Code)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		log.Println("Update Success")
		json.NewEncoder(w).Encode(department)

	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

func DeleteDepartement(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		db, err := config.ConnectDb()
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		defer db.Close()
		var department model.Departement
		err = json.NewDecoder(r.Body).Decode(&department)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		log.Print(department)

		_, err = db.Exec("DELETE FROM departments WHERE id = ?", department.ID)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		log.Println("Dellete Success")
		json.NewEncoder(w).Encode(department)

	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}
