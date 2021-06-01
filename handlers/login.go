package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"github.com/labstack/gommon/log"
	"github.com/wingfeng/idx/store"
	"github.com/wingfeng/idx/utils"
)

type LoginController struct {
	UserStore store.DbUserStore
}

func (ctrl *LoginController) LoginGet(ctx *gin.Context) {
	w := ctx.Writer

	store, err := session.Start(ctx.Request.Context(), ctx.Writer, ctx.Request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, exist := store.Get("LoggedInUserID"); exist {
		ctx.Redirect(http.StatusFound, "/consent")
		return
	}
	ctx.HTML(http.StatusFound, "login.html", nil)
}

func (ctrl *LoginController) LoginPost(ctx *gin.Context) {

	w := ctx.Writer
	r := ctx.Request
	//	core.DumpRequest(os.Stdout, "LoginPost", r)
	store, err := session.Start(ctx.Request.Context(), ctx.Writer, ctx.Request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	userName := r.Form.Get("username")
	pwd := r.Form.Get("password")
	user, err := ctrl.UserStore.GetUserByAccount(userName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pwdVerified, err := utils.VerifyPassword(user.PasswordHash, pwd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if pwdVerified {

		store.Set("LoggedInUserID", user.ID)
		store.Save()
		requireConsent := false
		if v, ok := store.Get("ReturnUri"); ok {

			form := v.(map[string]interface{})
			clientID := form["client_id"].([]interface{})[0].(string)
			requireConsent = needConsent(clientID, user.ID)
		}
		if requireConsent {
			w.Header().Set("Location", "/consent")
			w.WriteHeader(http.StatusFound)
		} else {
			ctx.Redirect(http.StatusFound, "/connect/authorize")
		}
		return
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

}
func needConsent(clientID, userID string) bool {
	client, err := ClientStore.GetByID(nil, clientID)
	if err != nil {
		log.Errorf("获取Client:%s信息错误!Error:%s", clientID, err.Error())
		return false
	}
	//todo:获取保存好的consent信息，如果已经有以保存的consent信息即可直接跳过。
	return client.GetRequireConsent()
}
