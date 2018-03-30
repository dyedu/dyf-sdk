package config

import (
	"fmt"
	"github.com/Centny/gwf/util"
	"gopkg.in/bson.v2"
	"strings"
	"testing"
)

func TestInitOpenAuthConf(t *testing.T) {
	func() {
		defer func() {
			err := recover()
			t.Log(err)
		}()

		InitOpenAuthConf("", "xx", 0)
	}()

	func() {
		defer func() {
			err := recover()
			t.Log(err)
		}()

		InitOpenAuthConf("xx", "", 0)
	}()

	appId := bson.NewObjectId().Hex()
	appSecret := bson.NewObjectId().Hex()

	InitOpenAuthConf(appId, appSecret, 321)
	if appId != AppId {
		t.Error("appId error")
		return
	}
	if appSecret != AppSecret {
		t.Error("appSecret error")
		return
	}
	if TimeDeviation != 321 {
		t.Error("TimeDeviation error")
		return
	}
}

func TestInitSrvAddr(t *testing.T) {
	func() {
		defer func() {
			err := recover()
			t.Log(err)
		}()

		InitSrvAddr(nil)
	}()

	cfg := util.NewFcfg3()
	for _, val := range []string{"appraise", "ars", "count", "course", "dms", "extra", "fs", "imsd", "order", "pes2", "recruit", "ucs", "tms", "rcp", "adm"} {
		key := fmt.Sprintf("public/%s_PUB_HOST", strings.ToUpper(val))
		cfg.SetVal(key, val)
	}

	InitSrvAddr(cfg)
	if AppraiseSrvAddr != "//appraise" {
		t.Error("AppraiseSrvAddr error")
		return
	}
	if ArsSrvAddr != "//ars" {
		t.Error("ArsSrvAddr error")
		return
	}
	if CountSrvAddr != "//count" {
		t.Error("CountSrvAddr error")
		return
	}
	if CourseSrvAddr != "//course" {
		t.Error("CourseSrvAddr error")
		return
	}
	if DmsSrvAddr != "//dms" {
		t.Error("DmsSrvAddr error")
		return
	}
	if ExtraSrvAddr != "//extra" {
		t.Error("ExtraSrvAddr error")
		return
	}
	if FsSrvAddr != "//fs" {
		t.Error("FsSrvAddr error")
		return
	}
	if ImsSrvAddr != "//imsd" {
		t.Error("ImsSrvAddr error")
		return
	}
	if OrderSrvAddr != "//order" {
		t.Error("OrderSrvAddr error")
		return
	}
	if Pes2SrvAddr != "//pes2" {
		t.Error("Pes2SrvAddr error")
		return
	}
	if RecruitSrvAddr != "//recruit" {
		t.Error("RecruitSrvAddr error")
		return
	}
	if UcsSrvAddr != "//ucs" {
		t.Error("UcsSrvAddr error")
		return
	}
	if TmsSrvAddr != "//tms" {
		t.Error("SrvAddr error")
		return
	}
	if RcpSrvAddr != "//rcp" {
		t.Error("RcpSrvAddr error")
		return
	}
	if AdmSrvAddr != "//adm" {
		t.Error("AdmSrvAddr error")
		return
	}
}
