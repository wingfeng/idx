package controller

import (
	"log/slog"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx-oauth2/endpoint"
	"github.com/wingfeng/idx-oauth2/model/request"
	"github.com/wingfeng/idx/models/dto"
	"github.com/wingfeng/idx/service"
)

type AuthController struct {
	userService *service.UserService
}
type changePwdRequest struct {
	OldPassword string `form:"oldpassword" binding:"required"`
	NewPassword string `form:"newpassword" binding:"required"`
	Username    string `form:"username" binding:"required"`
	redirect    string `form:"redirect"`
}

func NewAuthController(s *service.UserService) *AuthController {
	return &AuthController{userService: s}
}
func (ctrl *AuthController) RegistRoute(route gin.IRouter) {
	route.POST("/changepwd", ctrl.PostChangePassword)
	route.GET("/changepwd", ctrl.GetChangePassword)
}
func (ctrl *AuthController) GetLogin(ctx *gin.Context) {
	ctrl.loadLogin(ctx, "/idx/", "")
}
func (ctrl *AuthController) PostLogin(ctx *gin.Context) {
	loginRequest := &request.LoginRequest{}

	if err := ctx.ShouldBind(&loginRequest); err != nil {
		ctrl.loadLogin(ctx, loginRequest.Redirect, err.Error())
		return
	}
	if ctrl.userService.VerifyPassword(loginRequest.UserName, loginRequest.Password) {
		link := loginRequest.Redirect
		//if user is temporary password ask user change pwd first;

		user, err := ctrl.userService.GetUserByName(loginRequest.UserName)
		if err != nil {
			ctrl.loadLogin(ctx, link, "")
			return
		}
		u := user.(*dto.UserDto)
		if u.IsTemporaryPassword {
			ctx.HTML(401, "changepwd.html", gin.H{
				"errMsg": "Your password is temporary, please change it first",
				"tenant": "idx",
			})
			return
		}
		session := sessions.Default(ctx)
		session.Set(endpoint.Const_Principle, loginRequest.UserName)
		session.Save()

		if strings.EqualFold(link, "") {
			link, _ = url.JoinPath("idx", "/")
		}
		slog.Info("User logined", "user", loginRequest.UserName, "redirect to", link)
		ctx.Redirect(302, link)
	} else {
		ctrl.loadLogin(ctx, loginRequest.Redirect, "invalid username or password")
		return
	}
}
func (ctrl *AuthController) GetChangePassword(ctx *gin.Context) {
	ctx.HTML(200, "changepwd.html", gin.H{"errMsg": "", "tenant": "idx"})
}
func (ctrl *AuthController) PostChangePassword(ctx *gin.Context) {
	var request changePwdRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.HTML(400, "changepwd.html", gin.H{"errMsg": err.Error()})
		return
	}
	err := ctrl.userService.ChangePassword(request.Username, request.OldPassword, request.NewPassword)
	if err != nil {
		slog.Error("ChangePassword Error", "error", err)
		ctx.HTML(400, "changepwd.html", gin.H{"errMsg": "invalid old password", "tenant": "idx"})
		return
	}
	ctrl.loadLogin(ctx, request.redirect, "password changed please login again")
}
func (ctrl *AuthController) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete(endpoint.Const_Principle)
	session.Save()
	//root, _ := url.JoinPath(ctrl.Config.TenantPath, "/")
	ctx.Redirect(302, "")
}
func (ctrl *AuthController) loadLogin(ctx *gin.Context, redirect, msg string) {
	if strings.EqualFold(redirect, "") {
		redirect = "/idx/"
	}
	ctx.HTML(401, "login.html", gin.H{
		"redirect": redirect,
		"tenant":   "idx",
		"errMsg":   msg,
	})
}
