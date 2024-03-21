package fetchers_test

import (
	"testing"

	"code.cloudfoundry.org/go-log-cache/rpc/logcache_v1"
	"code.cloudfoundry.org/lager/v3"
	"code.cloudfoundry.org/lager/v3/lagertest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFetchers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fetchers Suite")
}

var (
	logger lager.Logger
)

var _ = BeforeSuite(func() {
	logger = lagertest.NewTestLogger("cumulative-usage-test")
})

type Metric struct {
	Usage       float64
	Entitlement float64
	Age         float64
}

func queryResult(samples ...*logcache_v1.PromQL_Sample) *logcache_v1.PromQL_InstantQueryResult {
	return &logcache_v1.PromQL_InstantQueryResult{
		Result: &logcache_v1.PromQL_InstantQueryResult_Vector{
			Vector: &logcache_v1.PromQL_Vector{
				Samples: samples,
			},
		},
	}
}

func sample(instanceID, procInstanceID string, point *logcache_v1.PromQL_Point) *logcache_v1.PromQL_Sample {
	return &logcache_v1.PromQL_Sample{
		Metric: map[string]string{
			"instance_id":         instanceID,
			"process_instance_id": procInstanceID,
		},
		Point: point,
	}
}

func point(time string, value float64) *logcache_v1.PromQL_Point {
	return &logcache_v1.PromQL_Point{Time: time, Value: value}
}
