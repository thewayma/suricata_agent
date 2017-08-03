package funcs

import (
    "github.com/thewayma/suricata_agent/g"
)

type FuncsAndInterval struct {
    Fs       []func() []*g.MetricData
    Interval int
}

var CollectorFuncs []FuncsAndInterval

func GenerateCollectorFuncs() {
    interval := g.Config().Transfer.Interval
    CollectorFuncs = []FuncsAndInterval {
        {
            Fs: []func() []*g.MetricData {
                //GetUptime,
               CpuMetrics,
               LoadAvgMetrics,
               MemMetrics,
               DiskIOMetrics,
               IOStatsMetrics,
            },
            Interval: interval,
        },
    }
}
