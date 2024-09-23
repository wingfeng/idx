package test

import (
	"testing"

	"github.com/wingfeng/idx/repo"
)

func TestClientRepository_GetClient(t *testing.T) {
	db := initTestDb()
	repo := repo.NewClientRepository(db)

	client, err := repo.GetClientByClientID("hybrid_client")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(client)
	t.Log("Secrets:", client.GetSecret())
	t.Log("Scopes:", client.GetScopes())
	t.Log("Origins:", client.GetWebOrigins())
	t.Log("LogoutRedirectUris:", client.GetPostLogoutUris())
	t.Log("RedirectUris:", client.GetRedirectUris())
}
