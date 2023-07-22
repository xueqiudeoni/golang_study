package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zap_log/dao/mysql"
	"zap_log/dao/redis"
	"zap_log/logger"
	"zap_log/routers"
	"zap_log/settings"
)

func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("init setting failed,err:", err)
		return
	}
	//2、初始化日志
	if err := logger.Init(); err != nil {
		fmt.Println("init logger failed,err:", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	//3、初始化MySQL连接
	if err := mysql.Init(); err != nil {
		fmt.Println("init mysql failed,err:", err)
		return
	}
	defer mysql.Close()
	//4、初始化redis连接
	if err := redis.Init(); err != nil {
		fmt.Println("init redis failed,err:", err)
		return
	}
	defer redis.Close()
	//5、注册路由
	r := routers.SetupRouter()
	//6、启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Info("listen err:", zap.Error(err))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("shutdown server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("server shutdown:", zap.Error(err))
	}
	zap.L().Info("server exiting")
}
