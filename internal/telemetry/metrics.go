package telemetry

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const meterName = "url-shortener"

type Metrics interface {
	Incr(ctx context.Context, name string, tags []attribute.KeyValue)
}

type otelMetrics struct {
	meter    metric.Meter
	counters map[string]metric.Int64Counter
	mu       sync.RWMutex
}

func NewMetrics() Metrics {
	m := &otelMetrics{
		meter:    otel.Meter(meterName),
		counters: make(map[string]metric.Int64Counter),
	}

	// Pre-register known counters
	counter, _ := m.meter.Int64Counter(HTTPRequestsTotal)
	m.counters[HTTPRequestsTotal] = counter

	return m
}

func (m *otelMetrics) Incr(ctx context.Context, name string, tags []attribute.KeyValue) {
	m.mu.RLock()
	counter, exists := m.counters[name]
	m.mu.RUnlock()

	if !exists {
		var err error
		counter, err = m.meter.Int64Counter(name)
		if err != nil {
			return
		}
		m.mu.Lock()
		m.counters[name] = counter
		m.mu.Unlock()
	}

	counter.Add(ctx, 1, metric.WithAttributes(tags...))
}
