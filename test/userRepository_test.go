package test

import (
	"testing"

	"github.com/wingfeng/idx/repo"
)

func TestDBUserRepository_GetUser(t *testing.T) {
	db := initTestDb()
	repo := repo.NewUserRepository(db)

	user, err := repo.GetUser("1838872840128958464")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
	t.Logf("UserID %s", user.GetId())
}
func TestDBUserRepository_GetUserByUserName(t *testing.T) {
	db := initTestDb()
	repo := repo.NewUserRepository(db)

	user, err := repo.GetUserByName("admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}
