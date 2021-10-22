package bardmetric

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"time"
)

// Customized what I want from https://github.com/zsais/go-gin-prometheus/blob/master/middleware.go

var defaultMetricPath = "/metrics"

// Standard default metrics
//	counter, counter_vec, gauge, gauge_vec,
//	histogram, histogram_vec, summary, summary_vec
var reqCnt = &Metric{
	ID:          "reqCnt",
	Name:        "requests_total",
	Description: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	Type:        "counter_vec",
	Args:        []string{"code", "method", "handler", "host", "url"}}

var reqDur = &Metric{
	ID:          "reqDur",
	Name:        "request_duration_seconds",
	Description: "The HTTP request latencies in seconds.",
	Type:        "histogram_vec",
	Args:        []string{"code", "method", "url"},
}

var resSz = &Metric{
	ID:          "resSz",
	Name:        "response_size_bytes",
	Description: "The HTTP response sizes in bytes.",
	Type:        "summary"}

var reqSz = &Metric{
	ID:          "reqSz",
	Name:        "request_size_bytes",
	Description: "The HTTP request sizes in bytes.",
	Type:        "summary"}

var standardMetrics = []*Metric{
	reqCnt,
	reqDur,
	resSz,
	reqSz,
}

type Metric struct {
	MetricCollector prometheus.Collector
	ID              string
	Name            string
	Description     string
	Type            string
	Args            []string
}

type GinPrometheus struct {
	requestCount              *prometheus.CounterVec
	requestDuration           *prometheus.HistogramVec
	requestSize, responseSize prometheus.Summary
	router                    *gin.Engine
	listenAddress             string

	MetricsList []*Metric
	MetricsPath string
}

// NewPrometheus generates a new set of metrics with a certain subsystem name
func NewPrometheus(subsystem string, customMetricsList ...[]*Metric) *GinPrometheus {

	var metricsList []*Metric

	if len(customMetricsList) > 1 {
		panic("Too many args. NewPrometheus( string, <optional []*Metric> ).")
	} else if len(customMetricsList) == 1 {
		metricsList = customMetricsList[0]
	}

	for _, metric := range standardMetrics {
		metricsList = append(metricsList, metric)
	}

	p := &GinPrometheus{
		MetricsList: metricsList,
		MetricsPath: defaultMetricPath,
	}

	p.registerMetrics(subsystem)

	return p
}

func (g *GinPrometheus) MustRegister(c prometheus.Collector) {
	prometheus.MustRegister(c)
}

func NewMetric(m *Metric, subsystem string) prometheus.Collector {
	var metric prometheus.Collector
	switch m.Type {
	case "counter_vec":
		metric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "counter":
		metric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "gauge_vec":
		metric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "gauge":
		metric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "histogram_vec":
		metric = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "histogram":
		metric = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "summary_vec":
		metric = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "summary":
		metric = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	}
	return metric
}

func (p *GinPrometheus) SetMetricsPath(e *gin.Engine) {

	if p.listenAddress != "" {
		p.router.GET(p.MetricsPath, prometheusHandler())
		p.runServer()
	} else {
		e.GET(p.MetricsPath, prometheusHandler())
	}
}

func (p *GinPrometheus) runServer() {
	if p.listenAddress != "" {
		go p.router.Run(p.listenAddress)
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (p *GinPrometheus) registerMetrics(subsystem string) {
	for _, metricDef := range p.MetricsList {
		metric := NewMetric(metricDef, subsystem)
		if err := prometheus.Register(metric); err != nil {
			log.Logger.Err(err).
				Str("metric_name", metricDef.Name).
				Msg("Metric could not be registered in Prometheus")
		}
		switch metricDef {
		case reqCnt:
			p.requestCount = metric.(*prometheus.CounterVec)
		case reqDur:
			p.requestDuration = metric.(*prometheus.HistogramVec)
		case resSz:
			p.responseSize = metric.(prometheus.Summary)
		case reqSz:
			p.requestSize = metric.(prometheus.Summary)
		}
		metricDef.MetricCollector = metric
	}
}

func (p *GinPrometheus) Use(e *gin.Engine) {
	e.Use(p.HandlerFunc())
	p.SetMetricsPath(e)
}

func (p *GinPrometheus) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == p.MetricsPath {
			c.Next()
			return
		}

		start := time.Now()
		reqSz := computeApproximateRequestSize(c.Request)

		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := time.Since(start).Seconds()
		responseSize := float64(c.Writer.Size())

		url := c.FullPath()
		p.requestDuration.WithLabelValues(status, c.Request.Method, url).Observe(elapsed)
		p.requestCount.WithLabelValues(status, c.Request.Method, c.HandlerName(), c.Request.Host, url).Inc()
		p.requestSize.Observe(float64(reqSz))
		p.responseSize.Observe(responseSize)
	}
}

// From https://github.com/DanielHeckrath/gin-prometheus/blob/master/gin_prometheus.go
func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}
