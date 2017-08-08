package g

import (
    "os"
    "fmt"
    "runtime"
)

// code == 0 => success
// code == 1 => bad request
type SimpleRpcResponse struct {
    Code int `json:"code"`
}

func (this *SimpleRpcResponse) String() string {
    return fmt.Sprintf("<Code: %d>", this.Code)
}

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
