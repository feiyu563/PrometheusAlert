package model

import "github.com/prometheus/client_golang/prometheus"

var (
	AlertsFromCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "alers_from_count",
		Help: "count alers from any where",
	},
		[]string{"from", "message", "level", "host", "index"},
	)
	//model.AlertsFromCounter.WithLabelValues("from","to","message","level","host","index").Add(1)
	AlertToCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "alers_to_count",
		Help: "count alers to any where",
	},
		[]string{"to", "message", "phone"},
	)
	//model.AlertToCounter.WithLabelValues("to","message").Add(1)
	AlertFailedCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "alers_send_failed_count",
		Help: "count alers send failed",
	},
		[]string{"to", "message", "phone"},
	)
	//model.AlertFailedCounter.WithLabelValues("to","message","phone").Add(1)
)

func MetricsInit() {
	prometheus.MustRegister(AlertsFromCounter)
	prometheus.MustRegister(AlertToCounter)
	prometheus.MustRegister(AlertFailedCounter)
}
