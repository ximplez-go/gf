// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gmetric

import (
	"github.com/ximplez-go/gf/encoding/gjson"
	"github.com/ximplez-go/gf/errors/gcode"
	"github.com/ximplez-go/gf/errors/gerror"
	"github.com/ximplez-go/gf/text/gregex"
)

// localMetric implements interface Metric.
type localMetric struct {
	MetricInfo
}

// newMetric creates and returns an object that implements interface Metric.
func (meter *localMeter) newMetric(
	metricType MetricType, metricName string, metricOption MetricOption,
) (Metric, error) {
	if metricName == "" {
		return nil, gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`error creating %s metric while given name is empty, option: %s`,
			metricType, gjson.MustEncodeString(metricOption),
		)
	}
	if !gregex.IsMatchString(MetricNamePattern, metricName) {
		return nil, gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid metric name "%s", should match regular expression pattern "%s"`,
			metricName, MetricNamePattern,
		)
	}
	return &localMetric{
		MetricInfo: meter.newMetricInfo(metricType, metricName, metricOption),
	}, nil
}

// Info returns the basic information of a Metric.
func (l *localMetric) Info() MetricInfo {
	return l.MetricInfo
}
