package g

import (
    "os"
    "fmt"
    log "github.com/sirupsen/logrus"
)

//!< Debug, Info, Warn, Error, Fatal 日志级别由低到高
func InitLog() (err error) {
    //!< tty, file
    if Config().Log.Output == "file" {
        logfile, err := os.OpenFile("run.log", os.O_RDWR|os.O_CREATE, 0)
        if err != nil {
            fmt.Printf("%s\n", err.Error())
            os.Exit(-1)
        }
        log.SetOutput(logfile)
    }

    //!< text, json
    if Config().Log.Type == "json" {
        log.SetFormatter(&log.JSONFormatter{})
    } else {
        log.SetFormatter(&log.TextFormatter{})
    }

    //!< debug, info, warn, error, fatal
    switch Config().Log.LogLevel {
    case "info":
        log.SetLevel(log.InfoLevel)
    case "debug":
        log.SetLevel(log.DebugLevel)
    case "warn":
        log.SetLevel(log.WarnLevel)
    default:
        log.Fatal("log conf only allow [info, debug, warn], please check your confguire")
    }

    return nil
}
