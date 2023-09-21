package test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"github.com/patrickmn/go-cache"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/store"
	"github.com/wingfeng/idx/utils"
	"gorm.io/gorm"
)

func TestClientStore_GetByID(t *testing.T) {
	initTest()

	gocache := cache.New(5*time.Minute, 60*time.Second)

	cs := &store.ClientStore{
		DB:    db,
		Cache: gocache,
	}

	ctx := context.Background()
	wg := sync.WaitGroup{}
	//
	// c := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			got, err := cs.GetByID(ctx, "local_test")

			if err != nil {
				t.Errorf("ClientStore.GetByID() error = %v, wantErr ", err)
				return
			}
			t.Logf("Client:%v", got)
			// c = c + j
			// t.Logf("Counter:%d", c)
			// if c == 4950 {
			// 	wg.Done()
			// }
			wg.Done()
		}(i)

	}
	wg.Wait()
}

func TestClientStore_VerifySecret(t *testing.T) {
	initTest()
	gocache := cache.New(5*time.Minute, 60*time.Second)

	cs := &store.ClientStore{
		DB:    db,
		Cache: gocache,
	}
	cs.ValidateSecret("hybrid_client", "hybrid_secret")
}

var db *gorm.DB

func initTest() {
	//初始化DB
	//db = utils.GetDB("mysql", "root:kXbXt2nLrL@tcp(localhost:3306)/idx?&parseTime=true")
	db = utils.GetDB("pgx", "host=localhost user=root password=pass@word1 dbname=idx port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	//
	models.Sync2Db(db)
}
func TestSeedData(t *testing.T) {
	//	node, err := snowflake.NewNode(1)
	initTest()

	ou := &models.OrganizationUnit{}
	ou.ID = "1328680589330485248"
	ou.Name = "翼火工作室"
	ou.DisplayName = "翼火工作室"

	err := db.Save(ou).Error
	if err != nil {
		panic(err)
	}

	user := &models.User{}
	user.ID = "7a45cb54-b0ff-4ecd-95b9-074d33aaac1e"
	user.Account = "admin"
	user.DisplayName = "管理员"
	user.Email = "admin@fire.loc"
	user.OUID = ou.ID
	user.OU = ou.DisplayName

	user.PasswordHash = utils.GenHashedPWD("fire@123")

	err = db.Save(user).Error
	if err != nil {
		panic(err)
	}
	role := &models.Role{}

	role.ID = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43b"
	role.Name = "SystemAdmin"
	addRole(role)
	addUserRole(user.ID, ou.ID, role.ID)
	role = &models.Role{}

	role.ID = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43c"
	role.Name = "科室主任"
	addRole(role)
	addUserRole(user.ID, ou.ID, role.ID)
	addClient("implicit_client", "implicit_secret", "implicit")
	addClient("hybrid_client", "hybrid_secret", "hybrid")
	addClient("code_client", "code_secret", "authorization_code")
	addClient("password_client", "password_secret", "password")
	addClient("local_test", "local_secret", "authorization_code")

}
func TestInsertHybrid(t *testing.T) {
	initTest()
	addClient("oidc-client-implicit.test", "secret", "hybrid")

}
func TestInsertAuthCode(t *testing.T) {
	initTest()
	addClient("local_test", "local_secret", "authorization_code")

}
func TestInsertPasswordClient(t *testing.T) {
	initTest()
	addClient("password_client", "password_secret", "password")

}
func addClient(clientID, secret, grantType string) {
	requireSecret := len(secret) > 0
	client := &models.Client{

		ClientCode:                       clientID,
		Enabled:                          true,
		ProtocolType:                     "oidc",
		RequireClientSecret:              requireSecret,
		ClientName:                       "Client",
		Domains:                          "http://localhost:9000,http://localhost:9001",
		GrantTypes:                       grantType,
		Scopes:                           "openid email profile roles",
		RequireConsent:                   true,
		AllowRememberConsent:             true,
		AlwaysIncludeUserClaimsInIDToken: false,
		AllowAccessTokensViaBrowser:      true,
		BackChannelLogoutSessionRequired: true,
		IDentityTokenLifetime:            300,
		AccessTokenLifetime:              3600,
		AbsoluteRefreshTokenLifetime:     2592000,
		SlidingRefreshTokenLifetime:      2592000,
		AuthorizationCodeLifetime:        300,
		RefreshTokenUsage:                1,
		RefreshTokenExpiration:           1,
		ClientClaimsPrefix:               "client_",
		DeviceCodeLifetime:               300,

		EnableLocalLogin: true,
		//UserSsoLifetime: , can be zero
	}

	db.Save(client).Where("ClientId=?", clientID)

	if client.RequireClientSecret {
		addClientScecret(secret, client.ID)
	}
}
func addClientScecret(secret string, clientid int) {
	sc := &models.ClientSecrets{
		Type:     "SHA256",
		ClientID: clientid,
	}
	sc.Value = utils.HashString(secret)
	sc.Expiration = time.Now().AddDate(1, 0, 0)

	err := db.Save(sc).Error
	if err != nil {
		panic(err)
	}
}

func addUserRole(uid, ouid, rid string) {

	ur := &models.UserRoles{
		RoleID: rid,
		UserID: uid,
		OUID:   ouid,
	}
	//联合主键的直接用engine来处理
	err := db.Save(ur).Error
	if err != nil {
		panic(err)
	}
}
func addRole(role *models.Role) {

	err := db.Save(role).Error
	if err != nil {
		panic(err)
	}
}

func TestValidateSecret(t *testing.T) {
	initTest()
	cs := &store.ClientStore{
		DB: db,
	}
	err := cs.ValidateSecret("local_test", "local_secret")
	assert.Equal(t, nil, err)
}
