package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/idx/core"
)

func Test_Codeflow(t *testing.T) {
	router := init_router()
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/connect/authorize?response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A7000%2Fcallback%3Fclient_name%3DOidcClient&state=2CWRZW1KxoM-EBpTCATOCAH7GMVoJEeuNNSaZn6jVP4&client_id=code_client&scope=openid+profile+email", nil)

	router.ServeHTTP(recorder, req)
	core.DumResponse(os.Stdout, "authorize GET", recorder.Result())
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
	req, _ = http.NewRequest("POST", "/connect/authorize", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	core.DumResponse(os.Stdout, "authorize POST", recorder.Result())
	recorder.Flush()
	header := recorder.HeaderMap
	t.Logf("Resp Redirect:%v", header["Location"])
	callbackURI, _ := url.Parse(header["Location"][0])
	code := callbackURI.Query().Get("code")

	body := fmt.Sprintf("code=%s&redirect_uri=%s&grant_type=authorization_code", code, "http%3A%2F%2Flocalhost%3A7000%2Fcallback%3Fclient_name%3DOidcClient")

	req, _ = http.NewRequest("POST", "/connect/token", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic Y29kZV9jbGllbnQ6Y29kZV9zZWNyZXQ=")
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	core.DumResponse(os.Stdout, "Token", recorder.Result())
	assert.Equal(t, recorder.Code, 200)
}
