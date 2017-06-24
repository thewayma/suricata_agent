package g

import (
    "os"
    "log"
    "fmt"
    "runtime"
)

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
