package tools

import (
	"fmt"
	"github.com/Centny/gwf/util"
	"github.com/dyedu/dyf-sdk/config"
	"gopkg.in/bson.v2"
	"testing"
)

func TestCreateOpenAuthSign(t *testing.T) {
	var (
		res string
	)

	config.InitOpenAuthConf(bson.NewObjectId().Hex(), bson.NewObjectId().Hex(), 0)

	if res = CreateOpenAuthSign("xx"); res != "" {
		t.Error("invalid method error")
		return
	}
	if res = CreateOpenAuthSign("md5"); res != util.Md5_b([]byte(fmt.Sprintf("appId=%v&appSecret=%v&authMethod=%v&timestamp=%v", config.AppId, config.AppSecret, "md5", util.Now()+config.TimeDeviation))) {
		t.Error("md5 result error")
		return
	}
	if res = CreateOpenAuthSign("sha1"); res != util.Sha1_b([]byte(fmt.Sprintf("appId=%v&appSecret=%v&authMethod=%v&timestamp=%v", config.AppId, config.AppSecret, "sha1", util.Now()+config.TimeDeviation))) {
		t.Error("sha1 result error")
		return
	}
}

func TestOpenAuthParam(t *testing.T) {
	var (
		res string
	)

	config.InitOpenAuthConf(bson.NewObjectId().Hex(), bson.NewObjectId().Hex(), 0)

	if res = OpenAuthParam("md5"); res != fmt.Sprintf("openAppId=%v&openAuthMethod=%v&openAuthSign=%v&openAuthTs=%v",
		config.AppId, "md5", CreateOpenAuthSign("md5"), util.Now()+config.TimeDeviation) {
		t.Error("md5 result error")
		return
	}
	if res = OpenAuthParam("sha1"); res != fmt.Sprintf("openAppId=%v&openAuthMethod=%v&openAuthSign=%v&openAuthTs=%v",
		config.AppId, "sha1", CreateOpenAuthSign("sha1"), util.Now()+config.TimeDeviation) {
		t.Error("sha1 result error")
		return
	}
}
