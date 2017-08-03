package g

import (
	"log"
	"sync"
	"time"
    "net/rpc"
	"math"
	"math/rand"
    "github.com/toolkits/net"
)

type TransferResponse struct {
    Message string
    Total   int
    Invalid int
    Latency int64
}

type SingleConnRpcClient struct {
    sync.Mutex
    rpcClient *rpc.Client
    RpcServer string
    Timeout   time.Duration
}

func (r *SingleConnRpcClient) close() {
    if r.rpcClient != nil {
        r.rpcClient.Close()
        r.rpcClient = nil
    }
}

func (this *SingleConnRpcClient) serverConn() error {
    if this.rpcClient != nil {
        return nil
    }

    var err error
    var retry int = 1

    for {
        if this.rpcClient != nil {
            return nil
        }

        this.rpcClient, err = net.JsonRpcClient("tcp", this.RpcServer, this.Timeout)
        if err != nil {
            log.Printf("dial %s fail: %v", this.RpcServer, err)
            if retry > 3 {
                return err
            }
            time.Sleep(time.Duration(math.Pow(2.0, float64(retry))) * time.Second)
            retry++
            continue
        }
        return err
    }
}

func (this *SingleConnRpcClient) Call(method string, args interface{}, reply interface{}) error {

    this.Lock()
    defer this.Unlock()

    err := this.serverConn()
    if err != nil {
        return err
    }

    timeout := time.Duration(10 * time.Second)
    done := make(chan error, 1)

    go func() {
        err := this.rpcClient.Call(method, args, reply)
        done <- err
    }()

    select {
    case <-time.After(timeout):
        log.Printf("[WARN] rpc call timeout %v => %v", this.rpcClient, this.RpcServer)
        this.close()
    case err := <-done:
        if err != nil {
            this.close()
            return err
        }
    }

    return nil
}

var (
	TransferClientsLock *sync.RWMutex                   = new(sync.RWMutex)
	TransferClients     map[string]*SingleConnRpcClient = map[string]*SingleConnRpcClient{}
)

func SendMetrics(metrics []*MetricData, resp *TransferResponse) {
	rand.Seed(time.Now().UnixNano())
	for _, i := range rand.Perm(len(Config().Transfer.Addrs)) {
		addr := Config().Transfer.Addrs[i]

		c := getTransferClient(addr)
		if c == nil {
			c = initTransferClient(addr)
		}

		if updateMetrics(c, metrics, resp) {
			break
		}
	}
}

func initTransferClient(addr string) *SingleConnRpcClient {
	var c *SingleConnRpcClient = &SingleConnRpcClient{
		RpcServer: addr,
		Timeout:   time.Duration(Config().Transfer.Timeout) * time.Millisecond,
	}
	TransferClientsLock.Lock()
	defer TransferClientsLock.Unlock()
	TransferClients[addr] = c

	return c
}

func updateMetrics(c *SingleConnRpcClient, metrics []*MetricData, resp *TransferResponse) bool {
	err := c.Call("Transfer.Update", metrics, resp)
	if err != nil {
		log.Println("call Transfer.Update fail:", c, err)
		return false
	}
	return true
}

func getTransferClient(addr string) *SingleConnRpcClient {
	TransferClientsLock.RLock()
	defer TransferClientsLock.RUnlock()

	if c, ok := TransferClients[addr]; ok {
		return c
	}
	return nil
}
