package g

import (
    "fmt"
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
    }

    for _, tag := range tags {
        str := strings.Split(tag, "=")
        fmt.Println("tagK= ", str[0], "tagV= ", str[1])
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
