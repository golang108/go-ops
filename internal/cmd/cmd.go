package cmd

import (
	"context"

	"github.com/gogf/gf/v2/protocol/goai"

	"osp/internal/consts"
	"osp/internal/controller"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/swagger/v2"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCORS)
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
					controller.ScritptTask,
					controller.PeerManagaer,
					controller.DownloadFileTask,
				)
			})

			s.SetIndexFolder(true)
			s.AddSearchPath("public")
			s.AddStaticPath("/public", "public")
			s.Plugin(&swagger.Swagger{})
			s.SetPort(8199)
			enhanceOpenAPIDoc(s)
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "GoFrame",
			URL:  "https://goframe.org",
		},
	}
}
