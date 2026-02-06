# Sangfor AD Exporter

[English](README_EN.md) | 中文文档

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/shaxiaozz/sangfor-ad-exporter)](https://goreportcard.com/report/github.com/shaxiaozz/sangfor-ad-exporter)

**Sangfor AD Exporter** 是一个用于采集深信服应用交付（Sangfor AD）设备指标的 Prometheus Exporter。如果不方便使用 SNMP 协议抓取监控数据，可以使用本 Exporter 通过 API 方式获取虚拟服务（Virtual Service）的运行状态和性能指标。

## ✨ 功能特性

- **全面的指标采集**：支持采集虚拟服务的 HTTP 请求速率、并发连接数、吞吐量（上行/下行/总计）、SSL 连接数等核心指标。
- **状态监控**：实时监控虚拟服务的健康状态（正常/故障/告警）和启禁用状态。
- **高性能**：基于 Go 语言开发，轻量级且高效。
- **Token 缓存**：内置 Token 缓存机制，减少对 AD 设备的认证请求频率。
- **Docker 支持**：提供 Docker 镜像，方便容器化部署。

## 🚀 快速开始

### 方式一：使用 Docker 部署（推荐）

1. **准备配置文件**

   创建 `config.yaml` 文件，并填写 Sangfor AD 的连接信息：

   ```yaml
   # config.yaml
   sangfor_ad:
     device_name: "INFRA-AD-01"  # 自定义设备标识
     username: "admin"           # 登录用户名
     password: "your_password"   # 登录密码
     url: "https://192.168.1.1"  # Sangfor AD 管理地址
   ```

2. **启动容器**

   ```bash
   docker run -d \
     --name sangfor-ad-exporter \
     -p 8080:8080 \
     -v $(pwd)/config.yaml:/sangfor-ad-exporter/config.yaml \
     --restart=always \
     shaxiaozz/sangfor-ad-exporter:latest
   ```
   *(注意：请先自行构建镜像或使用已有的镜像)*

### 方式二：二进制运行

1. **下载并编译**

   ```bash
   git clone https://github.com/shaxiaozz/sangfor-ad-exporter.git
   cd sangfor-ad-exporter
   go build -o sangfor-ad-exporter
   ```

2. **配置**

   在运行目录下创建 `config/config.yaml` 文件（参考 `config/config-example.yaml`）。

3. **运行**

   ```bash
   ./sangfor-ad-exporter start
   ```

## ⚙️ 配置说明

配置文件默认位于 `config/config.yaml`，支持以下配置项：

```yaml
sangfor_ad:
  device_name: "INFRA-AD-01"  # 自定义设备标识
  username: "admin"           # 登录用户名
  password: "your_password"   # 登录密码
  url: "https://192.168.1.1"  # Sangfor AD 管理地址
```

## 📊 导出指标 (Metrics)

Exporter 会在 `/metrics` 路径下暴露以下指标。所有指标均包含 `device_name`, `vs_name` (虚拟服务名称), `model` (模式), `unit` (单位) 等标签。

| 指标名称 (Metric Name) | 类型 | 描述 |
| :--- | :--- | :--- |
| `sangfor_ad_vs_http_request_rate` | Gauge | 虚拟服务 HTTP 请求速率 (个/秒) |
| `sangfor_ad_vs_connection_count` | Gauge | 虚拟服务并发连接数 |
| `sangfor_ad_vs_connection_rate` | Gauge | 虚拟服务新建连接速率 (个/秒) |
| `sangfor_ad_vs_upstream_throughput_bit_rate` | Gauge | 虚拟服务上行吞吐速率 (bps) |
| `sangfor_ad_vs_downstream_throughput_bit_rate` | Gauge | 虚拟服务下行吞吐速率 (bps) |
| `sangfor_ad_vs_general_throughput_bit_rate` | Gauge | 虚拟服务总吞吐速率 (bps) |
| `sangfor_ad_vs_client_connection_count` | Gauge | 客户端连接数 |
| `sangfor_ad_vs_server_connection_count` | Gauge | 服务端连接数 |
| `sangfor_ad_vs_connection_established_count` | Gauge | Established 状态的并发连接数 |
| `sangfor_ad_vs_pool_connection_rate_count` | Gauge | TCP 连接池新建速率 |
| `sangfor_ad_vs_ssl_connection_rate` | Gauge | SSL 新建连接速率 (个/秒) |
| `sangfor_ad_vs_ssl_connection_count` | Gauge | SSL 连接数 |
| `sangfor_ad_vs_health` | Gauge | 虚拟服务健康状态 (1=当前状态) <br> 标签 `status`: NORMAL (正常), FAILURE (故障), ALERT (告警) |
| `sangfor_ad_vs_state` | Gauge | 虚拟服务启用状态 (1=当前状态) <br> 标签 `state`: ENABLE (启用), DISABLE (禁用) |

## 🔌 Prometheus 配置

在 `prometheus.yml` 中添加如下 job：

```yaml
scrape_configs:
  - job_name: 'sangfor_ad'
    static_configs:
      - targets: ['localhost:8080']
    metrics_path: /metrics
    scrape_interval: 15s
```

## 🛠️ 开发与构建

**本地运行**
```bash
go run main.go start
```

**构建 Docker 镜像**
```bash
docker build -t sangfor-ad-exporter:latest .
```

## 📄 License

本项目采用 [Apache-2.0](LICENSE) 开源许可证。
