package items

import (
	"errors"
	"repo/cmd/models"
)

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

func (m *ItemsBase) CreateItem(create *models.CreateItem) error {
	m.Items = append(m.Items, models.Items{
		Id:    m.Items[len(m.Items)-1].Id + 1,
		Name:  create.Name,
		Price: create.Price,
	})
	return nil
}

func (d *ItemsBase) DeleteItem(delete *models.DeleteItem) error {
	for i, tmp := range d.Items {
		if tmp.Id == delete.Id {
			d.Items = append(d.Items[:i], d.Items[i+1:]...)
			return nil
		}
	}
	return errors.New("no found")
}
