package g

import (
    //"fmt"
    "strings"
    "strconv"
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

type MetricData struct {      //!< 统一agent,transporter data, 减小内存拷贝                 
    Endpoint    string              `json:"endpoint"`
    Metric      string              `json:"metric"`
    Value       float64             `json:"value"`
    Step        int64               `json:"step"`
    Type        string              `json:"Type"`
    Tags        map[string]string   `json:"tags"`
    Timestamp   int64               `json:"timestamp"`
}

func NewMetric(metric string, v interface{}, dataType string, tags ...string) *MetricData {
    //!< 在agent端判断数据类型, 避免transporter的内存拷贝
    var vv float64
    switch cv := v.(type) {
    case string:
        vv, _ = strconv.ParseFloat(cv, 64)
    case float64:
        vv = cv
    case int64:
        vv = float64(cv)
    }

    mv := MetricData {
        Metric: metric,
        Value:  vv,
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

func GaugeValue(metric string, val interface{}, tags ...string) *MetricData {
    return NewMetric(metric, val, "GAUGE", tags...)		//!< 瞬时型监控值
}

func CounterValue(metric string, val interface{}, tags ...string) *MetricData {
    return NewMetric(metric, val, "COUNTER", tags...)	//!< 累加型监控值
}

