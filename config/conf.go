package config

import (
	"fmt"
	"github.com/Centny/gwf/util"
)

var (
	AppId         string //开放授权平台id
	AppSecret     string //开放授权平台密钥
	TimeDeviation int64  //服务器时间差

	AppraiseSrvAddr string
	ArsSrvAddr      string
	CountSrvAddr    string
	CourseSrvAddr   string
	DmsSrvAddr      string
	ExtraSrvAddr    string
	FsSrvAddr       string
	ImsSrvAddr      string
	OrderSrvAddr    string
	Pes2SrvAddr     string
	RecruitSrvAddr  string
	UcsSrvAddr      string
	TmsSrvAddr      string
	RcpSrvAddr      string
	AdmSrvAddr      string
)

func InitOpenAuthConf(appId, appSecret string, timeDeviation int64) {
	if appId == "" {
		panic("appId is nil")
	}

	if appSecret == "" {
		panic("appSecret is nil")
	}

	AppId = appId
	AppSecret = appSecret
	TimeDeviation = timeDeviation
}

func InitSrvAddr(cfg *util.Fcfg) {
	if cfg == nil {
		panic("config is nil")
	}

	cfg.Print()

	AppraiseSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/APPRAISE_PUB_HOST"))
	ArsSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/ARS_PUB_HOST"))
	CountSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/COUNT_PUB_HOST"))
	CourseSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/COURSE_PUB_HOST"))
	DmsSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/DMS_PUB_HOST"))
	ExtraSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/EXTRA_PUB_HOST"))
	FsSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/FS_PUB_HOST"))
	ImsSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/IMSD_PUB_HOST"))
	OrderSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/ORDER_PUB_HOST"))
	Pes2SrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/PES2_PUB_HOST"))
	RecruitSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/RECRUIT_PUB_HOST"))
	UcsSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/UCS_PUB_HOST"))
	TmsSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/TMS_PUB_HOST"))
	RcpSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/RCP_PUB_HOST"))
	AdmSrvAddr = fmt.Sprintf("%s//%s", cfg.Val("public/PUBLIC_PROTO"), cfg.Val("public/ADM_PUB_HOST"))
}

func InitTestSrv() {
	cfg := util.NewFcfg3()
	cfg.InitWithUri("http://pb.dev.jxzy.com/dyf_sandbox_conf.properties")

	InitSrvAddr(cfg)
}
