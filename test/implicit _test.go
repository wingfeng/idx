package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/idx/core"
)

func Test_Implicit(t *testing.T) {
	router := init_router()
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/connect/authorize?response_type=id_token%20token&client_id=implicit_client&redirect_uri=http%3A%2F%2Flocalhost:9000%2Fcb&scope=openid%20profile&state=af0ifjsldkj&nonce=n-0S6_WzA2Mj", nil)
	router.ServeHTTP(recorder, req)
	//	core.DumResponse(os.Stdout, "authorize GET", recorder.Result())
	cookies := recorder.Result().Cookies()
	assert.Equal(t, 302, recorder.Code)

	// assert.Equal(t, "pong", recorder.Body.String())

	req, err := http.NewRequest("POST", "/login", nil)
	if err != nil {
		t.Logf("Error:%s", err.Error())
	}
	for _, c := range cookies {
		req.AddCookie(c)
	}
	req.Form = make(url.Values)
	req.Form.Add("username", "admin")
	req.Form.Add("password", "fire@123")
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	core.DumResponse(os.Stdout, "Login POST", recorder.Result())
	assert.Equal(t, 302, recorder.Code)
	req, _ = http.NewRequest("GET", "/connect/authorize?response_type=id_token%20token&client_id=implicit_client&redirect_uri=http%3A%2F%2Flocalhost:9000%2Fcb&scope=openid%20profile&state=af0ifjsldkj&nonce=n-0S6_WzA2Mj", nil)

	for _, c := range cookies {
		req.AddCookie(c)
	}
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	core.DumResponse(os.Stdout, "authorize GET", recorder.Result())
	assert.Equal(t, recorder.Code, 302)
}
