package basic

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// 定义结构体封装所有指标描述符
type basicMetrics struct {
	cpuUsageDesc *prometheus.Desc
	memUsageDesc *prometheus.Desc
	totalMemDesc *prometheus.Desc
	stgUsageDesc *prometheus.Desc
	totalStgDesc *prometheus.Desc
}

// 构造函数：创建指标描述符
func NewMetrics() *basicMetrics {
	return &basicMetrics{
		cpuUsageDesc: prometheus.NewDesc(
			"system_cpu_usage_percentage",
			"CPU usage percentage.",
			nil, nil,
		),
		memUsageDesc: prometheus.NewDesc(
			"system_memory_usage_bytes",
			"Memory usage bytes.",
			nil, nil,
		),
		totalMemDesc: prometheus.NewDesc(
			"system_total_memory_bytes",
			"Memory total bytes.",
			nil, nil,
		),
		stgUsageDesc: prometheus.NewDesc(
			"system_stg_usage_bytes",
			"Storage usage bytes.",
			nil, nil,
		),
		totalStgDesc: prometheus.NewDesc(
			"system_total_stg_bytes",
			"Storage total bytes.",
			nil, nil,
		),
	}
}

// Describe 实现 prometheus.Collector 接口
func (m *basicMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.cpuUsageDesc
	ch <- m.memUsageDesc
	ch <- m.totalMemDesc
	ch <- m.stgUsageDesc
	ch <- m.totalStgDesc
}

// Collect 实现 prometheus.Collector 接口
func (m *basicMetrics) Collect(ch chan<- prometheus.Metric) {
	// 获取 CPU 使用率
	if cpuStats, err := cpu.Percent(time.Second, false); err == nil && len(cpuStats) > 0 {
		ch <- prometheus.MustNewConstMetric(
			m.cpuUsageDesc,
			prometheus.GaugeValue,
			cpuStats[0],
		)
	}

	// 获取内存使用情况
	if v, err := mem.VirtualMemory(); err == nil {
		ch <- prometheus.MustNewConstMetric(
			m.memUsageDesc,
			prometheus.GaugeValue,
			float64(v.Used),
		)
		ch <- prometheus.MustNewConstMetric(
			m.totalMemDesc,
			prometheus.GaugeValue,
			float64(v.Total),
		)
	}

	// 获取硬盘使用情况
	if diskStats, err := disk.Usage("/"); err == nil {
		ch <- prometheus.MustNewConstMetric(
			m.stgUsageDesc,
			prometheus.GaugeValue,
			float64(diskStats.Used),
		)
		ch <- prometheus.MustNewConstMetric(
			m.totalStgDesc,
			prometheus.GaugeValue,
			float64(diskStats.Total),
		)
	}
}
