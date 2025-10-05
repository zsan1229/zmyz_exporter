package network

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"zmyz_exporter/utils"

	"github.com/prometheus/client_golang/prometheus"
)

func PingIP() []map[string]string {
	IPS := utils.ReadNetWorkConfig()
	var cmd *exec.Cmd
	var match []string
	var res []map[string]string
	for _, v := range IPS {
		cmd = exec.Command("ping", "-c", "1", "-W", "2", v.Value)
		output, err := cmd.CombinedOutput()
		outStr := string(output)
		// fmt.Println(outStr)
		// 匹配 time=后面的数字 (可以包含小数)
		re := regexp.MustCompile(`time=([\d.]+)`)
		match = re.FindStringSubmatch(outStr)
		var time string
		if len(match) <= 1 || err != nil {
			time = "0"
			fmt.Println("No match or Time out")
		} else {
			time = match[1]
			fmt.Println("Time:", match[1]) // 输出 189
		}
		res = append(res, map[string]string{v.Name: time})
	}
	fmt.Println(res[0]["上海联通"])
	return res
}

type PingCollector struct {
	latencyDesc *prometheus.Desc
}

func NewPingCollector() *PingCollector {
	return &PingCollector{
		latencyDesc: prometheus.NewDesc(
			"ping_latency_ms",
			"Ping latency time in milliseconds",
			[]string{"target"},
			nil,
		),
	}
}

// Describe 描述指标
func (c *PingCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.latencyDesc
}

// Collect 收集指标
func (c *PingCollector) Collect(ch chan<- prometheus.Metric) {
	results := PingIP()

	for _, item := range results {
		for name, val := range item {
			value, err := strconv.ParseFloat(val, 64)
			if err != nil {
				value = 0
			}
			ch <- prometheus.MustNewConstMetric(
				c.latencyDesc,
				prometheus.GaugeValue,
				value,
				name,
			)
		}
	}
}
