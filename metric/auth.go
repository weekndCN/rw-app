package metric

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/weekndCN/rw-app/core"
)

var noContext = context.Background()

// AuthCount auth user number
func AuthCount(auths core.AuthStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: "rwplus_user_count",
				Help: "Total number of register users",
			}, func() float64 {
				i, _ := auths.Count(noContext)
				return float64(i)
			}),
	)
}
