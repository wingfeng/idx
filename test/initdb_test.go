package test

import (
	"testing"
	"time"

	"github.com/wingfeng/idx-oauth2/utils"
	"github.com/wingfeng/idx/models"
	idxutils "github.com/wingfeng/idx/utils"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	//初始化DB
	//db = utils.GetDB("mysql", "root:kXbXt2nLrL@tcp(localhost:3306)/idx?&parseTime=true")
	db = idxutils.GetDB("pgx", "host=localhost user=root password=pass@word1 dbname=idx port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	//
	models.Sync2Db(db)
}

func TestSeedData(t *testing.T) {
	//	node, err := snowflake.NewNode(1)
	initDB()

	ou := &models.OrganizationUnit{}
	ou.Id = "1328680589330485248"
	ou.Name = "集团"
	ou.DisplayName = "XXX集团"
	ou.Path = "0"
	err := db.Save(ou).Error
	if err != nil {
		panic(err)
	}

	ou = &models.OrganizationUnit{}
	ou.Id = "1328680589330485249"
	ou.Name = "子公司"
	ou.DisplayName = "XXX集团子公司"
	ou.Parent = null.NewString("1328680589330485248", true)
	ou.Path = "0.1"
	err = db.Save(ou).Error
	if err != nil {
		panic(err)
	}

	user := &models.User{}
	user.Id = "7a45cb54-b0ff-4ecd-95b9-074d33aaac1e"
	user.Account = "admin"
	user.DisplayName = "管理员"
	user.Email = "admin@idx.local"
	user.EmailConfirmed = true
	user.OUId = ou.Id
	user.OU = ou.DisplayName

	user.PasswordHash, _ = utils.HashPassword("password1")

	err = db.Save(user).Error
	if err != nil {
		panic(err)
	}
	role := &models.Role{}

	role.Id = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43b"
	role.Name = "SystemAdmin"
	addRole(role)
	addUserRole(user.Id, ou.Id, role.Id)
	role = &models.Role{}

	role.Id = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43c"
	role.Name = "科室主任"
	addRole(role)
	addUserRole(user.Id, ou.Id, role.Id)
	addClient("implicit_client", "implicit_secret", "implicit", t)
	addClient("hybrid_client", "hybrid_secret", "authorization_code implicit device_code password client_credential", t)
	addClient("code_client", "code_secret", "authorization_code", t)
	addClient("password_client", "password_secret", "password", t)
	addClient("local_test", "local_secret", "authorization_code", t)

}

func addClient(clientId, secret, grantType string, t *testing.T) {
	//requireSecret := len(secret) > 0
	pwdHash, _ := utils.HashPassword(secret)
	client := &models.Client{

		ClientId: clientId,

		ClientName: "Client",

		GrantTypes: grantType,

		Scopes:         "openid email profile roles",
		RequireConsent: true,

		//UserSsoLifetime: , can be zero
	}

	var result *gorm.DB
	if db.Table("clients").Where("client_id=?", clientId).First(&models.Client{}).RowsAffected > 0 {
		result = db.Table("clients").Where("client_id=?", clientId).Updates(client)

	} else {
		result = db.Table("clients").Where("client_id=?", clientId).Save(client)
	}
	if result.Error != nil {
		t.Logf("insert client error: %v", result.Error)
		panic(result.Error)
	}
	addClientScecret(pwdHash, client.Id)
}
func addClientScecret(secret string, clientid int) {
	sc := &models.ClientSecrets{
		Type:     "SHA256",
		ClientId: clientid,
	}
	sc.Value, _ = utils.HashPassword(secret)
	sc.Expiration = time.Now().AddDate(1, 0, 0)

	err := db.Save(sc).Error
	if err != nil {
		panic(err)
	}
}

func addUserRole(uid, ouid, rid string) {

	ur := &models.UserRoles{
		RoleId: rid,
		UserId: uid,
		OUId:   ouid,
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
