package main

import (
    "os"
    "fmt"
    "flag"
    "github.com/thewayma/suricata_agent_go/g"
    "github.com/thewayma/suricata_agent_go/funcs"
    //"github.com/thewayma/suricata_agent_go/cron"
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

    funcs.GetUptime()

}
