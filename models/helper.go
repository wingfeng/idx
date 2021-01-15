package models

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

//Sync2Db 将struct同步数据结构到数据库
func Sync2Db(x *gorm.DB) {
	//	x.DisableForeignKeyConstraintWhenMigrating = true

	// 同步结构体与数据表
	err := x.AutoMigrate(

		new(APIClaims),
		new(APIProperties),
		new(APIResources),
		new(APIScopeClaims),
		new(APIScopes),
		new(APISecrets),
		new(AuditLog),

		new(ClientClaims),
		new(ClientCorsOrigins),
		new(ClientGrantTypes),
		new(ClientIDPRestrictions),
		new(ClientRedirectURIs),
		new(ClientProperties),
		new(ClientPostLogoutRedirectURIs),
		new(Client),
		new(ClientScopes),
		new(ClientSecrets),
		new(DeviceCodes),
		new(IDentityClaims),
		new(IDentityProperties),
		new(IDentityResources),
		new(OrganizationUnit),
		new(PersistedGrants),
		new(RoleClaims),
		new(Role),
		new(UserClaims),
		new(UserLogins),
		new(UserRoles),
		new(User),
		new(UserTokens),
	)
	if err != nil {
		log.Errorf("同步数据结构错误,Error:%v", err)
	}
}
