package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(AvsInfo)
	prometheus.MustRegister(MemoryTotal)
}

var (
	AvsInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avs_info",
			Help: "AVS Info include: ip, hex_address, cpu model...",
		},
		[]string{"ip", "hex_address", "model"},
	)

	MemoryTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "worker_memory_total_bytes",
			Help: "Total memory of the worker in bytes",
		},
	)
)

func UpdateHostMetrics(avsAddr string) {
	// 更新指标
	ip := GetOutboundIP()
	cpuModel := GetCPUModel()
	AvsInfo.WithLabelValues(ip, avsAddr, cpuModel).Set(1)

	memTotal := GetTotalMemory()
	MemoryTotal.Set(memTotal)

}
