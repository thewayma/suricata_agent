package main

import (
    "os"
    "fmt"
    "flag"

    "github.com/thewayma/suricata_agent/g"
    "github.com/thewayma/suricata_agent/funcs"
    "github.com/thewayma/suricata_agent/cron"
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

    funcs.BuildMappers()

    go cron.InitDataHistory()
    cron.Collect()

    /*
    //!< commandTest
    funcs.GetVersion()
    funcs.GetRunningMode()
    funcs.GetCaptureMode()
    funcs.GetUptime()
    funcs.GetAllPortStats()
    funcs.ReloadRules()
    //funcs.ShutDown()
    */
    select {}
}
