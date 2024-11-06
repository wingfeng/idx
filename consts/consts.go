package consts

import "time"

const (
	CONST_TTL = time.Minute * 15

	CONST_CLIENTKEY = "idx:client:id:"

	CONST_ConsentKey = "idx:consent:%s:%s"

	CONST_UserIDKey      = "idx:user:id:"
	CONST_USERNAMEKEY    = "idx:user:name:"
	CONST_USERPWDHashKEY = "idx:user:pwdhash:"
)
