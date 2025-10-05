package main

import (
	"fmt"
	"log"
	"net/http"
	"zmyz_exporter/basic"
	"zmyz_exporter/network"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// "zmyz_exporter/utils"
)

func Registry() {
	// 创建一个 Prometheus 注册表
	reg := prometheus.NewRegistry()

	// 创建并注册自定义指标收集器
	basicExporterMetrics := basic.NewMetrics()
	pingCollector := network.NewPingCollector()
	reg.MustRegister(basicExporterMetrics, pingCollector)

	// 设置 HTTP 处理程序来暴露 /metrics 端点
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		// 禁用 Go 运行时指标，我们只关心自定义指标
		DisableCompression: true,
	}))

	// 启动 HTTP 服务器
	port := 9100
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting Go Node Exporter on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	Registry()
	//utils.ReadNetWorkConfig()
}
