package db

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"net"
	"time"
)

type WithDbMetrics struct {
	db             DBTX
	dbname         string
	errorCounter   *prometheus.CounterVec
	queryDuration  prometheus.Histogram
}

func NewDbMetrics(db DBTX, dbname string) *WithDbMetrics {
	subsystem := "go_sql"

	return &WithDbMetrics{
		db: db,
		errorCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem:   subsystem,
				Name:        "errors",
				Help:        "Counts sqlc errors",
				ConstLabels: prometheus.Labels{"db_name": dbname},
			},
			[]string{"class", "name"}),
		queryDuration: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      "query_duration_seconds",
				Help:      "Histogram of touches",
				ConstLabels: prometheus.Labels{"db_name": dbname},
				//Buckets: prometheus.ExponentialBuckets(1, 10, 6),
			},
		),
	}
}

func (m *WithDbMetrics) Collectors() []prometheus.Collector {
	return []prometheus.Collector {
		m.errorCounter,
		m.queryDuration,
	}
}

func (m *WithDbMetrics) ExecContext(context context.Context, sql string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	result, err := m.db.ExecContext(context, sql, args...)
	elapsed := time.Since(start).Seconds()
	m.queryDuration.Observe(elapsed)

	if err != nil {
		if errPq, ok := err.(*pq.Error); ok {
			m.errorCounter.WithLabelValues(errPq.Code.Class().Name(), errPq.Code.Name()).Inc()
		} else if _, ok := err.(net.Error); ok {
			m.errorCounter.WithLabelValues("tcp", "connection").Inc()
		} else {
			m.errorCounter.WithLabelValues("n/a", "n/a").Inc()
		}
	}
	return result, err
}

func (m *WithDbMetrics) QueryContext(context context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	result, err := m.db.QueryContext(context, sql, args...)
	elapsed := time.Since(start).Seconds()
	m.queryDuration.Observe(elapsed)

	if err != nil {
		if errPq, ok := err.(*pq.Error); ok {
			m.errorCounter.WithLabelValues(errPq.Code.Class().Name(), errPq.Code.Name()).Inc()
		} else if _, ok := err.(net.Error); ok {
			m.errorCounter.WithLabelValues("tcp", "connection").Inc()
		} else {
			m.errorCounter.WithLabelValues("n/a", "n/a").Inc()
		}
	}
	return result, err
}

func (m *WithDbMetrics) PrepareContext(context context.Context, sql string) (*sql.Stmt, error) {
	return m.db.PrepareContext(context, sql)
}

func (m *WithDbMetrics) QueryRowContext(context context.Context, sql string, args ...interface{}) *sql.Row {
	start := time.Now()
	result := m.db.QueryRowContext(context, sql, args...)
	elapsed := time.Since(start).Seconds()
	m.queryDuration.Observe(elapsed)
	return result
}
