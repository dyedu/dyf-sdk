package sso

import (
	"fmt"
	"github.com/Centny/gwf/log"
	"github.com/Centny/gwf/routing"
	"github.com/Centny/gwf/util"
	"net/http"
	"net/url"
	"regexp"
)

var TokenReg = regexp.MustCompile("\\&?\\b*token=[^\\$\\&=#]*")

type AuthFilter struct {
	SsoAuthUrl  string
	SsoLoginUrl string
	Pre         string
	M           string //the filter model,C is return the 302 code.
	Options     []*regexp.Regexp
	Optioned    bool
}

func (a *AuthFilter) SrvHTTP(hs *routing.HTTPSession) routing.HResult {
	turl := TokenReg.ReplaceAllString(hs.R.URL.String(), "")
	protocol := "http"
	if hs.R.TLS != nil {
		protocol = "https"
	}

	redirectUrl := url.QueryEscape(fmt.Sprintf("%s://%s%s%s", protocol, hs.R.Host, a.Pre, turl))
	loginUrl := fmt.Sprintf(a.SsoLoginUrl, redirectUrl)

	token := hs.CheckValA("token")
	log.I("AuthFilter->doing by token(%v),redirect(%v)", token, redirectUrl)
	if token == "" {
		return a.OnNotLogin(hs, loginUrl)
	}

	uid, err := a.requestAuth(token)
	if err != nil {
		return a.OnNotLogin(hs, loginUrl)
	}

	var u = false
	if uid != hs.StrVal("uid") {
		hs.SetVal("uid", uid) //set uid to session
		u = true
	}
	if token != hs.StrVal("token") {
		hs.SetVal("token", token) //set token to session
		u = true
	}
	if token != hs.Cookie("token") {
		hs.SetCookie("token", token) //set token to cookie
		u = true
	}
	if u {
		hs.S.Flush()
	}
	return routing.HRES_CONTINUE
}

func (a *AuthFilter) OnNotLogin(hs *routing.HTTPSession, loginUrl string) routing.HResult {
	hs.SetVal("token", "")
	hs.SetCookie("token", "")
	if a.Optioned {
		return routing.HRES_CONTINUE
	}
	for _, reg := range a.Options {
		if reg.MatchString(hs.R.URL.Path) {
			return routing.HRES_CONTINUE
		}
	}
	m := hs.RVal("m")
	if len(m) < 1 {
		m = a.M
	}
	switch m {
	case "C":
		hs.MsgRes2(http.StatusMovedPermanently, loginUrl)
	default:
		hs.Redirect(loginUrl)
	}
	return routing.HRES_RETURN
}

func (a *AuthFilter) requestAuth(token string) (string, error) {
	res, err := util.HGet2(a.SsoAuthUrl, token)
	if err != nil {
		err = util.Err("requestAuth by token(%v) result error(%v)", token, err)
		log.E("[AuthFilter.requestAuth] %v", err)
		return "", nil
	}

	uid := res.StrVal("data")
	if uid == "" {
		err = util.Err("requestAuth by token(%v) result uid is nil, dmsg: %v", token, res.StrVal("dmsg"))
		log.E("[AuthFilter.requestAuth] %v", err)
		return "", err
	}
	return uid, nil
}

/**
@arg:
	authUrl		鉴权接口地址
	loginUrl	登录页地址
	pre			接口前缀
@desc:
	api鉴权过滤器，通过url上token参数或者cookie里的token获取uid
@ret:
	*AuthFilter	api鉴权过滤器
@author:
    lujz create on 2018-03-06
*/
func NewAuthFilter(authUrl, loginUrl, pre string) *AuthFilter {
	if authUrl == "" || loginUrl == "" {
		panic(fmt.Sprintf("initial AuthFilter error(auth: %v, login: %v, pre: %v)", authUrl, loginUrl, pre))
	}

	log.I("initial AuthFilter(auth: %v, login: %v, pre: %v)", authUrl, loginUrl, pre)
	return &AuthFilter{
		SsoAuthUrl:  authUrl,
		SsoLoginUrl: loginUrl,
		Pre:         pre,
		Options:     []*regexp.Regexp{},
	}
}
