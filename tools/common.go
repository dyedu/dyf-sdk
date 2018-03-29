package tools

import (
	"fmt"
	"github.com/Centny/gwf/util"
	"github.com/dyedu/dyf-sdk/config"
)

func CreateOpenAuthSign(method string) string {
	raw := fmt.Sprintf("appId=%v&appSecret=%v&authMethod=%v&timestamp=%v",
		config.AppId, config.AppSecret, method, util.Now()+config.TimeDeviation)
	switch method {
	case "md5":
		return util.Md5_b([]byte(raw))
	case "sha1":
		return util.Sha1_b([]byte(raw))
	default:
		return ""
	}
}

func OpenAuthParam(method string) string {
	return fmt.Sprintf("openAppId=%v&openAuthMethod=%v&openAuthSign=%v&openAuthTs=%v",
		config.AppId, method, CreateOpenAuthSign(method), util.Now()+config.TimeDeviation)
}
