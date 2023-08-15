

功能目标：利用go实现一个SSO单点登录服务器   
参考Skoruba.IdentityServer4.Admin逐步满足IdentityServer4的功能，基于github.com/go-oauth2/oauth2/v4，魔改了一下，实现OpenID部分。   
支持Saml登录（todo）   
支持Oauth2   
支持OpenID   
支持LDAP?   
支持命令行控制CLI   

通过Gin实现对OIDC的Flow进行单元测试，单元测试过程参看:https://openid.net/specs/openid-connect-core-1_0.html   
经过测试，已经可以支持Wordpress,NextCloud的OIDC插件实现SSO。   

增加K8S Helm Chart
Docker的配置以IDX作为前缀          
IDX_CONNECTION    
IDX_PORT   
IDX_DRIVER   
IDX_HTTPSCHEME   

生成测试数据
/test/dbclientstore_test.go