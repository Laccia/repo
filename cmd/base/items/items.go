package items

import "repo/cmd/models"

type ItemsBase struct {
	Items []models.Items
}

func New() *ItemsBase {
	tmp := ItemsBase{[]models.Items{
		{
			Id:    1,
			Name:  "test",
			Price: 0.01,
		},
	}}
	return &tmp
}
