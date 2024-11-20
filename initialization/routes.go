package initialization

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"watchAlert/internal/global"
	"watchAlert/internal/middleware"
	"watchAlert/internal/routers"
	"watchAlert/internal/routers/v1"
)

func InitRoute() {
	logc.Info(context.Background(), "服务启动")

	mode := global.Config.Server.Mode
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	ginEngine := gin.New()

	ginEngine.Use(
		// 启用CORS中间件
		middleware.Cors(),
		// 自定义请求日志格式
		gin.Recovery(),
		//gin.LoggerWithFormatter(middleware.RequestLoggerFormatter),
	)
	allRouter(ginEngine)

	err := ginEngine.Run(":" + global.Config.Server.Port)
	if err != nil {
		logc.Error(context.Background(), "服务启动失败:", err)
		return
	}
}

func allRouter(engine *gin.Engine) {

	routers.HealthCheck(engine)
	v1.Router(engine)

}
