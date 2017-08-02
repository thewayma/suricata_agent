package g

import (
    "fmt"
    "strings"
)

type MetricValue struct {
    Endpoint  string
    Metric    string
    Value     interface{}
    Step      int64
    Type      string
    Timestamp int64
    //Tags    string    //!< tag机制暂时不需要, 每个上传的metric都必须带上Endpoint, Metric !!!
}

func NewMetricValue(metric string, val interface{}, dataType string, tags ...string) *MetricValue {
    mv := MetricValue {
        Metric: metric,
        Value:  val,
        Type:   dataType,
    }

    /*暂时去掉tag切片
    size := len(tags)

    if size > 0 {
        mv.Tags = strings.Join(tags, ",")
    }
    */
    return &mv
}

func GaugeValue(metric string, val interface{}, tags ...string) *MetricValue {
    return NewMetricValue(metric, val, "GAUGE", tags...)
}

func CounterValue(metric string, val interface{}, tags ...string) *MetricValue {
    return NewMetricValue(metric, val, "COUNTER", tags...)
}
