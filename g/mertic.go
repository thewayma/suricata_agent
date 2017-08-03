package g

import (
    "fmt"
    "strings"
)

//!< 监控项结果
type MetricValue struct {
    Endpoint  string      `json:"endpoint"`
    Metric    string      `json:"metric"`
    Value     interface{} `json:"value"`
    Step      int64       `json:"step"`
    Type      string      `json:"counterType"`
    Tags      map[string]string `json:"tags"`
    Timestamp int64       `json:"timestamp"`
}

func (this *MetricValue) String() string {
    return fmt.Sprintf(
        "<Endpoint:%s, Metric:%s, Type:%s, Tags:%s, Step:%d, Time:%d, Value:%v>",
        this.Endpoint,
        this.Metric,
        this.Type,
        this.Tags,
        this.Step,
        this.Timestamp,
        this.Value,
    )
}

func NewMetricValue(metric string, val interface{}, dataType string, tags ...string) *MetricValue {
    mv := MetricValue {
        Metric: metric,
        Value:  val,
        Type:   dataType,
    }
/*
    size := len(tags)

    if size > 0 {
        mv.Tags = strings.Join(tags, ",")
    }
*/

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
