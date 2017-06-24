package funcs

import (
    "fmt"
    "os"
    "net"
    "encoding/json"
    "github.com/thewayma/suricata_agent_go/g"
)

type Version struct {
    Version string `json:"version"`
}

type Command struct {
    Command string `json:"command"`
}

var (
    protocolMap map[string]string
)

func init() {
    protocolMap = make(map[string]string)
    protocolMap["version"] = `{"version": "0.1"}`
    protocolMap["command"] = `{"command": "%s"}`
}

func suriConnect(sock string) (net.Conn, error) {
    return nil, nil
}

func suriMakeCommand(com string) string {
    return fmt.Sprintf(protocolMap["command"], com)
}

func suriSendCommand(data []byte) {

}


func GetUptime() {
    unixSockFile := g.Config().UnixSockFile
    //fmt.Println(unixSockFile)

    conn, err := net.Dial("unix", unixSockFile)
    if err != nil {
        fmt.Printf("Unix File %s not found\n", unixSockFile)
        os.Exit(-1)
    }
    defer conn.Close()

    //fmt.Printf("Unix Socket %s Connection Ok\n", unixSockFile)




    buf := make([]byte,1024)

    ver := Version{Version: "0.1"}
    data, err := json.Marshal(ver)
    if err != nil {
        fmt.Printf("JSON marshaling failed: %s", err)
    }
    fmt.Printf("%s\n", data)

    conn.Write([]byte(data))

    conn.Read(buf)
    fmt.Printf("%s\n", buf)


    ver1 := Command{Command: "uptime"}
    data1, err := json.Marshal(ver1)
    if err != nil {
        fmt.Printf("JSON marshaling failed: %s", err)
    }
    fmt.Printf("%s\n", data1)

    conn.Write([]byte(data1))

    conn.Read(buf)
    fmt.Printf("%s\n", buf)



}
