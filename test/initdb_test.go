package test

import (
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
	constants "github.com/wingfeng/idx-oauth2/const"
	"github.com/wingfeng/idx-oauth2/utils"
	"github.com/wingfeng/idx/models"
	idxutils "github.com/wingfeng/idx/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var db *gorm.DB

// initDB Initializes the database connection and synchronizes the database schema.
//
// No parameters.
// No return values.
func initDB() {
	//初始化DB
	//	db = idxutils.GetDB("mysql", "root:password1@tcp(localhost:3306)/idx?&parseTime=true")
	db = idxutils.GetDB("pgx", "host=localhost user=root password=pass@word1 dbname=idx port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	//
	models.Sync2Db(db)
}

// TestSeedData tests the seeding of initial data into the database.
//
// Parameter t is a pointer to testing.T, which is used to report test failures.
// No return values.
func TestSeedData(t *testing.T) {

	initDB()

	ou := &models.OrganizationUnit{}
	ou.Id = 1328680589330485248
	ou.Name = "集团"
	ou.DisplayName = "XXX集团"
	ou.Path = "0"
	err := db.Save(ou).Error
	if err != nil {
		panic(err)
	}

	ou = &models.OrganizationUnit{}
	ou.Id = 1328680589330485249
	ou.Name = "子公司"
	ou.DisplayName = "XXX集团子公司"
	ou.ParentId = 1328680589330485248
	//  sql.NullInt64{
	// 	Valid: true,
	// 	Int64: 1328680589330485248,
	// }
	ou.Path = "0.1"

	err = db.Save(ou).Error
	if err != nil {
		panic(err)
	}

	user := &models.User{}
	user.Id = 1838872840128958464
	user.Account = "admin"
	user.DisplayName = "管理员"
	user.Email = "admin@idx.local"
	user.EmailConfirmed = true
	user.OUId = ou.Id
	user.OU = ou.DisplayName
	user.Claims = datatypes.JSON([]byte(`{"alias":"db_admin"}`))
	user.PasswordHash, _ = utils.HashPassword("password1")

	err = db.Save(user).Error
	if err != nil {
		panic(err)
	}
	role := &models.Role{}

	role.Id = 1838872840128958465
	role.Name = "admin"
	addRole(role)
	addUserRole(user.Id, role.Id)
	role = &models.Role{}

	role.Id = 1838872840128958466
	role.Name = "科室主任"
	addRole(role)
	addUserRole(user.Id, role.Id)
	addClient("implicit_client", "secret", "implicit", t)
	addClient("hybrid_client", "secret", "authorization_code implicit "+string(constants.DeviceCode)+" password client_credential", t)
	addClient("code_client", "secret", "authorization_code", t)
	addClient("password_client", "secret", "password", t)
	addClient("local_test", "secret", "authorization_code", t)
	addClient("client_credentials_client", "secret", "client_credentials", t)
	addClient("device_code_client", "secret", string(constants.DeviceCode), t)

}

// addClient adds a new client to the database with the given client ID, secret, grant type, and testing context.
//
// Parameters:
// - clientId: the ID of the client to be added (string)
// - secret: the secret of the client to be added (string)
// - grantType: the grant type of the client to be added (string)
// - t: the testing context (testing.T)
//
// Return:
// - None
func addClient(clientId, secret, grantType string, t *testing.T) {
	requireSecret := len(secret) > 0

	client := &models.Client{

		ClientId:   clientId,
		Enabled:    true,
		ClientName: "Client",

		GrantTypes: grantType,

		Scopes:         "openid email profile roles",
		RequireConsent: true,
		RequireSecret:  requireSecret,
		//UserSsoLifetime: , can be zero
	}

	var result *gorm.DB
	var newClient models.Client
	if db.Table("clients").Where("client_id=?", clientId).First(&newClient).RowsAffected > 0 {
		result = db.Table("clients").Where("client_id=?", clientId).Updates(client)

	} else {
		result = db.Table("clients").Where("client_id=?", clientId).Save(client)
	}
	if result.Error != nil {
		t.Logf("insert client error: %v", result.Error)
		panic(result.Error)
	}
	if requireSecret {
		addClientScecret(secret, newClient.Id)
	}

	addClientOrigin("*", newClient.Id)
}
func addClientOrigin(origin string, clientid int64) {

	db.Unscoped().Delete(&models.ClientCorsOrigins{}, "client_id = ?", clientid)

	co := &models.ClientCorsOrigins{
		ClientId: clientid,
		Origin:   origin,
	}

	err := db.Save(co).Error
	if err != nil {
		panic(err)
	}
}
func addClientScecret(secret string, clientid int64) {
	db.Unscoped().Delete(&models.ClientSecrets{}, "client_id = ?", clientid)
	sc := &models.ClientSecrets{

		ClientId: clientid,
	}
	sc.Value, _ = utils.HashPassword(secret)
	sc.Expiration = time.Now().AddDate(1, 0, 0)

	err := db.Save(sc).Error
	if err != nil {
		panic(err)
	}
}

func addUserRole(uid, rid snowflake.ID) {

	ur := &models.UserRoles{
		RoleId: rid,
		UserId: uid,
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
