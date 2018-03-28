package config

import (
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

	AppraiseSrvAddr = cfg.Val("APPRAISE_PUB_HOST")
	ArsSrvAddr = cfg.Val("ARS_PUB_HOST")
	CountSrvAddr = cfg.Val("COUNT_PUB_HOST")
	CourseSrvAddr = cfg.Val("COURSE_PUB_HOST")
	DmsSrvAddr = cfg.Val("DMS_PUB_HOST")
	ExtraSrvAddr = cfg.Val("EXTRA_PUB_HOST")
	FsSrvAddr = cfg.Val("FS_PUB_HOST")
	ImsSrvAddr = cfg.Val("IMSD_PUB_HOST")
	OrderSrvAddr = cfg.Val("ORDER_PUB_HOST")
	Pes2SrvAddr = cfg.Val("PES2_PUB_HOST")
	RecruitSrvAddr = cfg.Val("RECRUIT_PUB_HOST")
	UcsSrvAddr = cfg.Val("UCS_PUB_HOST")
	TmsSrvAddr = cfg.Val("TMS_PUB_HOST")
}
