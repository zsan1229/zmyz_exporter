package basic

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

// 定义一个结构体来封装指标收集器
type basicMetrics struct {
	cpuUsage      prometheus.Gauge
	memUsage      prometheus.Gauge
	totalmemUsage prometheus.Gauge
	stgUsage      prometheus.Gauge
	totalstgUsage prometheus.Gauge
}

// 构造函数：创建并返回一个 basicMetrics 实例
func NewMetrics() *basicMetrics {
	return &basicMetrics{
		cpuUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "system_cpu_usage_percentage",
			Help: "CPU usage percentage.",
		}),
		memUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_usage_bytes",
			Help: "Memory usage bytes.",
		}),
		totalmemUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "system_total_memory_bytes",
			Help: "Memory total bytes.",
		}),
		stgUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "system_stg_usage_bytes",
			Help: "stg usage bytes.",
		}),
		totalstgUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "system_total_stg_bytes",
			Help: "stg total bytes.",
		}),
	}
}

// Collector 接口的 Describe 方法，用于描述所有指标
func (m *basicMetrics) Describe(ch chan<- *prometheus.Desc) {
	m.cpuUsage.Describe(ch)
	m.memUsage.Describe(ch)
	m.totalmemUsage.Describe(ch)
	m.stgUsage.Describe(ch)
	m.totalstgUsage.Describe(ch)
}

// Collector 接口的 Collect 方法，用于收集并更新指标
func (m *basicMetrics) Collect(ch chan<- prometheus.Metric) {
	// 获取 CPU 使用率
	cpuStats, err := cpu.Percent(time.Second, false)
	if err == nil && len(cpuStats) > 0 {
		m.cpuUsage.Set(cpuStats[0])
	}

	// 获取内存使用情况、总内存
	v, err := mem.VirtualMemory()
	if err == nil {
		m.memUsage.Set(float64(v.Used))
		m.totalmemUsage.Set(float64(v.Total))
	}

	// 获取硬盘使用情况
	diskStats, err := disk.Usage("D:\\")
	if err == nil {
		m.totalstgUsage.Set(float64(diskStats.Total))
		m.stgUsage.Set(float64(diskStats.Used))
	}

	// 将指标发送到 channel
	m.cpuUsage.Collect(ch)
	m.memUsage.Collect(ch)
	m.totalmemUsage.Collect(ch)
	m.stgUsage.Collect(ch)
	m.totalstgUsage.Collect(ch)

}
