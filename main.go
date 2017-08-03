package main

import (
    "flag"
    "github.com/thewayma/suricata_agent/g"
    "github.com/thewayma/suricata_agent/cron"
    "github.com/thewayma/suricata_agent/http"
    "github.com/thewayma/suricata_agent/funcs"
)

func main() {
    cfg := flag.String("c", "cfg.json", "configuration file")

    flag.Parse()
    g.ParseConfig(*cfg)

    g.InitLog()
    g.InitRpcClients()

    funcs.GenerateCollectorFuncs()

    go cron.PreCollect()

    cron.Collect()
    cron.ReportAgentStatus()

    go http.Start()

    select {}
}
