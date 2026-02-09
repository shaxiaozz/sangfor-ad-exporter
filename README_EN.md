# Sangfor AD Exporter

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/shaxiaozz/sangfor-ad-exporter)](https://goreportcard.com/report/github.com/shaxiaozz/sangfor-ad-exporter)

**Sangfor AD Exporter** is a Prometheus Exporter for collecting metrics from Sangfor Application Delivery (AD) devices. If it is inconvenient to use the SNMP protocol to scrape monitoring data, you can use this exporter to obtain the running status and performance metrics of Virtual Services via API.

## ✨ Features

- **Comprehensive Metric Collection**: Supports collecting core metrics such as HTTP request rate, concurrent connections, throughput (upstream/downstream/total), and SSL connections of virtual services.
- **Status Monitoring**: Real-time monitoring of virtual service health status (NORMAL/FAILURE/ALERT) and enable/disable status.
- **High Performance**: Developed in Go, lightweight and efficient.
- **Token Caching**: Built-in Token caching mechanism to reduce the frequency of authentication requests to AD devices.
- **Docker Support**: Provides Docker images for easy containerized deployment.

## 🚀 Quick Start

### Method 1: Deploy using Docker (Recommended)

1. **Prepare Configuration File**

   Create a `config.yaml` file and fill in the connection information for Sangfor AD:

   ```yaml
   # config.yaml
   sangfor_ad:
     device_name: "INFRA-AD-01"  # Custom device identifier
     username: "admin"           # Login username
     password: "your_password"   # Login password
     url: "https://192.168.1.1"  # Sangfor AD management address
   ```

2. **Start Container**

   ```bash
   docker run -d \
     --name sangfor-ad-exporter \
     -p 9098:9098 \
     -v $(pwd)/config.yaml:/sangfor-ad-exporter/config.yaml \
     --restart=always \
     shaxiaozz/sangfor-ad-exporter:latest
   ```
   *(Note: Please build the image yourself or use an existing one)*

### Method 2: Run Binary

1. **Download and Compile**

   ```bash
   git clone https://github.com/shaxiaozz/sangfor-ad-exporter.git
   cd sangfor-ad-exporter
   go build -o sangfor-ad-exporter
   ```

2. **Configure**

   Create a `config/config.yaml` file in the running directory (refer to `config/config-example.yaml`).

3. **Run**

   ```bash
   ./sangfor-ad-exporter start
   ```

## ⚙️ Configuration

The configuration file is located at `config/config.yaml` by default and supports the following options:

```yaml
sangfor_ad:
  device_name: "INFRA-AD-01"  # device_name label value in exported metrics
  username: "admin"           # Recommend creating a read-only API account
  password: "your_password"
  url: "https://192.168.1.1"  # Must include protocol (http/https)
```

## 📊 Exported Metrics

The Exporter exposes the following metrics at the `/metrics` path. All metrics include labels such as `device_name`, `vs_name` (Virtual Service Name), `model`, `unit`, etc.

| Metric Name | Type | Description |
| :--- | :--- | :--- |
| `sangfor_ad_vs_http_request_rate` | Gauge | Virtual Service HTTP Request Rate (requests/sec) |
| `sangfor_ad_vs_connection_count` | Gauge | Virtual Service Concurrent Connections |
| `sangfor_ad_vs_connection_rate` | Gauge | Virtual Service New Connection Rate (connections/sec) |
| `sangfor_ad_vs_upstream_throughput_bit_rate` | Gauge | Virtual Service Upstream Throughput Rate (bps) |
| `sangfor_ad_vs_downstream_throughput_bit_rate` | Gauge | Virtual Service Downstream Throughput Rate (bps) |
| `sangfor_ad_vs_general_throughput_bit_rate` | Gauge | Virtual Service Total Throughput Rate (bps) |
| `sangfor_ad_vs_client_connection_count` | Gauge | Client Connection Count |
| `sangfor_ad_vs_server_connection_count` | Gauge | Server Connection Count |
| `sangfor_ad_vs_connection_established_count` | Gauge | Concurrent Connections in Established State |
| `sangfor_ad_vs_pool_connection_rate_count` | Gauge | TCP Connection Pool New Creation Rate |
| `sangfor_ad_vs_ssl_connection_rate` | Gauge | SSL New Connection Rate (connections/sec) |
| `sangfor_ad_vs_ssl_connection_count` | Gauge | SSL Connection Count |
| `sangfor_ad_vs_health` | Gauge | Virtual Service Health Status (1=Current State) <br> Label `status`: NORMAL, FAILURE, ALERT |
| `sangfor_ad_vs_state` | Gauge | Virtual Service Enable State (1=Current State) <br> Label `state`: ENABLE, DISABLE |

## 🔌 Prometheus Configuration

Add the following job to `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: 'sangfor_ad'
    static_configs:
      - targets: ['localhost:9098']
    metrics_path: /metrics
    scrape_interval: 15s
```

## 🛠️ Development & Build

**Run Locally**
```bash
go run main.go start
```

**Build Docker Image**
```bash
docker build -t sangfor-ad-exporter:latest .
```

## 📄 License

This project is licensed under the [Apache-2.0](LICENSE) License.
