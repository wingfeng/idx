package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx/store"
)

type UserInfoController struct {
	UserStore *store.DbUserStore
}
type emptyStruct struct{}

func (ctrl *UserInfoController) UserInfo(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Header("Content-Type", "application/json")

	token, err := Srv.ValidationBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if token != nil {
		id := token.GetUserID()
		sp := strings.Split(token.GetScope(), " ")
		scopes := make(map[string]*emptyStruct)
		for _, sc := range sp {
			scopes[sc] = &emptyStruct{}
		}
		user, err := ctrl.UserStore.GetUserByID(id)

		result := make(map[string]interface{})
		result["sub"] = user.ID
		result["display_name"] = user.DisplayName
		result["preferred_username"] = user.Account

		if scopes["email"] != nil {
			result["email"] = user.Email
			result["email_verified"] = user.EmailConfirmed
		}

		if scopes["profile"] != nil {
			result["ou"] = user.OU
			result["ouid"] = user.OUID
		}
		if err != nil {
			log.Errorf("获取用户错误,Error:%s", err.Error())
		}
		json.NewEncoder(w).Encode(result)
	} else {
		log.Errorf("解析Token错误,Validate:%b,Error:%s", err.Error())
	}
}
