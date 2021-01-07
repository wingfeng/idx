package store

import (
	"context"
	"database/sql"
	"reflect"
	"strings"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/wingfeng/idx/models"
	idxmodels "github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/store"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestClientStore_GetByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idxmodels.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &store.ClientStore{
				DB: tt.fields.DB,
			}
			got, err := cs.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientStore.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientStore.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
func GetDB(driver string, connection string) *gorm.DB {
	if strings.EqualFold(driver, "") {
		driver = "mysql"
	}

	var err error
	sqlDB, err := sql.Open(driver, connection)
	x, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if nil != err {
		log.Error("init" + err.Error())
	}

	return x
}

var db *gorm.DB

func TestSeedData(t *testing.T) {
	//	node, err := snowflake.NewNode(1)

	//初始化DB
	db = GetDB("mysql", "root:123456@tcp(localhost:3306)/sso?&parseTime=true")
	models.Sync2Db(db)

	client := &models.Client{
		ID:                               2,
		ClientID:                         "222222",
		Enabled:                          true,
		ProtocolType:                     "oidc",
		RequireClientSecret:              false,
		ClientName:                       "go Client",
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

	db.Save(client)
	cg := &models.ClientGrantTypes{
		ID:        1,
		ClientID:  1,
		GrantType: "implicit",
	}

	err := db.Save(cg).Error
	if err != nil {
		panic(err)
	}

	addRedirectURI(1, "http://localhost:9094/oauth2", client.ID)
	addRedirectURI(2, "http://localhost:9000", client.ID)
	addClientScope(1, "openid", client.ID)
	addClientScope(2, "profile", client.ID)
	addClientScope(3, "roles", client.ID)
}
func addRedirectURI(id int, uri string, clientid int) {

	redUris := &models.ClientRedirectURIs{
		ID:          id,
		RedirectURI: uri,
		ClientID:    clientid,
	}

	err := db.Save(redUris).Error
	if err != nil {

		panic(err)
	}
}
func addClientScope(id int, scope string, clientid int) {
	sc := &models.ClientScopes{
		ID:       id,
		Scope:    scope,
		ClientID: clientid,
	}

	err := db.Save(sc).Error
	if err != nil {
		panic(err)
	}
}
