package users

import "repo/cmd/models"

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
