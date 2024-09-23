package test

import (
	"testing"

	"github.com/wingfeng/idx/repo"
)

func TestDBUserRepository_GetUser(t *testing.T) {
	db := initTestDb()
	repo := repo.NewUserRepository(db)

	user, err := repo.GetUser("7a45cb54-b0ff-4ecd-95b9-074d33aaac1e")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
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
