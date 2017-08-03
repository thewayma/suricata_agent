package cron

import (
    //"fmt"
    "time"
    //"log"
    log "github.com/sirupsen/logrus"
    "github.com/thewayma/suricata_agent/g"
    "github.com/thewayma/suricata_agent/funcs"
)

func PreCollect() {
    for {
        funcs.UpdateCpuStat()
        funcs.UpdateDiskStats()
        time.Sleep(time.Second) //!< TODO: hardcode
    }
}

func Collect() {
	if !g.Config().Transfer.Enabled {
		return
	}

	if len(g.Config().Transfer.Addrs) == 0 {
		return
	}

	for _, v := range funcs.CollectorFuncs {
		go collect(int64(v.Interval), v.Fs)
	}
}

func collect(sec int64, fns []func() []*g.MetricValue) {
	t := time.NewTicker(time.Second * time.Duration(sec)).C
	for {
		<-t
        /*
		hostname, err := g.Hostname()
		if err != nil {
			continue
		}
        */
        ip := g.IP()

		mvs := []*g.MetricValue{}

		for _, fn := range fns {
			items := fn()
			if items == nil {
				continue
			}

			if len(items) == 0 {
				continue
			}

			for _, mv := range items {
                mvs = append(mvs, mv)
			}
		}

        dt := g.Config().DefaultTags
		now := time.Now().Unix()
		for j := 0; j < len(mvs); j++ { //!< Metric, Endpoint等在GaugeValue构造填充
			mvs[j].Step = sec
			//mvs[j].Endpoint = fmt.Sprintf("%s_%s", hostname, ip)
			mvs[j].Endpoint  = ip
			mvs[j].Timestamp = now

            if len(dt) > 0 {
                for k, v := range dt {
                    mvs[j].Tags[k] = v
                }
            }

        }


        //g.SendToTransfer(mvs)
        log.Debug("\nCron:", mvs)
	}
}
