package model

// 登录请求结构体
type SangforAdLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录请求响应结构体
type SangforAdLoginResp struct {
	Timeout          int    `json:"timeout"`
	CreateTimestamp  int    `json:"create_timestamp"`
	ExpiredTimestamp int    `json:"expired_timestamp"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	PermitCtl        []struct {
		Role    string `json:"role"`
		Project string `json:"project"`
	} `json:"permit_ctl"`
}

// 获取虚拟服务状态信息响应结构体
type SangforAdVirtualServiceStatResp struct {
	Items []struct {
		Name           string `json:"name"`
		ConnectionRate struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"connection_rate"`
		Connection struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"connection"`
		ConnectionEstablished struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"connection_established"`
		HttpRequestRate struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_request_rate"`
		ClientConnection struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"client_connection"`
		ServerConnection struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"server_connection"`
		PoolConnectionRate struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"pool_connection_rate"`
		UpstreamThroughput struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"upstream_throughput"`
		DownstreamThroughput struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"downstream_throughput"`
		GeneralThroughput struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"general_throughput"`
		SslConnectionRate struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"ssl_connection_rate"`
		SslConnection struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"ssl_connection"`
		HttpCompressionRawThroughput struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_compression_raw_throughput"`
		HttpCompressionThroughput struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_compression_throughput"`
		SslUpstreamThroughput struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"ssl_upstream_throughput"`
		SslDownstreamThroughput struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"ssl_downstream_throughput"`
		Health            string `json:"health"`
		State             string `json:"state"`
		MaximumConnection struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"maximum_connection"`
		TotalConnection struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"total_connection"`
		TotalHttpRequest struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"total_http_request"`
		UpstreamData struct {
			Model     string `json:"model"`
			Value     int64  `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"upstream_data"`
		DownstreamData struct {
			Model     string `json:"model"`
			Value     int64  `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"downstream_data"`
		UpstreamPacket struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"upstream_packet"`
		DownstreamPacket struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"downstream_packet"`
		PoolConnectionUsage struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"pool_connection_usage"`
		HttpCacheHit struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_cache_hit"`
		HttpCacheResponseData struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_cache_response_data"`
		HttpCompressionDataReduction struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_compression_data_reduction"`
		HttpCompressionSavingRatio struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_compression_saving_ratio"`
		Pool struct {
			Total  int `json:"total"`
			Health struct {
				Normal  []string      `json:"normal"`
				Failure []interface{} `json:"failure"`
				Busy    []interface{} `json:"busy"`
				Alert   []interface{} `json:"alert"`
			} `json:"health"`
		} `json:"pool"`
		VsHealthReason    string `json:"vs_health_reason"`
		HttpCacheCapacity struct {
			Model     string `json:"model"`
			Value     int64  `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_cache_capacity"`
		HttpCacheUsed struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_cache_used"`
		HttpCacheSingleUsed struct {
			Model     string `json:"model"`
			Value     int    `json:"value"`
			Timestamp int    `json:"timestamp"`
			Unit      string `json:"unit"`
		} `json:"http_cache_single_used"`
		EffectiveDevice string   `json:"effective_device"`
		Service         string   `json:"service"`
		Icon            string   `json:"icon"`
		Netns           string   `json:"netns"`
		Vips            []string `json:"vips"`
		Vports          []string `json:"vports"`
		Description     string   `json:"description"`
	} `json:"items"`
	TotalPages  int `json:"total_pages"`
	PageNumber  int `json:"page_number"`
	PageSize    int `json:"page_size"`
	TotalItems  int `json:"total_items"`
	ItemsOffset int `json:"items_offset"`
	ItemsLength int `json:"items_length"`
}
