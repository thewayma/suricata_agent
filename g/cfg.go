package g

import (
	"os"
    "net"
    "sync"
    "time"
    "strings"
    "encoding/json"

    "github.com/toolkits/file"
)

var (
	LocalIp string
    ConfigFile string
    config *GlobalConfig
    lock = new(sync.RWMutex)
)

type LogConfig struct {
    LogLevel string
    Output   string
}

type HeartbeatConfig struct {
    Enabled  bool
    Addr     string
    Interval int
    Timeout  int
}

type TransferConfig struct {
    Enabled  bool
    Addrs    []string
    Interval int        //!< 监控项采集周期
    Timeout  int
}

type HttpConfig struct {
    Enabled  bool
    Listen   string
}

type GlobalConfig struct {
	Hostname		string
	Ip				string
    UnixSockFile    string
    Log             *LogConfig
    Heartbeat		*HeartbeatConfig
	Transfer		*TransferConfig
    Http            *HttpConfig
    DefaultTags     map[string]string
}

func InitLocalIp() {
    if Config().Transfer.Enabled {
        conn, err := net.DialTimeout("tcp", Config().Transfer.Addrs[0], time.Second*10)
        if err != nil {
            Log.Error("get local addr failed !")
        } else {
            LocalIp = strings.Split(conn.LocalAddr().String(), ":")[0]
            conn.Close()
        }
    } else {
        Log.Error("hearbeat is not enabled, can't get localip")
    }
}

func Config() *GlobalConfig {
    lock.RLock()
    defer lock.RUnlock()
    return config
}

func Hostname() (string, error) {
    hostname := Config().Hostname
    if hostname != "" {
        return hostname, nil
    }

    hostname, err := os.Hostname()
    if err != nil {
        Log.Error("ERROR: os.Hostname() fail", err)
    }
    return hostname, err
}

func IP() string {
    ip := Config().Ip
    if ip != "" {
        // use ip in configuration
        return ip
    }

    if len(LocalIp) > 0 {
        ip = LocalIp
    }

    return ip
}

func ParseConfig(cfg string) {
    if cfg == "" {
        Log.Critical("use -c to specify configuration file")
    }

    if !file.IsExist(cfg) {
        Log.Critical("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
    }

    ConfigFile = cfg

    configContent, err := file.ToTrimString(cfg)
    if err != nil {
        Log.Critical("read config file:", cfg, "fail:", err)
    }

    var c GlobalConfig
    err = json.Unmarshal([]byte(configContent), &c)
    if err != nil {
        Log.Critical("parse config file:", cfg, "fail:", err)
    }

    lock.Lock()
    defer lock.Unlock()

    config = &c

    Log.Debug("read config file:", cfg, "successfully")
}
