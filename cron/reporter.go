package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/thewayma/suricata_agent/g"
	"github.com/thewayma/suricata_agent/funcs"
)

type AgentReportRequest struct {
    Hostname      string
    IP            string
    AgentVersion  string	//!< ids engine version
	Uptime		  int64		//!< ids engine uptime
}

func (this *AgentReportRequest) String() string {
    return fmt.Sprintf(
        "<Hostname:%s, IP:%s, engineVersion:%s, engineUptime:%s>",
        this.Hostname,
        this.IP,
        this.AgentVersion,
        this.Uptime,
    )
}

func ReportAgentStatus() {
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		go reportAgentStatus(time.Duration(g.Config().Heartbeat.Interval) * time.Second)
	}
}

func reportAgentStatus(interval time.Duration) {
	for {
		hostname, err := g.Hostname()
		if err != nil {
			hostname = fmt.Sprintf("error:%s", err.Error())
		}

		req := AgentReportRequest{
			Hostname:      hostname,
			IP:            g.IP(),
			AgentVersion:  g.VERSION,
			Uptime:		   funcs.GetUptime(),
		}

		var resp g.SimpleRpcResponse
		err = g.HbsClient.Call("Agent.ReportStatus", req, &resp)
		if err != nil || resp.Code != 0 {
			log.Println("call Agent.ReportStatus fail:", err, "Request:", req, "Response:", resp)
		}

		time.Sleep(interval)
	}
}
