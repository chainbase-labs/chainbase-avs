package metrics

import (
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	metrics.Metrics
	IncNumTasksReceived()
	IncNumTasksAcceptedByCoordinator()
	// This metric would either need to be tracked by the coordinator itself,
	// or we would need to write a collector that queries onchain for this info
	// AddPercentageStakeSigned(percentage float64)
}

// AvsAndEigenMetrics contains instrumented metrics that should be incremented by the avs node using the methods below
type AvsAndEigenMetrics struct {
	metrics.Metrics
	numTasksReceived prometheus.Counter
	// if numSignedTaskResponsesAcceptedByCoordinator != numTasksReceived, then there is a bug
	numSignedTaskResponsesAcceptedByCoordinator prometheus.Counter
}

const chainbaseNamespace = "chainbase"

func NewAvsAndEigenMetrics(avsName string, eigenMetrics *metrics.EigenMetrics, reg prometheus.Registerer) *AvsAndEigenMetrics {
	return &AvsAndEigenMetrics{
		Metrics: eigenMetrics,
		numTasksReceived: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: chainbaseNamespace,
				Name:      "num_tasks_received",
				Help:      "The number of tasks received by reading from the avs service manager contract",
			}),
		numSignedTaskResponsesAcceptedByCoordinator: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: chainbaseNamespace,
				Name:      "num_signed_task_responses_accepted_by_coordinator",
				Help:      "The number of signed task responses accepted by the coordinator",
			}),
	}
}

func (m *AvsAndEigenMetrics) IncNumTasksReceived() {
	m.numTasksReceived.Inc()
}

func (m *AvsAndEigenMetrics) IncNumTasksAcceptedByCoordinator() {
	m.numSignedTaskResponsesAcceptedByCoordinator.Inc()
}
