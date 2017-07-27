package g

import (
	"github.com/toolkits/net"
	"log"
	"math"
	"net/rpc"
	"sync"
	"time"
)

var (
    HbsClient *RpcClient
)

type RpcClient struct {
	sync.Mutex
	rpcClient *rpc.Client
	RpcServer string
	Timeout   time.Duration
}

func InitRpcClients() {
    if Config().Heartbeat.Enabled {
        HbsClient = &RpcClient{
            RpcServer: Config().Heartbeat.Addr,
            Timeout: time.Duration(Config().Heartbeat.Timeout) * time.Millisecond,
        }
    }
}

func (this *RpcClient) close() {
	if this.rpcClient != nil {
		this.rpcClient.Close()
		this.rpcClient = nil
	}
}

func (this *RpcClient) serverConn() error {
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

func (this *RpcClient) Call(method string, args interface{}, reply interface{}) error {
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
