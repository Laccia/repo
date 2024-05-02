package base_test

import (
	"context"
	"repo/cmd/base"
	"repo/cmd/models"
	"testing"
)

func TestPg(t *testing.T) {
	ctx := context.Background()
	t.Run("1", func(t *testing.T) {
		pgtest, _ := base.NewPG(ctx)
		pgtest.NewUser(ctx, models.CreateUser{Account: "test", Name: "test", Pass: "test", Email: "test"})
	})
}
