package model

type Employe struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Position string `json:"position"`
}
type Employes []Employe
