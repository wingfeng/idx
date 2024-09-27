package models

import (
	"log/slog"

	"gorm.io/gorm"
)

// Sync2Db 将struct同步数据结构到数据库
func Sync2Db(x *gorm.DB) {
	x.DisableForeignKeyConstraintWhenMigrating = true

	// 同步结构体与数据表
	err := x.AutoMigrate(

		new(APIResources),

		new(APIScopes),
		new(APISecrets))
	if err != nil {
		slog.Error("同步数据结构错误,Error:", "error", err)
	}
	err = x.AutoMigrate(

		new(ClientCorsOrigins),

		new(ClientPostLogoutRedirectURIs),
		new(Client),

		new(ClientSecrets),

		new(IdentityClaims),
		new(IdentityProperties),
		new(IdentityResources),
		new(OrganizationUnit),
		new(PersistedGrants),

		new(Role),

		new(UserLogins),
		new(UserRoles),
		new(User),
		new(UserTokens),
	)
	if err != nil {
		slog.Error("同步数据结构错误,Error", "error", err)
	}
	err = x.AutoMigrate(new(Authorization_fake))
	if err != nil {
		slog.Error("同步数据结构错误 Authorization,Error", "error", err)
	}

}
