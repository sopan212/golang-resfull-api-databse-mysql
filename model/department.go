package model

type Departement struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
type Department Departement
