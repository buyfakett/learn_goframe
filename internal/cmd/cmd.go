package cmd

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"learn_golang/internal/controller/user"
)

var (
	Main = gcmd.Command{
		Name:  "learn_golang",
		Usage: "learn_golang",
		Brief: "learning golang",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 加载 Swagger HTML 模板文件
			htmlFilePath := "resource/template/swagger-ui.html" // 确保路径正确
			htmlContent, err := ioutil.ReadFile(htmlFilePath)
			if err != nil {
				log.Fatalf("无法加载 Swagger HTML 模板文件: %v", err)
			}

			// 启动服务器
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/user", func(users *ghttp.RouterGroup) {
					users.Bind(user.Router())
				})
			})

			// 设置 Swagger UI 路径
			s.SetSwaggerPath("/swagger-ui")
			// 设置加载的 Swagger UI 模板
			s.SetSwaggerUITemplate(string(htmlContent))
			s.SetServerRoot("resource/public")

			// 启动服务
			s.Run()
			return nil
		},
	}
)
