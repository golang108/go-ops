package controller

import (
	"go-ops/internal/service"
	"go-ops/pkg/util"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

const (
	GoOpsHeaderAppId     = "GO-OPS-X-APPID"
	GoOpsHeaderSignature = "GO-OPS-X-SIGNATURE"
	GoOpsHeaderTimestamp = "GO-OPS-X-TIMESTAMP"
	GoOpsHeaderNonce     = "GO-OPS-X-NONCE"
	GoOpsHeaderToken     = "GO-OPS-X-TOKEN"
)

func MiddlewareGetApp(r *ghttp.Request) {
	appid := r.GetHeader(GoOpsHeaderAppId)
	if appid == "" {
		glog.Error(r.GetCtx(), "appid is empty")
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}

	app, err := service.App().GetApp(r.GetCtx(), appid)
	if err != nil {
		glog.Errorf(r.GetCtx(), "get app:%s err:%v", appid, err)
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}

	signature := r.GetHeader(GoOpsHeaderSignature)
	timestamp := r.GetHeader(GoOpsHeaderTimestamp)
	nonce := r.GetHeader(GoOpsHeaderNonce)
	body := r.GetBody()
	sign := util.GetSign(app.ApiKey, app.SecKey, nonce, timestamp, body)

	if sign != signature {
		glog.Errorf(r.GetCtx(), "sign is not match, req sign:%s server sign:%s", signature, sign)
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}

	r.Middleware.Next()

}

func AuthUser(r *ghttp.Request) {

	appid := r.GetHeader(GoOpsHeaderAppId)
	if appid != "" {
		// 如果是app
		MiddlewareGetApp(r)
	} else {
		service.Auth().MiddlewareFunc()(r)
		r.Middleware.Next()
	}

}
