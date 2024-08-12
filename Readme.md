
# IDx an Identity Server by Golang  

## About  

功能目标：利用go实现一个SSO单点登录服务器  
参考Skoruba.IdentityServer4.Admin逐步满足IdentityServer4的功能，基于[github.com/wingfeng/idx-oauth2](https://pkg.go.dev/github.com/wingfeng/idx-oauth2)重写了oauth2的go实现。 

## Features

* 支持Oauth2  
* 支持OpenID  
* 支持Saml登录（todo)
* 支持LDAP?(todo)  
* 支持命令行控制CLI（todo)  

通过Gin实现对OIDC的Flow进行单元测试，单元测试过程参看:<https://openid.net/specs/openid-connect-core-1_0.html>  
经过测试，已经可以支持Wordpress,NextCloud的OIDC插件实现SSO。  

增加K8S Helm Chart
Docker的配置以IDX作为前缀  
IDX_CONNECTION  
IDX_PORT  
IDX_DRIVER  
IDX_HTTPSCHEME  

生成测试数据
func /test/initdb_test.go/TestSeedData  
