package users

import (
	"repo/cmd/models"
)

type UsersBase struct {
	Users []models.Users
}

func New() *UsersBase {
	tmp := UsersBase{[]models.Users{
		{
			Account: "doe",
			Name:    "Test",
			Pass:    "1234",
			Email:   "test@test.com",
			Group:   1,
		},
	}}
	return &tmp
}

func (v *UsersBase) Validate(Account, Pass string) bool {

	for _, tmp := range v.Users {
		if tmp.Account == Account {
			if tmp.Pass == Pass {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func (g *UsersBase) ValidateRule(Account string) bool {
	for _, tmp := range g.Users {
		if tmp.Account == Account {
			if tmp.Group == 1 {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func (m *UsersBase) CreateUser(create *models.CreateUser) error {
	m.Users = append(m.Users, models.Users{
		Account: create.Account,
		Name:    create.Name,
		Pass:    create.Pass,
		Email:   create.Email,
		Group:   0,
	})
	return nil
}
