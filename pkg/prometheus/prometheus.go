package prometheus

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(AvsInfo)
	prometheus.MustRegister(MemoryTotal)
	prometheus.MustRegister(JobManagerStatus)

}

var (
	JOB_NAME = "chainbase-avs"

	AvsInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avs_info",
			Help: "AVS Info include: ip, hex_address, cpu model...",
		},
		[]string{"ip", "hex_address", "model", "job"},
	)
	MemoryTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "worker_memory_total_bytes",
			Help: "Total memory of the worker in bytes",
		},
		[]string{"hex_address"},
	)

	JobManagerStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avs_job_manager_status",
			Help: "Status of the AVS job manager (0 for inactive, 1 for active)",
		},
		[]string{"hex_address"},
	)
)

func UpdateHostMetrics(avsAddr string) {

	ip := GetOutboundIP()
	cpuModel := GetCPUModel()
	memTotal := GetTotalMemory()
	flinkJobManagerStatus := GetFlinkJobManagerStatus()
	AvsInfo.WithLabelValues(ip, avsAddr, cpuModel, JOB_NAME).Set(1)
	MemoryTotal.WithLabelValues(avsAddr).Set(memTotal)
	JobManagerStatus.WithLabelValues(avsAddr).Set(flinkJobManagerStatus)
}

func GetFlinkJobManagerStatus() float64 {

	flinkHostName := os.Getenv("FLINK_CONNECT_ADDRESS")
	flinkHostPort := os.Getenv("FLINK_JOBMANAGER_PORT")
	resp, err := http.Get((fmt.Sprintf("http://%s:%s/config", flinkHostName, flinkHostPort)))

	if err != nil {

		slog.Error("failed to get flink job manager status", "error", err)
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		slog.Error("failed to get flink job manager status", "status_code", resp.StatusCode)
		return 0

	}

	return 1
}
