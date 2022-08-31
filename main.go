package main

import (
	"golang_basic_rest_api/route"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", route.GetHome)
	http.HandleFunc("/getEmployes", route.GetEmployes)
	http.HandleFunc("/post-employe", route.Post_employes)
	http.HandleFunc("/putEmploye", route.PutEmploye)
	http.HandleFunc("/deleteEmploye", route.DeleteEmploye)

	http.HandleFunc("/getDepartement", route.GtDepartement)
	http.HandleFunc("/postDepartement", route.PostDepartement)
	http.HandleFunc("/putDepartement", route.PutDepartement)
	http.HandleFunc("/deleteDepartement", route.DeleteDepartement)
	log.Printf("Server Run in localhost:3000")
	http.ListenAndServe(":3000", nil)
}
