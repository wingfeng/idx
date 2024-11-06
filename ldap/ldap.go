package ldap

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strings"

	"github.com/jimlambrt/gldap"
	"github.com/wingfeng/idx/service"
)

var us *service.UserService

func StartLdapServer(s *service.UserService) {
	us = s
	slog.Info("Starting LDAP server")
	ldapServer, err := gldap.NewServer()
	if err != nil {
		log.Fatalf("unable to create server: %s", err.Error())
	}
	r, err := gldap.NewMux()
	if err != nil {
		log.Fatalf("unable to create router: %s", err.Error())
	}
	r.Bind(bindHandler)
	r.Search(searchHandler)
	r.Add(func(w *gldap.ResponseWriter, r *gldap.Request) {
		slog.Debug("Add request")
		resp := r.NewResponse(gldap.WithResponseCode(gldap.ResultNotSupported))
		defer func() {
			w.Write(resp)
		}()
	})
	r.Delete(func(w *gldap.ResponseWriter, r *gldap.Request) {
		slog.Debug("Delete request")

		resp := r.NewResponse(gldap.WithResponseCode(gldap.ResultNotSupported))
		defer func() {
			w.Write(resp)
		}()
	})
	ldapServer.Router(r)
	go ldapServer.Run(":10389") // listen on port 10389

	// stop server gracefully when ctrl-c, sigint or sigterm occurs
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	select {
	case <-ctx.Done():
		log.Printf("\nstopping directory")
		ldapServer.Stop()
	}
	// ldapServer.AddRootDN("dc=example,dc=com")
	// ldapServer.AddBindFunc("", BindHandler)
	// ldapServer.AddSearchFunc("", SearchHandler)
	// ldapServer.ListenAndServe(":1389")
	// log.Fatal(ldapServer.ListenAndServe(":1389"))
}
func bindHandler(w *gldap.ResponseWriter, r *gldap.Request) {
	resp := r.NewBindResponse(
		gldap.WithResponseCode(gldap.ResultInvalidCredentials),
	)
	defer func() {
		w.Write(resp)
	}()

	m, err := r.GetSimpleBindMessage()
	if err != nil {
		log.Printf("not a simple bind message: %s", err)
		return
	}

	var userName string
	if strings.Contains(m.UserName, "@") {
		userName = strings.Split(m.UserName, "@")[0]
	} else if strings.Contains(m.UserName, "\\") {
		userName = strings.Split(m.UserName, "\\")[1]
	} else {
		userName = m.UserName
	}
	slog.Debug("bindHandler:", "user name", userName)
	if us.VerifyPassword(userName, string(m.Password)) {
		resp.SetResultCode(gldap.ResultSuccess)
		slog.Info("LDAP 用户认证成功", "username", m.UserName)
		return
	}
}
func searchHandler(w *gldap.ResponseWriter, r *gldap.Request) {
	resp := r.NewSearchDoneResponse()
	defer func() {
		w.Write(resp)
	}()
	m, err := r.GetSearchMessage()
	if err != nil {
		log.Printf("not a search message: %s", err)
		return
	}
	log.Printf("search base dn: %s", m.BaseDN)
	log.Printf("search scope: %d", m.Scope)
	log.Printf("search filter: %s", m.Filter)
	us.GetUserByName(m.Filter)
	if strings.Contains(m.Filter, "uid=admin") || m.BaseDN == "ou=users,dc=idx,dc=com" {
		entry := r.NewSearchResponseEntry(
			"uid=admin,ou=people,cn=example,dc=org",
			gldap.WithAttributes(map[string][]string{
				"objectclass": {"top", "person", "organizationalPerson", "inetOrgPerson"},
				"uid":         {"alice"},
				"cn":          {"alice eve smith"},
				"givenname":   {"alice"},
				"sn":          {"smith"},
				"ou":          {"people"},
				"description": {"friend of Rivest, Shamir and Adleman"},
				"password":    {"{SSHA}U3waGJVC7MgXYc0YQe7xv7sSePuTP8zN"},
			}),
		)
		entry.AddAttribute("email", []string{"alice@example.org"})
		w.Write(entry)
		resp.SetResultCode(gldap.ResultSuccess)
	}
	if m.BaseDN == "ou=users,dc=idx,dc=com" {
		entry := r.NewSearchResponseEntry(
			"ou=people,cn=example,dc=org",
			gldap.WithAttributes(map[string][]string{
				"objectclass": {"organizationalUnit"},
				"ou":          {"people"},
			}),
		)
		w.Write(entry)
		resp.SetResultCode(gldap.ResultSuccess)
	}
	return
}
