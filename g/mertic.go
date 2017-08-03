package g

import (
    //"fmt"
    "strings"
)

type MetricValue struct {
    Endpoint  string      `json:"endpoint"`
    Metric    string      `json:"metric"`
    Value     interface{} `json:"value"`
    Step      int64       `json:"step"`
    Type      string      `json:"counterType"`
    Tags      map[string]string `json:"tags"`
    Timestamp int64       `json:"timestamp"`
}

func NewMetricValue(metric string, val interface{}, dataType string, tags ...string) *MetricValue {
    mv := MetricValue {
        Metric: metric,
        Value:  val,
        Type:   dataType,
        Tags:   make(map[string]string),
    }

    for _, tag := range tags {
        str := strings.Split(tag, "=")
        //fmt.Printf("\n\n\n k=%s, v=%s\n\n\n", str[0], str[1])
        mv.Tags[str[0]] = str[1]
    }

    return &mv
}

func GaugeValue(metric string, val interface{}, tags ...string) *MetricValue {
    return NewMetricValue(metric, val, "GAUGE", tags...)
}

func CounterValue(metric string, val interface{}, tags ...string) *MetricValue {
    return NewMetricValue(metric, val, "COUNTER", tags...)
}

