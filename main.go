package main

import (
    "fmt"
    "github.com/thewayma/suricata_agent_go/g"
    "github.com/thewayma/suricata_agent_go/funcs"
    //"github.com/thewayma/suricata_agent_go/cron"
)

func main() {

    fmt.Println(g.UnixSockFile)

    fun.GetUptime(g.UnixSockFile)

}
