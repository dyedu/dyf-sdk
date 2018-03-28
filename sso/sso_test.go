package amazonapi

import (
	"fmt"
	"github.com/Centny/gwf/routing"
	"github.com/Centny/gwf/routing/httptest"
	"github.com/Centny/gwf/util"
	"github.com/dyedu/dyf-sdk/config"
	"strings"
	"testing"
)

var ts *httptest.Server

var SrvAddr = func() string {
	panic("SrvAddr is not initial")
}

func init() {
	ts = httptest.NewMuxServer()

	func() {
		defer func() {
			recover()
		}()
		SrvAddr()
	}()

	SrvAddr = func() string {
		return ts.URL
	}

	config.InitOpenAuthConf("kxOAea310912d49b2423b8540010d2eace52c9c9c4b85179b201c41ba22210cf2b0d",
		"bac39fe36553d7369302c721bb42087a3656184778097f33134db6321b740a1d7ab3a947e2f709654958d62c9d181e500794a259b3479886c04e5d496d2bb7fa",
		-10000)

	cfg := util.NewFcfg3()
	for _, val := range []string{"appraise", "ars", "count", "course", "dms", "extra", "fs", "imsd", "order", "pes2", "recruit", "ucs", "tms"} {
		key := fmt.Sprintf("%s_PUB_HOST", strings.ToUpper(val))
		cfg.SetVal(key, fmt.Sprintf("https://%s.dev.gdy.io", val))
	}
	config.InitSrvAddr(cfg)
}

func TestAuthFilter(t *testing.T) {
	usrFilter := NewAuthFilter("https://sso.dev.gdy.io/sso/api/auth?token=%v", "https://sso.dev.gdy.io/sso/index.html?url=%v", "")
	usrFilter.M = "C"
	pubFilter := NewAuthFilter("https://sso.dev.gdy.io/sso/api/auth?token=%v", "https://sso.dev.gdy.io/sso/index.html?url=%v", "")
	pubFilter.M = "C"
	pubFilter.Optioned = true
	ts.Mux.HFilter("^/.*pub/.*$", pubFilter)
	ts.Mux.HFilter("^/.*usr/.*$", usrFilter)

	ts.Mux.HFunc("^/usr/api/uid", func(hs *routing.HTTPSession) routing.HResult {
		uid := hs.StrVal("uid")
		return hs.MsgRes(uid)
	})
	ts.Mux.HFunc("^/pub/api/uid", func(hs *routing.HTTPSession) routing.HResult {
		uid := hs.StrVal("uid")
		return hs.MsgRes(uid)
	})

	var doUsr = func(token string) (string, error) {
		res, err := util.HGet2("%v/usr/api/uid?token=%v", SrvAddr(), token)
		if err != nil {
			return "", nil
		}
		if res.IntVal("code") != 0 {
			return "", util.Err(util.S2Json(res))
		}
		return res.StrVal("data"), nil
	}
	var doPub = func(token string) (string, error) {
		res, err := util.HGet2("%v/pub/api/uid?token=%v", SrvAddr(), token)
		if err != nil {
			return "", nil
		}
		if res.IntVal("code") != 0 {
			return "", util.Err(util.S2Json(res))
		}
		return res.StrVal("data"), nil
	}

	//
	var (
		res string
		err error
	)

	uid, token, err := loginToRcp("ucsadmin", "ucsadmin")
	if err != nil {
		t.Error(err)
		return
	}

	if res, err = doUsr(""); err == nil {
		t.Error("usr error")
		return
	}
	if res, err = doUsr(token); err != nil {
		t.Error(err)
		return
	}
	if res != uid {
		t.Error("uid error")
		return
	}

	if res, err = doPub(""); err != nil {
		t.Error(err)
		return
	}
	if res != "" {
		t.Error("pub error")
		return
	}
	if res, err = doPub(token); err != nil {
		t.Error(err)
		return
	}
	if res != uid {
		t.Error("pub with token error")
		return
	}
}

func loginToRcp(usr, pwd string) (string, string, error) {
	res, err := util.HGet2("https://sso.dev.gdy.io/sso/api/login?usr=%v&pwd=%v&source=PC", usr, pwd)
	if err != nil {
		return "", "", err
	}
	if res.IntVal("code") != 0 {
		return "", "", util.Err(util.S2Json(res))
	}

	return res.StrValP("data/usr/id"), res.StrValP("data/token"), nil
}
