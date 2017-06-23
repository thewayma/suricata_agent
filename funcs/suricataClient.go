package funcs

import (
    "fmt"
    "os"
    "net"
)


func GetUptime(unixSockFile string) {
    fmt.Println(unixSockFile)

    con, err := net.Dial("unix", unixSockFile)
    if err != nil {
        fmt.Printf("Unix File %s not found\n", unixSockFile)
        os.Exit(-1)
    }
    defer con.Close()

    fmt.Println("Connection OK.")

}
