package controller

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shaxiaozz/sangfor-ad-exporter/cache"
	"github.com/shaxiaozz/sangfor-ad-exporter/global"
	"github.com/shaxiaozz/sangfor-ad-exporter/model"
	"github.com/shaxiaozz/sangfor-ad-exporter/pkg/sangfor_ad"
)

type MetricsCollector struct {
}

var tokenCache = cache.NewTokenCache()

// 初始化相关指标
var (
	// 初始化指标: nacos_service_instance_count
	SangforAdVirtualServiceHttpRequestRate = prometheus.NewDesc(
		"sangfor_ad_vs_http_request_rate",
		"Virtual service HTTP request rate.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)
)

// Prometheus Describe接口传递指标描述符到 channel
func (m *MetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- SangforAdVirtualServiceHttpRequestRate
}

func (m *MetricsCollector) Collect(ch chan<- prometheus.Metric) {
	// 从缓存获取token
	token, err := getToken()
	if err != nil {
		global.Logger.Fatal(fmt.Sprintf("获取 Sangfor AD Token 失败: %v", err))
	}

	// 获取虚拟服务状态信息
	data, err := sangfor_ad.VirtualServiceStat(token.Name)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("获取Sangfor AD 虚拟服务状态信息失败: %v", err))
		return
	}

	for _, v := range data.Items {
		// 写入nacos_service_instance_count指标
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceHttpRequestRate,
			prometheus.GaugeValue, float64(v.HttpRequestRate.Value), global.Config.SangforAd.DeviceName, v.Name, v.HttpRequestRate.Model, v.HttpRequestRate.Unit)
	}
}

func getToken() (*model.SangforAdLoginResp, error) {
	return tokenCache.Get(func() (*model.SangforAdLoginResp, error) {
		return sangfor_ad.Login(&model.SangforAdLoginReq{
			Username: global.Config.SangforAd.Username,
			Password: global.Config.SangforAd.Password,
		})
	})
}
