package tag

import (
	"cs-api/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type instrumenting struct {
	requestCount   *prometheus.CounterVec
	requestLatency *prometheus.HistogramVec
	h              ITagHandler
}

func NewInstrumenting(h ITagHandler) ITagHandler {
	labels := []string{"op"}

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: types.RequestCount,
		Help: types.RequestCountHelp,
	}, labels)

	latency := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: types.RequestLatency,
		Help: types.RequestLatencyHelp,
	}, labels)

	prometheus.MustRegister(counter, latency)

	return &instrumenting{
		requestCount:   counter,
		requestLatency: latency,
		h:              h,
	}
}

func (i *instrumenting) ListTag(c *gin.Context) {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("ListTag").Add(1)
		i.requestLatency.WithLabelValues("ListTag").Observe(time.Since(begin).Seconds())
	}(time.Now())

	i.h.ListTag(c)
}

func (i *instrumenting) CreateTag(c *gin.Context) {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("CreateTag").Add(1)
		i.requestLatency.WithLabelValues("CreateTag").Observe(time.Since(begin).Seconds())
	}(time.Now())

	i.h.CreateTag(c)
}

func (i *instrumenting) GetTag(c *gin.Context) {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("GetTag").Add(1)
		i.requestLatency.WithLabelValues("GetTag").Observe(time.Since(begin).Seconds())
	}(time.Now())

	i.h.GetTag(c)
}

func (i *instrumenting) UpdateTag(c *gin.Context) {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("UpdateTag").Add(1)
		i.requestLatency.WithLabelValues("UpdateTag").Observe(time.Since(begin).Seconds())
	}(time.Now())

	i.h.UpdateTag(c)
}

func (i *instrumenting) DeleteTag(c *gin.Context) {
	defer func(begin time.Time) {
		i.requestCount.WithLabelValues("DeleteTag").Add(1)
		i.requestLatency.WithLabelValues("DeleteTag").Observe(time.Since(begin).Seconds())
	}(time.Now())

	i.h.DeleteTag(c)
}
