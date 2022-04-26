package controller

import (
	"fmt"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	GoOpsHeaderAppId = "GO-OPS-X-APPID"
)

func MiddlewareGetApp(r *ghttp.Request) {

	appid := r.GetHeader(GoOpsHeaderAppId)
	if appid == "" {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}

	fmt.Println("appid=", appid)
	r.Middleware.Next()

}
