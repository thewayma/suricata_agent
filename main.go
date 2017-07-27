package main

import (
    "os"
    "fmt"
    "flag"

    "github.com/thewayma/suricata_agent/g"
    "github.com/thewayma/suricata_agent/cron"
    "github.com/thewayma/suricata_agent/http"
    "github.com/thewayma/suricata_agent/funcs"
)

func main() {
    cfg := flag.String("c", "cfg.json", "configuration file")
    version := flag.Bool("v", false, "show version")

    flag.Parse()

    if *version {
        fmt.Println(g.VERSION)
        os.Exit(0)
    }

    g.ParseConfig(*cfg)

    g.InitLog()
    g.InitRpcClients()

    funcs.GenerateCollectorFuncs()

    go cron.PreCollect()

    cron.Collect()

    go http.Start()

    select {}
}
