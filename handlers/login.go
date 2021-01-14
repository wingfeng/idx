package handlers

import (
	"net/http"
	"os"

	"github.com/go-session/session"
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

			w.Header().Set("Location", "/auth")
			w.WriteHeader(http.StatusFound)
			return
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

	}
	outputHTML(w, r, "../static/login.html")
}
func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
