package controller

import (
	"fmt"
	"go-ops/internal/service"
	"go-ops/pkg/util"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
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
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}

	app, err := service.App().GetApp(r.GetCtx(), appid)
	if err != nil {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}

	signature := r.GetHeader(GoOpsHeaderSignature)
	timestamp := r.GetHeader(GoOpsHeaderTimestamp)
	nonce := r.GetHeader(GoOpsHeaderNonce)
	body := r.GetBody()
	sign := util.GetSign(app.ApiKey, app.SecKey, nonce, timestamp, body)

	if sign != signature {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}

	fmt.Println("appid=", appid)
	r.Middleware.Next()

}

func AuthUser(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
