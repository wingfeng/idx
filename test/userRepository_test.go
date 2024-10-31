package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wingfeng/idx/repo"
)

func TestDBUserRepository_GetUser(t *testing.T) {
	db := initTestDb()
	repo := repo.NewUserRepository()
	repo.DB = db

	user, err := repo.GetUser("1838872840128958464")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
	t.Logf("UserID %s", user.GetId())
}
func TestDBUserRepository_GetUserByUserName(t *testing.T) {
	db := initTestDb()
	repo := repo.NewUserRepository()
	repo.DB = db

	user, err := repo.GetUserByName("admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestChangePassword(t *testing.T) {
	db := initTestDb()
	repo := repo.NewUserRepository()
	repo.DB = db

	err := repo.ChangePassword("admin", "123456", "654321")
	assert.Error(t, err)
	resetedPwd, err := repo.ResetPassword("admin")
	assert.NoError(t, err)
	err = repo.ChangePassword("admin", resetedPwd, "password1")
	assert.NoError(t, err)
}
func TestResetPassword(t *testing.T) {
	db := initTestDb()
	repo := repo.NewUserRepository()
	repo.DB = db
	newPwd, err := repo.ResetPassword("admin")
	assert.NoError(t, err)
	assert.NotEmpty(t, newPwd)
	assert.Equal(t, 8, len(newPwd))
	t.Log(newPwd)
}
