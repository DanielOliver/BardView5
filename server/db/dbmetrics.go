package db

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"net"
)

type WithDbMetrics struct {
	db           DBTX
	dbname       string
	errorCounter *prometheus.CounterVec
	successCounter *prometheus.CounterVec
}

func NewDbMetrics(db DBTX, dbname string) *WithDbMetrics {
	return &WithDbMetrics{
		db: db,
		errorCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem:   "go_pq",
				Name:        "errors",
				Help:        "Counts sqlc errors",
				ConstLabels: prometheus.Labels{"db_name": dbname},
			},
			[]string{"class", "name"}),
		successCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem:   "go_pq",
				Name:        "touch",
				Help:        "Counts sqlc queries and exec",
				ConstLabels: prometheus.Labels{"db_name": dbname},
			},
			[]string{"type"}),
	}
}

func (m *WithDbMetrics) Collectors() []prometheus.Collector {
	return []prometheus.Collector {
		m.errorCounter,
		m.successCounter,
	}
}

func (m *WithDbMetrics) ExecContext(context context.Context, sql string, args ...interface{}) (sql.Result, error) {
	result, err := m.db.ExecContext(context, sql, args...)
	if err != nil {
		if errPq, ok := err.(*pq.Error); ok {
			m.errorCounter.WithLabelValues(errPq.Code.Class().Name(), errPq.Code.Name()).Inc()
		} else if _, ok := err.(net.Error); ok {
			m.errorCounter.WithLabelValues("tcp", "connection").Inc()
		} else {
			m.errorCounter.WithLabelValues("n/a", "n/a").Inc()
		}
	} else {
		m.successCounter.WithLabelValues("exec").Inc()
	}
	return result, err
}

func (m *WithDbMetrics) QueryContext(context context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	result, err := m.db.QueryContext(context, sql, args...)
	if err != nil {
		if errPq, ok := err.(*pq.Error); ok {
			m.errorCounter.WithLabelValues(errPq.Code.Class().Name(), errPq.Code.Name()).Inc()
		} else if _, ok := err.(net.Error); ok {
			m.errorCounter.WithLabelValues("tcp", "connection").Inc()
		} else {
			m.errorCounter.WithLabelValues("n/a", "n/a").Inc()
		}
	}else {
		m.successCounter.WithLabelValues("query").Inc()
	}
	return result, err
}

func (m *WithDbMetrics) PrepareContext(context context.Context, sql string) (*sql.Stmt, error) {
	return m.db.PrepareContext(context, sql)
}

func (m *WithDbMetrics) QueryRowContext(context context.Context, sql string, args ...interface{}) *sql.Row {
	return m.db.QueryRowContext(context, sql, args...)
}
