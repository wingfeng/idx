package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/go-session/session"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	var form map[string]interface{}
	if v, ok := store.Get("ReturnUri"); ok {
		// mapVal := v.(map[string]interface{})
		// for m, val := range mapVal {
		// 	form[m] = val.(string)
		// }
		form = v.(map[string]interface{})
	}
	if v, ok := store.Get("state"); ok {
		//r.Form.Set("state", v.(string))
		log.Infof("State:%s", v)
	}
	// 解析指定文件生成模板对象
	tem, err := template.ParseFiles("../static/auth.html")
	if err != nil {
		fmt.Println("读取文件失败,err", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	tem.Execute(w, form)
}
