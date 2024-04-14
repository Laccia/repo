package models

type Users struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Pass    string `json:"-"`
	Email   string `json:"email"`
	Group   int    `json:"-"`
}

type Items struct {
	Id    uint
	Name  string
	Price float64
}
