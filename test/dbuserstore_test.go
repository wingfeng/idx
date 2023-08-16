package test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/idx/store"
)

func TestGetUser(t *testing.T) {
	initTest()
	us := store.NewDbUserStore(db)
	user, err := us.GetUserByAccount("admin")
	assert.Equal(t, err, nil, "Error should be null")
	t.Logf("User %s, Password: %s", user.Account, user.PasswordHash)
}
