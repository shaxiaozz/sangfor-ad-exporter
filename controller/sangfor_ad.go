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

	// sangfor_ad_vs_general_throughput_rate ( 上/下行总吞吐量速率 BIT/秒 )
	SangforAdVirtualServiceGeneralThroughputBitRate = prometheus.NewDesc(
		"sangfor_ad_vs_general_throughput_bit_rate",
		"Virtual service General throughput rate.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_client_connection_count ( 客户端连接数 )
	SangforAdVirtualServiceClientConnectionCount = prometheus.NewDesc(
		"sangfor_ad_vs_client_connection_count",
		"Virtual service Client connection count.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_server_connection_count ( 服务端连接数 )
	SangforAdVirtualServiceServerConnectionCount = prometheus.NewDesc(
		"sangfor_ad_vs_server_connection_count",
		"Virtual service Server connection count.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_connection_established_count ( Established并发连接数 )
	SangforAdVirtualServiceConnectionEstablishedCount = prometheus.NewDesc(
		"sangfor_ad_vs_connection_established_count",
		"Virtual service Established concurrent connections count.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_connection_established_count ( TCP连接池新建速率 )
	SangforAdVirtualServicePoolConnectionRateCount = prometheus.NewDesc(
		"sangfor_ad_vs_pool_connection_rate_count",
		"Virtual service TCP connection pool new creation rate count.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_ssl_connection_rate ( SSL新建连接数 个/秒)
	SangforAdVirtualServiceSslConnectionRate = prometheus.NewDesc(
		"sangfor_ad_vs_ssl_connection_rate",
		"Virtual service SSL new connection rate.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_ssl_connection_count ( SSL连接数 )
	SangforAdVirtualServiceSslConnectionCount = prometheus.NewDesc(
		"sangfor_ad_vs_ssl_connection_count",
		"Virtual service SSL connection count.",
		[]string{"device_name", "vs_name", "model", "unit"}, nil)

	// sangfor_ad_vs_health ( 健康状态 ）（NORMAL-正常/FAILURE-故障/ALERT-告警）
	SangforAdVirtualServiceHealth = prometheus.NewDesc(
		"sangfor_ad_vs_health",
		"Virtual service NORMAL/FAILURE/ALERT health status",
		[]string{"device_name", "vs_name", "status"}, nil)

	// sangfor_ad_vs_state ( 配置启/禁用开关 ）（ENABLE-启用/DISABLE-禁用）
	SangforAdVirtualServiceState = prometheus.NewDesc(
		"sangfor_ad_vs_state",
		"Virtual service enable/disable state",
		[]string{"device_name", "vs_name", "state"}, nil)
)

// Prometheus Describe接口传递指标描述符到 channel
func (m *MetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- SangforAdVirtualServiceHttpRequestRate
	ch <- SangforAdVirtualServiceConnectionCount
	ch <- SangforAdVirtualServiceConnectionRate
	ch <- SangforAdVirtualServiceUpstreamThroughputBitRate
	ch <- SangforAdVirtualServiceDownstreamThroughputBitRate
	ch <- SangforAdVirtualServiceClientConnectionCount
	ch <- SangforAdVirtualServiceConnectionEstablishedCount
	ch <- SangforAdVirtualServicePoolConnectionRateCount
	ch <- SangforAdVirtualServiceSslConnectionRate
	ch <- SangforAdVirtualServiceSslConnectionCount
	ch <- SangforAdVirtualServiceHealth
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

		// sangfor_ad_vs_upstream_throughput_bit_rate
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceUpstreamThroughputBitRate,
			prometheus.GaugeValue, float64(v.UpstreamThroughput.Value), global.Config.SangforAd.DeviceName, v.Name, v.UpstreamThroughput.Model, v.UpstreamThroughput.Unit)

		// sangfor_ad_vs_downstream_throughput_bit_rate
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceDownstreamThroughputBitRate,
			prometheus.GaugeValue, float64(v.DownstreamThroughput.Value), global.Config.SangforAd.DeviceName, v.Name, v.DownstreamThroughput.Model, v.DownstreamThroughput.Unit)

		// sangfor_ad_vs_general_throughput_bit_rate
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceGeneralThroughputBitRate,
			prometheus.GaugeValue, float64(v.GeneralThroughput.Value), global.Config.SangforAd.DeviceName, v.Name, v.GeneralThroughput.Model, v.GeneralThroughput.Unit)

		// sangfor_ad_vs_client_connection_count
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceClientConnectionCount,
			prometheus.GaugeValue, float64(v.ClientConnection.Value), global.Config.SangforAd.DeviceName, v.Name, v.ClientConnection.Model, v.ClientConnection.Unit)

		// sangfor_ad_vs_server_connection_count
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceServerConnectionCount,
			prometheus.GaugeValue, float64(v.ServerConnection.Value), global.Config.SangforAd.DeviceName, v.Name, v.ServerConnection.Model, v.ServerConnection.Unit)

		// sangfor_ad_vs_connection_established_count
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceConnectionEstablishedCount,
			prometheus.GaugeValue, float64(v.ConnectionEstablished.Value), global.Config.SangforAd.DeviceName, v.Name, v.ConnectionEstablished.Model, v.ConnectionEstablished.Unit)

		// sangfor_ad_vs_pool_connection_rate_count
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServicePoolConnectionRateCount,
			prometheus.GaugeValue, float64(v.PoolConnectionRate.Value), global.Config.SangforAd.DeviceName, v.Name, v.PoolConnectionRate.Model, v.PoolConnectionRate.Unit)

		// sangfor_ad_vs_ssl_connection_rate
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceSslConnectionRate,
			prometheus.GaugeValue, float64(v.SslConnectionRate.Value), global.Config.SangforAd.DeviceName, v.Name, v.SslConnectionRate.Model, v.SslConnectionRate.Unit)

		// sangfor_ad_vs_ssl_connection_count
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceSslConnectionCount,
			prometheus.GaugeValue, float64(v.SslConnection.Value), global.Config.SangforAd.DeviceName, v.Name, v.SslConnection.Model, v.SslConnection.Unit)

		// sangfor_ad_vs_health
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceHealth,
			prometheus.GaugeValue, 1, global.Config.SangforAd.DeviceName, v.Name, v.Health)

		// sangfor_ad_vs_state
		state := "DISABLE"
		if v.State == "ENABLE" {
			state = "ENABLE"
		}
		ch <- prometheus.MustNewConstMetric(SangforAdVirtualServiceState,
			prometheus.GaugeValue, 1, global.Config.SangforAd.DeviceName, v.Name, state)
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
