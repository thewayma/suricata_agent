package funcs

import (
    "fmt"
    "os"
    "net"
)


func GetUptime(unixSockFile string) {
    con, err := net.Dial("unix", unixSockFile)
    defer con.Close()

    if err != nil {
        fmt.Printf("Unix File %s not found\n", unixSockFile)
        os.Exit(-1)
    }
    fmt.Println("Connection OK.")

}
