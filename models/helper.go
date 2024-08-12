package models

import (
	"log/slog"

	"gorm.io/gorm"
)

// Sync2Db 将struct同步数据结构到数据库
func Sync2Db(x *gorm.DB) {
	x.DisableForeignKeyConstraintWhenMigrating = true

	// 同步结构体与数据表
	err := x.Debug().AutoMigrate(

		new(APIClaims),
		new(APIProperties),
		new(APIResources),
		new(APIScopeClaims),
		new(APIScopes),
		new(APISecrets))
	if err != nil {
		slog.Error("同步数据结构错误,Error:", "error", err)
	}
	err = x.Debug().AutoMigrate(

		new(ClientClaims),
		new(ClientCorsOrigins),

		new(ClientIdPRestrictions),

		new(ClientProperties),
		new(ClientPostLogoutRedirectURIs),
		new(Client),

		new(ClientSecrets),

		new(IdentityClaims),
		new(IdentityProperties),
		new(IdentityResources),
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
		slog.Error("同步数据结构错误,Error", "error", err)
	}
	err = x.AutoMigrate(new(Authorization_fake))
	if err != nil {
		slog.Error("同步数据结构错误 Authorization,Error", "error", err)
	}

}
