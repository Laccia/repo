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
	Name  string  `json:"name"`
	Price float64 `Json:"price"`
}

type CreateUser struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Pass    string `json:"pass"`
	Email   string `json:"email"`
}

type Login struct {
	Account string `json:"account"`
	Pass    string `json:"pass"`
}

type CreateItem struct {
	Name  string  `json:"name"`
	Price float64 `Json:"price"`
}

type DeleteItem struct {
	Id uint
}
