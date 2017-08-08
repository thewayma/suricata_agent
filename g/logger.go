package g

import (
    "time"
    "github.com/alecthomas/log4go"
)

var Log = make(log4go.Logger)

//!< 日志等级从低到高: FINEST, FINE, DEBUG, TRACE, INFO, WARNING, ERROR, CRITICAL
func InitLog() error {
    loglevel := log4go.INFO

    switch Config().Log.LogLevel {
    case "debug":
        loglevel = log4go.DEBUG
    case "trace":
        loglevel = log4go.TRACE
    case "info":
        loglevel = log4go.INFO
    case "warn":
        loglevel = log4go.WARNING
    case "error":
        loglevel = log4go.ERROR
    case "critical":
        loglevel = log4go.CRITICAL
    }


    if Config().Log.Output == "file" {
        file := log4go.NewFileLogWriter("run.log", true)
        file.SetRotateLines(3600)
        file.SetRotateDaily(true)
        Log.AddFilter("file", loglevel, file)
    } else if Config().Log.Output == "console" {
        Log.AddFilter("stdout", loglevel, log4go.NewConsoleLogWriter())
    }

    Log.Info("Log Framework Inited, Start time: %s", time.Now().Format("15:04:05 MST 2006/01/02"))

    return nil
}
