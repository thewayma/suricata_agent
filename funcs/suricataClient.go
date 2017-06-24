package funcs

import (
    "fmt"
    "os"
    "net"
    //"encoding/json"
    "github.com/thewayma/suricata_agent_go/g"
)

/*
type Version struct {
    Version string `json:"version"`
}

type Command struct {
    Command string `json:"command"`
}
*/

var (
    protocolMap map[string]string
    buf = make([]byte, 1024)
)

func init() {
    protocolMap = make(map[string]string)
    protocolMap["version"] = `{"version": "0.1"}`
    protocolMap["command"] = `{"command": "%s"}`
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
    fmt.Printf("SND: %s\n", protocolMap["version"])
    conn.Write([]byte(protocolMap["version"]))

    conn.Read(buf)
    fmt.Printf("RCV: %s\n", buf)

    //!< TODO: OK, NOK
}

func suriSendCommand(conn net.Conn, data string) {
    conn.Write([]byte(data))
    fmt.Printf("SND: %s\n", data)

    conn.Read(buf)
    fmt.Printf("RCV: %s\n", buf)

    //!< TODO: OK,NOK; 提取结果
}


func GetUptime() {
    conn := suriConnect()
    defer conn.Close()

    suriSendVersion(conn)
    com := suriMakeCommand(conn, "uptime")
    suriSendCommand(conn, com)
}
