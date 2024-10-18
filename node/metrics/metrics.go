package metrics

import (
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	metrics.Metrics
	IncNumTaskReceived()
	IncNumTaskSucceed()
	IncNumTaskFailed()
	SetTaskExecutionTime(executionTime float64)
	UpdateNodeMetrics(operatorAddress, jobName, jobManagerHost, jobManagerPort string) (string, float64)
}

// AvsAndEigenMetrics contains instrumented metrics that should be incremented by the avs node using the methods below
type AvsAndEigenMetrics struct {
	metrics.Metrics
	// local metrics
	numTaskReceived prometheus.Counter
	// if numTaskSucceed != numTaskReceived, then there is a bug
	numTaskSucceed    prometheus.Counter
	numTaskFailed     prometheus.Counter
	taskExecutionTime prometheus.Gauge
	// remote metrics
	nodeInfo         *prometheus.GaugeVec
	memoryTotal      *prometheus.GaugeVec
	jobManagerStatus *prometheus.GaugeVec
}

func NewAvsAndEigenMetrics(avsName string, eigenMetrics *metrics.EigenMetrics, reg prometheus.Registerer) *AvsAndEigenMetrics {
	return &AvsAndEigenMetrics{
		Metrics: eigenMetrics,
		numTaskReceived: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: avsName,
				Name:      "num_task_received",
				Help:      "The number of tasks received from coordinator",
			}),
		numTaskSucceed: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: avsName,
				Name:      "num_task_succeed",
				Help:      "The number of signed task responses accepted by the coordinator",
			}),
		numTaskFailed: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: avsName,
				Name:      "num_task_failed",
				Help:      "The number of task execute failed",
			}),
		taskExecutionTime: promauto.With(reg).NewGauge(
			prometheus.GaugeOpts{
				Namespace: avsName,
				Name:      "task_execution_time_minutes",
				Help:      "Task execution time in minutes",
			}),
		nodeInfo: promauto.With(reg).NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: avsName,
				Name:      "node_info",
				Help:      "Node Info include: ip, hex_address ...",
			},
			[]string{"hex_address", "ip", "job"},
		),
		memoryTotal: promauto.With(reg).NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: avsName,
				Name:      "node_memory_total_bytes",
				Help:      "Total memory of the node in bytes",
			},
			[]string{"hex_address", "job"},
		),
		jobManagerStatus: promauto.With(reg).NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: avsName,
				Name:      "node_job_manager_status",
				Help:      "Status of the job manager (0 for inactive, 1 for active)",
			},
			[]string{"hex_address", "job"},
		),
	}
}

func (m *AvsAndEigenMetrics) IncNumTaskReceived() {
	m.numTaskReceived.Inc()
}

func (m *AvsAndEigenMetrics) IncNumTaskSucceed() {
	m.numTaskSucceed.Inc()
}

func (m *AvsAndEigenMetrics) IncNumTaskFailed() {
	m.numTaskFailed.Inc()
}

func (m *AvsAndEigenMetrics) SetTaskExecutionTime(executionTime float64) {
	m.taskExecutionTime.Set(executionTime)
}

func (m *AvsAndEigenMetrics) UpdateNodeMetrics(operatorAddress, jobName, jobManagerHost, jobManagerPort string) (string, float64) {
	ip := GetOutboundIP()
	memTotal := GetTotalMemory()
	jobManagerStatus := GetJobManagerStatus(jobManagerHost, jobManagerPort)
	m.nodeInfo.WithLabelValues(ip, operatorAddress, jobName).Set(1)
	m.memoryTotal.WithLabelValues(operatorAddress, jobName).Set(memTotal)
	m.jobManagerStatus.WithLabelValues(operatorAddress, jobName).Set(jobManagerStatus)
	return ip, jobManagerStatus
}
