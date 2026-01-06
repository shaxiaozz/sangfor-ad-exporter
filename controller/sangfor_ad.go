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

/*
初始化相关指标:
CounterValue  // 只增不减（累计值）
GaugeValue    // 可增可减（瞬时值）
UntypedValue  // 不知道 or 不想承诺语义
*/

var (
	// sangfor_ad_vs_http_request_rate ( 请求速率 个/秒)
	SangforAdVirtualServiceHttpRequestRate = prometheus.NewDesc(
		"sangfor_ad_vs_http_request_rate",
		"Virtual service HTTP request rate.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_connection_count ( 并发连接 )
	SangforAdVirtualServiceConnectionCount = prometheus.NewDesc(
		"sangfor_ad_vs_connection_count",
		"Virtual service Concurrent connections count.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_connection_rate ( 新建连接 个/秒 )
	SangforAdVirtualServiceConnectionRate = prometheus.NewDesc(
		"sangfor_ad_vs_connection_rate",
		"Virtual service Number of new connections rate.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_upstream_throughput_rate ( 上行吞吐量速率 BIT/秒 )
	SangforAdVirtualServiceUpstreamThroughputBitRate = prometheus.NewDesc(
		"sangfor_ad_vs_upstream_throughput_bit_rate",
		"Virtual service Uplink throughput rate.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_downstream_throughput_rate ( 下行吞吐量速率 BIT/秒 )
	SangforAdVirtualServiceDownstreamThroughputBitRate = prometheus.NewDesc(
		"sangfor_ad_vs_downstream_throughput_bit_rate",
		"Virtual service Downlink throughput rate.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)
)

// Prometheus Describe接口传递指标描述符到 channel
func (m *MetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- SangforAdVirtualServiceHttpRequestRate
	ch <- SangforAdVirtualServiceConnectionCount
	ch <- SangforAdVirtualServiceConnectionRate
	ch <- SangforAdVirtualServiceUpstreamThroughputBitRate
	ch <- SangforAdVirtualServiceDownstreamThroughputBitRate
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
		// sangfor_ad_vs_http_request_rate
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceHttpRequestRate,
			prometheus.GaugeValue, float64(v.HttpRequestRate.Value), global.Config.SangforAd.DeviceName, v.Name, v.HttpRequestRate.Model, v.HttpRequestRate.Unit)

		// sangfor_ad_vs_connection_count
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceConnectionCount,
			prometheus.GaugeValue, float64(v.Connection.Value), global.Config.SangforAd.DeviceName, v.Name, v.Connection.Model, v.Connection.Unit)

		// sangfor_ad_vs_connection_rate
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceConnectionRate,
			prometheus.GaugeValue, float64(v.ConnectionRate.Value), global.Config.SangforAd.DeviceName, v.Name, v.ConnectionRate.Model, v.ConnectionRate.Unit)

		// sangfor_ad_vs_upstream_throughput_rate
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceUpstreamThroughputBitRate,
			prometheus.GaugeValue, float64(v.UpstreamThroughput.Value), global.Config.SangforAd.DeviceName, v.Name, v.UpstreamThroughput.Model, v.UpstreamThroughput.Unit)

		// sangfor_ad_vs_downstream_throughput_rate
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceDownstreamThroughputBitRate,
			prometheus.GaugeValue, float64(v.DownstreamThroughput.Value), global.Config.SangforAd.DeviceName, v.Name, v.DownstreamThroughput.Model, v.DownstreamThroughput.Unit)
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
