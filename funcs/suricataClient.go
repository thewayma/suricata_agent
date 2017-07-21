package funcs

import (
    "fmt"
    "os"
    "net"
    "github.com/thewayma/suricata_agent/g"
    "github.com/antonholmquist/jason"
)

var (
    protocolMap map[string]string
    buf = make([]byte, 1024)

    maxPortNum int
    ifaceMap map[int]string     //!< portId <-> portName
)

func init() {
    protocolMap = make(map[string]string)
    protocolMap["version"] = `{"version": "0.1"}`
    protocolMap["command"] = `{"command": "%s"}`

    maxPortNum = 16
    ifaceMap = make(map[int]string)
}

func suriConnect() net.Conn {
    conn, err := net.Dial("unix", g.Config().UnixSockFile)
    //g.checkError(err)
    if err != nil {
        fmt.Printf("Unix File %s not found\n", g.Config().UnixSockFile)
        os.Exit(-1)
    }

    return conn
}

func suriMakeCommand(conn net.Conn, com string) string {
    return fmt.Sprintf(protocolMap["command"], com)
}

func suriSendVersion(conn net.Conn) {
    //fmt.Printf("SND: %s\n", protocolMap["version"])
    conn.Write([]byte(protocolMap["version"]))

    conn.Read(buf)
    //fmt.Printf("RCV: %s\n", buf)

    //!< TODO: OK, NOK
}

func suriSendCommandGetInt(conn net.Conn, data string) (int64, error) {
    conn.Write([]byte(data))
    //fmt.Printf("SND: %s\n", data)

    conn.Read(buf)
    //fmt.Printf("RCV: %s\n", buf)

    j, _ := jason.NewObjectFromBytes([]byte(buf))

    if res, _ := j.GetString("return"); res == "OK" {
        return j.GetInt64("message")
    } else {
        return -299, fmt.Errorf("%s Command Error", data)
    }

}

func suriSendCommandGetString(conn net.Conn, data string) (string, error) {
    conn.Write([]byte(data))
    //fmt.Printf("SND: %s\n", data)

    conn.Read(buf)
    //fmt.Printf("RCV: %s\n", buf)

    j, _ := jason.NewObjectFromBytes([]byte(buf))

    if res, _ := j.GetString("return"); res == "OK" {
        return j.GetString("message")
    } else {
        return "error", fmt.Errorf("%s Command Error", data)
    }
}

func suriSendCommandGetIface(conn net.Conn, data string) (string, error) {
    conn.Write([]byte(data))
    //fmt.Printf("SND: %s\n", data)

    conn.Read(buf)
    //fmt.Printf("RCV: %s\n", buf)

    j, _ := jason.NewObjectFromBytes([]byte(buf))

    if res, _ := j.GetString("return"); res == "OK" {
        messObj, _ := j.GetObject("message")
        ifaceObj, _ := messObj.GetStringArray("ifaces")

        for index, dataItem := range ifaceObj {
            ifaceMap[index] = dataItem

        }



        return "11", nil

    } else {
        return "error", fmt.Errorf("%s Command Error", data)
    }
}

//!< 周期性采集
func GetUptime() []*g.MetricValue {
    conn := suriConnect()
    defer conn.Close()

    suriSendVersion(conn)
    com := suriMakeCommand(conn, "uptime")
    ret, _ := suriSendCommandGetInt(conn, com)

    //fmt.Println("Uptime:", g.GaugeValue("suricata_uptime", ret))
    return []*g.MetricValue{g.GaugeValue("suricata_uptime", ret)}
}

//!< 以下为非周期性采集动作
func ShutDown() {
    conn := suriConnect()
    defer conn.Close()

    suriSendVersion(conn)
    com := suriMakeCommand(conn, "shutdown")
    ret, _ := suriSendCommandGetString(conn, com)

    fmt.Println(ret)
}

func ReloadRules() {
    conn := suriConnect()
    defer conn.Close()

    suriSendVersion(conn)
    com := suriMakeCommand(conn, "reload-rules")
    ret, _ := suriSendCommandGetString(conn, com)

    fmt.Println(ret)
}

func GetVersion() {
    conn := suriConnect()
    defer conn.Close()

    suriSendVersion(conn)
    com := suriMakeCommand(conn, "version")
    ret, _ := suriSendCommandGetString(conn, com)

    fmt.Println(ret)
}

func GetRunningMode() {
    conn := suriConnect()
    defer conn.Close()

    suriSendVersion(conn)
    com := suriMakeCommand(conn, "running-mode")
    ret, _ := suriSendCommandGetString(conn, com)

    fmt.Println(ret)
}

func GetCaptureMode() {
    conn := suriConnect()
    defer conn.Close()

    suriSendVersion(conn)
    com := suriMakeCommand(conn, "capture-mode")
    ret, _ := suriSendCommandGetString(conn, com)

    fmt.Println(ret)
}

func GetProfilingCouters() {
    conn := suriConnect()
    defer conn.Close()

    suriSendVersion(conn)
    com := suriMakeCommand(conn, "dump-counters")

    conn.Write([]byte(com))

    buf = make([]byte, 10240)
    conn.Read(buf)

    fmt.Printf("ProfilingCounters: %s", buf)
}

func GetAllPortStats() {
    conn := suriConnect()
    defer conn.Close()

    suriSendVersion(conn)
    com := suriMakeCommand(conn, "iface-list")

    suriSendCommandGetIface(conn, com)



}
