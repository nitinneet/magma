// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/prometheus/common/model"

	time "time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// PrometheusAPI is an autogenerated mock type for the PrometheusAPI type
type PrometheusAPI struct {
	mock.Mock
}

// Alerts provides a mock function with given fields: ctx
func (_m *PrometheusAPI) Alerts(ctx context.Context) (v1.AlertsResult, error) {
	ret := _m.Called(ctx)

	var r0 v1.AlertsResult
	if rf, ok := ret.Get(0).(func(context.Context) v1.AlertsResult); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(v1.AlertsResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LabelNames provides a mock function with given fields: ctx
func (_m *PrometheusAPI) LabelNames(ctx context.Context) ([]string, v1.Warnings, error) {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context) v1.Warnings); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LabelValues provides a mock function with given fields: ctx, label
func (_m *PrometheusAPI) LabelValues(ctx context.Context, label string) (model.LabelValues, v1.Warnings, error) {
	ret := _m.Called(ctx, label)

	var r0 model.LabelValues
	if rf, ok := ret.Get(0).(func(context.Context, string) model.LabelValues); ok {
		r0 = rf(ctx, label)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.LabelValues)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context, string) v1.Warnings); ok {
		r1 = rf(ctx, label)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, label)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Metadata provides a mock function with given fields: ctx, metric, limit
func (_m *PrometheusAPI) Metadata(ctx context.Context, metric string, limit string) (map[string][]v1.Metadata, error) {
	ret := _m.Called(ctx, metric, limit)

	var r0 map[string][]v1.Metadata
	if rf, ok := ret.Get(0).(func(context.Context, string, string) map[string][]v1.Metadata); ok {
		r0 = rf(ctx, metric, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]v1.Metadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, metric, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: ctx, query, ts
func (_m *PrometheusAPI) Query(ctx context.Context, query string, ts time.Time) (model.Value, v1.Warnings, error) {
	ret := _m.Called(ctx, query, ts)

	var r0 model.Value
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) model.Value); ok {
		r0 = rf(ctx, query, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Value)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time) v1.Warnings); ok {
		r1 = rf(ctx, query, ts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, time.Time) error); ok {
		r2 = rf(ctx, query, ts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// QueryRange provides a mock function with given fields: ctx, query, r
func (_m *PrometheusAPI) QueryRange(ctx context.Context, query string, r v1.Range) (model.Value, v1.Warnings, error) {
	ret := _m.Called(ctx, query, r)

	var r0 model.Value
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.Range) model.Value); ok {
		r0 = rf(ctx, query, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Value)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.Range) v1.Warnings); ok {
		r1 = rf(ctx, query, r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, v1.Range) error); ok {
		r2 = rf(ctx, query, r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Series provides a mock function with given fields: ctx, matches, startTime, endTime
func (_m *PrometheusAPI) Series(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]model.LabelSet, v1.Warnings, error) {
	ret := _m.Called(ctx, matches, startTime, endTime)

	var r0 []model.LabelSet
	if rf, ok := ret.Get(0).(func(context.Context, []string, time.Time, time.Time) []model.LabelSet); ok {
		r0 = rf(ctx, matches, startTime, endTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.LabelSet)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context, []string, time.Time, time.Time) v1.Warnings); ok {
		r1 = rf(ctx, matches, startTime, endTime)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, time.Time, time.Time) error); ok {
		r2 = rf(ctx, matches, startTime, endTime)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TargetsMetadata provides a mock function with given fields: ctx, matchTarget, metric, limit
func (_m *PrometheusAPI) TargetsMetadata(ctx context.Context, matchTarget string, metric string, limit string) ([]v1.MetricMetadata, error) {
	ret := _m.Called(ctx, matchTarget, metric, limit)

	var r0 []v1.MetricMetadata
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) []v1.MetricMetadata); ok {
		r0 = rf(ctx, matchTarget, metric, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1.MetricMetadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, matchTarget, metric, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}