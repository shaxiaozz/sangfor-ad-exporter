package cmd

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shaxiaozz/sangfor-ad-exporter/cache"
	"github.com/shaxiaozz/sangfor-ad-exporter/config"
	"github.com/shaxiaozz/sangfor-ad-exporter/constant"
	"github.com/shaxiaozz/sangfor-ad-exporter/controller"
	"github.com/shaxiaozz/sangfor-ad-exporter/global"
	"github.com/shaxiaozz/sangfor-ad-exporter/model"
	"github.com/shaxiaozz/sangfor-ad-exporter/pkg/sangfor_ad"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start sangfor-ad-exporter",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[sangfor-ad-exporter version] " + constant.Version)
		start()
	},
}

var tokenCache = cache.NewTokenCache()

func start() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// 1、初始化配置
	global.Config = config.InitConfig()

	// 2、初始化日志
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	global.Logger, _ = cfg.Build()

	// 3、获取 token
	data, err := getToken()
	if err != nil {
		global.Logger.Fatal(fmt.Sprintf("获取 Sangfor AD Token 失败: %v", err))
	}
	global.SangforAdToken = data.Name

	// 4、注册 Prometheus Collector
	mc := &controller.MetricsCollector{}
	if err := prometheus.Register(mc); err != nil {
		if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
			global.Logger.Warn("collector 已注册，复用旧实例")
			mc = are.ExistingCollector.(*controller.MetricsCollector)
		} else {
			global.Logger.Fatal(fmt.Sprintf("注册 Prometheus collector 失败: %v", err))
		}
	}

	// 5、HTTP Server
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})
	
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// 6. 启动 HTTP
	go func() {
		global.Logger.Info("sangfor-ad-exporter 启动成功，监听 :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Fatal(fmt.Sprintf("HTTP server 启动失败: %v", err))
		}
	}()

	// 7. 阻塞等待退出信号
	<-ctx.Done()
	global.Logger.Info("收到退出信号，正在优雅关闭...")

	// 8. 优雅关闭（给 Prometheus scrape 留时间）
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		global.Logger.Error(fmt.Sprintf("HTTP server 优雅关闭失败: %v", err))
	}

	global.Logger.Info("sangfor-ad-exporter 已退出")
}

func getToken() (*model.SangforAdLoginResp, error) {
	return tokenCache.Get(func() (*model.SangforAdLoginResp, error) {
		return sangfor_ad.Login(&model.SangforAdLoginReq{
			Username: global.Config.SangforAd.Username,
			Password: global.Config.SangforAd.Password,
		})
	})
}
