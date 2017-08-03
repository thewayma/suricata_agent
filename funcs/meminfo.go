package funcs

import (
	"github.com/thewayma/suricata_agent/g"
	"github.com/toolkits/nux"
	"log"
)

func MemMetrics() []*g.MetricData {
	m, err := nux.MemInfo()
	if err != nil {
		log.Println(err)
		return nil
	}

	memFree := m.MemFree + m.Buffers + m.Cached
	memUsed := m.MemTotal - memFree

	pmemFree := 0.0
	pmemUsed := 0.0
	if m.MemTotal != 0 {
		pmemFree = float64(memFree) * 100.0 / float64(m.MemTotal)
		pmemUsed = float64(memUsed) * 100.0 / float64(m.MemTotal)
	}

	pswapFree := 0.0
	pswapUsed := 0.0
	if m.SwapTotal != 0 {
		pswapFree = float64(m.SwapFree) * 100.0 / float64(m.SwapTotal)
		pswapUsed = float64(m.SwapUsed) * 100.0 / float64(m.SwapTotal)
	}

	return []*g.MetricData{
		g.GaugeValue("mem.memtotal", m.MemTotal),
		g.GaugeValue("mem.memused", memUsed),
		g.GaugeValue("mem.memfree", memFree),
		g.GaugeValue("mem.swaptotal", m.SwapTotal),
		g.GaugeValue("mem.swapused", m.SwapUsed),
		g.GaugeValue("mem.swapfree", m.SwapFree),
		g.GaugeValue("mem.memfree.percent", pmemFree),
		g.GaugeValue("mem.memused.percent", pmemUsed),
		g.GaugeValue("mem.swapfree.percent", pswapFree),
		g.GaugeValue("mem.swapused.percent", pswapUsed),
	}

}
