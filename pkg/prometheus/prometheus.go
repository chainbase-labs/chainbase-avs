package prometheus

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(AvsInfo)

}

var (
	AvsInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avs_info",
			Help: "AVS Info include: ip, hex_address, worker_memory_total_bytes, cpu model...",
		},
		[]string{"ip", "hex_address", "model", "worker_memory_total_bytes", "job_manager_status"},
	)

	MemoryTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "worker_memory_total_bytes",
			Help: "Total memory of the worker in bytes",
		},
	)
)

func UpdateHostMetrics(avsAddr string) {

	ip := GetOutboundIP()
	cpuModel := GetCPUModel()
	memTotal := GetTotalMemory()
	flinkJobManagerStatus := GetFlinkJobManagerStatus()
	AvsInfo.WithLabelValues(ip, avsAddr, cpuModel, fmt.Sprintf("%.2f", memTotal), flinkJobManagerStatus).Set(1)
}

func GetFlinkJobManagerStatus() string {
	resp, err := http.Get("http://flink-jobmanager:8081/config")
	if err != nil {

		slog.Error("failed to get flink job manager status", "error", err)
		return "0"
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		slog.Error("failed to get flink job manager status", "status_code", resp.StatusCode)
		return "0"

	}
	return "1"
}
