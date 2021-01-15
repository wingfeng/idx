package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/cihub/seelog"
)

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := verifyAuthorizationToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if claims != nil {
		id := claims["sub"]
		user, err := UserStore.GetUserByID(id.(string))
		result := make(map[string]interface{})
		result["sub"] = user.ID
		result["email"] = user.Email
		result["email_verified"] = user.EmailConfirmed
		result["display_name"] = user.DisplayName
		result["ou"] = user.OU
		result["ouid"] = user.OUID
		if err != nil {
			log.Errorf("获取用户错误,Error:%s", err.Error())
		}
		json.NewEncoder(w).Encode(result)
	} else {
		log.Errorf("解析Token错误,Validate:%b,Error:%s", err.Error())
	}
}
