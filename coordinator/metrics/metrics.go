package metrics

import (
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	metrics.Metrics
	IncNumTaskCreated()
	IncNumTaskCompleted()
}

// CoordinatorMetrics contains instrumented metrics that should be incremented by the avs node using the methods below
type CoordinatorMetrics struct {
	metrics.Metrics
	numTaskCreated prometheus.Counter
	// if numTaskCompleted != numTaskCreated, then there is a bug
	numTaskCompleted prometheus.Counter
}

func NewCoordinatorMetrics(avsName string, coordinatorMetrics *metrics.EigenMetrics, reg prometheus.Registerer) *CoordinatorMetrics {
	return &CoordinatorMetrics{
		Metrics: coordinatorMetrics,
		numTaskCreated: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: avsName,
				Name:      "num_task_created",
				Help:      "The number of tasks send to contract",
			}),
		numTaskCompleted: promauto.With(reg).NewCounter(
			prometheus.CounterOpts{
				Namespace: avsName,
				Name:      "num_task_completed",
				Help:      "The number of task response to contract",
			}),
	}
}

func (m *CoordinatorMetrics) IncNumTaskCreated() {
	m.numTaskCreated.Inc()
}

func (m *CoordinatorMetrics) IncNumTaskCompleted() {
	m.numTaskCompleted.Inc()
}
