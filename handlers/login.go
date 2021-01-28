package handlers

import (
	"net/http"

	"github.com/go-session/session"
	"github.com/labstack/gommon/log"
	"github.com/wingfeng/idx/store"
	"github.com/wingfeng/idx/utils"
)

var UserStore *store.DbUserStore

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, exist := store.Get("LoggedInUserID"); exist {
		w.Header().Set("Location", "/auth")
		w.WriteHeader(http.StatusFound)
		return
	}
	if r.Method == "POST" {
		if r.Form == nil {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		userName := r.Form.Get("username")
		pwd := r.Form.Get("password")
		user, err := UserStore.GetUserByAccount(userName)

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
				w.Header().Set("Location", "/auth")
				w.WriteHeader(http.StatusFound)
			} else {
				w.Header().Set("Location", "/connect/authorize")
				w.WriteHeader(http.StatusFound)
			}

			return
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

	}
	outputHTML(w, r, "../static/login.html")
}
func needConsent(clientID, userID string) bool {
	client, err := ClientStore.GetByID(nil, clientID)
	if err != nil {
		log.Errorf("获取Client:%s信息错误!Error:%s", clientID, err.Error())
		return false
	}
	//todo:获取保存好的consent信息，如果已经有以保存的consent信息即可直接跳过。
	return client.RequireConsent
}
