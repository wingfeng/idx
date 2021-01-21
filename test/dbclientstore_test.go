package test

import (
	"context"
	"database/sql"
	"reflect"
	"strings"
	"testing"
	"time"

	log "github.com/cihub/seelog"
	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/idx/models"
	idxmodels "github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/store"
	"github.com/wingfeng/idx/utils"
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

func initTest() {
	//初始化DB
	db = GetDB("mysql", "root:123456@tcp(localhost:3306)/sso?&parseTime=true")
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
	addClient("vue_client", "vue_secret", "implicit")
}
func addClient(clientID, secret, grantType string) {
	client := &models.Client{

		ClientID:                         clientID,
		Enabled:                          true,
		ProtocolType:                     "oidc",
		RequireClientSecret:              false,
		ClientName:                       "Client",
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

	db.Save(client).Where("clientId=?", clientID)
	cg := &models.ClientGrantTypes{
		ClientID:  client.ID,
		GrantType: grantType,
	}

	err := db.Save(cg).Error
	if err != nil {
		panic(err)
	}

	addRedirectURI("http://localhost:9094/oauth2", client.ID)
	addRedirectURI("http://localhost:9000/", client.ID)
	addClientScope("openid", client.ID)
	addClientScope("profile", client.ID)
	addClientScope("email", client.ID)
	addClientScope("roles", client.ID)
	addClientScecret(secret, client.ID)
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
func addRedirectURI(uri string, clientid int) {

	redUris := &models.ClientRedirectURIs{

		RedirectURI: uri,
		ClientID:    clientid,
	}

	err := db.Save(redUris).Error
	if err != nil {

		panic(err)
	}
}
func addClientScope(scope string, clientid int) {
	sc := &models.ClientScopes{

		Scope:    scope,
		ClientID: clientid,
	}

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

func TestGetClientScopes(t *testing.T) {
	initTest()
	cs := &store.ClientStore{
		DB: db,
	}
	scopes := cs.GetClientScopes("local_test")
	assert.Equal(t, len(scopes), 4)
}
func TestValidateSecret(t *testing.T) {
	initTest()
	cs := &store.ClientStore{
		DB: db,
	}
	err := cs.ValidateSecret("local_test", "local_secret")
	assert.Equal(t, nil, err)
}
