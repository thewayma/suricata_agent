package funcs

import (
    "github.com/thewayma/suricata_agent_go/g"
)

type FuncsAndInterval struct {
    Fs       []func() []*g.MetricValue
    Interval int
}

var Mappers []FuncsAndInterval

func BuildMappers() {
    interval := g.Config().Transfer.Interval
    Mappers = []FuncsAndInterval {
        {
            Fs: []func() []*g.MetricValue {
                GetUptime,
            },
            Interval: interval,
        },
    }
}
