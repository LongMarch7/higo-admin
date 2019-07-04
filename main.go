package main

import (
    "flag"
    "fmt"
    "github.com/LongMarch7/higo/config"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/util/global"
    "github.com/LongMarch7/higo/util/log"
    "google.golang.org/grpc/grpclog"
    "os"
    "os/signal"
    "syscall"
)

func init() {
    s := make(chan os.Signal, 1)
    signal.Notify(s, syscall.Signal(0xc))
    go func() {
        for {
            <-s
            conf :=&config.Config{}
            err := config.Load(config.ConfigFilePath, conf)
            if err != nil {
                fmt.Println("reload config failed")
                continue
            }
            zap.SetLogLevel(conf.Logger.Level)
            grpclog.Info(" ReLoad config;level: ", conf.Logger.Level)
        }
    }()
}
func main(){
    Mode       := flag.String("mode","","set mode")
    svrName    := flag.String("name","","service name")
    configPath := flag.String("conf","./config.json","config path")
    port       := flag.String("port","","server port")
    adPort     := flag.String("ad_port","","ad server port")
    flag.Parse()

    conf := config.ConfigInit(*Mode,*svrName,*configPath, *port, *adPort)
    if conf == nil{
        fmt.Println("star error by init config")
        return
    }
    if global.AppMode == define.CliMode {
        GateWay(conf)
    }else if global.AppMode == define.SvrMode{
        Server(conf)
    }else{
        Initialization(conf)
    }
}