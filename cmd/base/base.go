package base

import (
	"repo/cmd/base/items"
	"repo/cmd/base/users"
)

type Memory struct {
	*items.ItemsBase
	*users.UsersBase
}

func New() *Memory {
	return &Memory{
		items.New(),
		users.New(),
	}
}
