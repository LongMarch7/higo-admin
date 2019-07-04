package main

import (
    "github.com/LongMarch7/higo/config"
    "github.com/LongMarch7/higo/app"
    "github.com/LongMarch7/higo/middleware"
    "github.com/LongMarch7/higo/middleware/hystrix"
    "github.com/LongMarch7/higo/middleware/ratelimit"
    "github.com/LongMarch7/higo/middleware/zipkin"
    "github.com/LongMarch7/higo/util/global"
    "strings"
    "time"
    "github.com/LongMarch7/higo/router"
    "context"
    "github.com/LongMarch7/higo/util/log"
)

func middlewareResolving(mw config.Middleware) ([]middleware.MOption, []zipkin.ZOption){
    var mOpts []middleware.MOption

    var hOptions []hystrix.HOption
    if mw.HystrixEPT != 0 {
        hOptions = append(hOptions,hystrix.ErrorPercentThreshold(mw.HystrixEPT))
    }
    if mw.HystrixMCR != 0 {
        hOptions = append(hOptions,hystrix.MaxConcurrentRequests(mw.HystrixMCR))
    }
    if mw.HystrixRVT != 0 {
        hOptions = append(hOptions,hystrix.RequestVolumeThreshold(mw.HystrixRVT))
    }
    if mw.HystrixSW != 0 {
        hOptions = append(hOptions,hystrix.SleepWindow(mw.HystrixSW))
    }
    if mw.HystrixTimeout != 0 {
        hOptions = append(hOptions,hystrix.Timeout(mw.HystrixTimeout))
    }
    if len(hOptions) > 0 {
        mOpts = append(mOpts,middleware.HOptions(hOptions))
    }

    var rOptions   []ratelimit.ROption
    if mw.RatelimitBurst > 0 {
        rOptions = append(rOptions,ratelimit.Burst(mw.RatelimitBurst))
    }
    if mw.RatelimitInterval > 0 {
        rOptions = append(rOptions,ratelimit.Interval(time.Duration(mw.RatelimitInterval)))
    }
    if len(rOptions) > 0 {
        mOpts = append(mOpts,middleware.ROptions(rOptions))
    }

    var zOptions   []zipkin.ZOption
    zOptions = append(zOptions,zipkin.Debug(mw.ZipkinDebug))
    if len(mw.ZipkinhostPort) > 0 {
        zOptions = append(zOptions,zipkin.HostPort(mw.ZipkinhostPort))
    }
    if len(mw.ZipkinUrl) > 0 {
        zOptions = append(zOptions,zipkin.Url(mw.ZipkinUrl))
    }
    if mw.ZipkinMaxLogs > 5 {
        zOptions = append(zOptions,zipkin.MaxLogsPerSpan(mw.ZipkinMaxLogs))
    }
    if len(zOptions) > 0 {
        mOpts = append(mOpts,middleware.ZOptions(zOptions))
    }
    return mOpts, zOptions
}
func cliResolving(config *config.Configer) *cliConfig{
    cliConf := new(cliConfig)
    cliConf.router = router.NewRouter()

    cliConf.mOpts = append(cliConf.mOpts,middleware.Prefix(config.Name),middleware.MethodName("request"))
    resolves,zipkinOpts := middlewareResolving(config.Config.CliMw)
    if len(resolves) > 0 {
        cliConf.mOpts = append(cliConf.mOpts,resolves...)
    }
    zipkinOpts = append(zipkinOpts,zipkin.Name(config.Name))
    cliConf.cOpts = append(cliConf.cOpts,app.CzOptions(zipkinOpts))

    if config.Config.RetryTime > 0 {
        cliConf.cOpts = append(cliConf.cOpts,app.CRetryTime(time.Duration(config.Config.RetryTime)))
    }
    if config.Config.RetryCount > 0 {
        cliConf.cOpts = append(cliConf.cOpts,app.CRetryCount(config.Config.RetryCount))
    }
    if len(config.Config.ConsulServer) > 0 {
        cliConf.cOpts = append(cliConf.cOpts,app.CConsulAddr(config.Config.ConsulServer))
    }
    if len(config.Config.StaticPath) > 0 {
        global.StaticPath = config.Config.RootPath + config.Config.StaticPath
    }else{
        global.StaticPath = config.Config.RootPath + "public/static"
    }
    if len(config.Config.UploadPath) > 0 {
        global.UploadPath = config.Config.RootPath + config.Config.UploadPath
    }else{
        global.UploadPath = config.Config.RootPath + "public/upload"
    }
    if len(config.Config.Port) > 0 {
        cliConf.port = config.Config.Port
    }else{
        cliConf.port = "8080"
    }
    cliConf.serviceList = config.Config.ServiceList
    return cliConf
}

func svrResolving(config *config.Configer) *svrConfig{
    svrConf :=new(svrConfig)

    svrConf.mOpts = append(svrConf.mOpts,middleware.Prefix(config.Name))
    resolves,zipkinOpts := middlewareResolving(config.Config.SvrMw)
    if len(resolves) > 0 {
        svrConf.mOpts = append(svrConf.mOpts,resolves...)
    }
    zipkinOpts = append(zipkinOpts,zipkin.Name(config.Name))
    svrConf.sOpts = append(svrConf.sOpts,app.SzOptions(zipkinOpts))

    if len(config.Name) > 0 {
        svrConf.sOpts = append(svrConf.sOpts,app.SPrefix(config.Name))
    }
    if len(config.Config.ConsulServer) > 0 {
        svrConf.sOpts = append(svrConf.sOpts,app.SConsulAddr(config.Config.ConsulServer))
    }

    svrConf.sOpts = append(svrConf.sOpts,app.SCtx(context.Background()))

    if len(config.Config.TemplatePath) > 0 {
        svrConf.templatePath = config.Config.RootPath + config.Config.TemplatePath
    }else{
        svrConf.templatePath = config.Config.RootPath + "template"
    }
    notFound := true
    for _,value := range config.Config.ServiceList{
        if value.Name == config.Name{
            if len(value.Addr) > 0 {
                svrConf.sOpts = append(svrConf.sOpts,app.SServerAddr(value.Addr))
            }
            if len(value.AdAddr) > 0 {
                svrConf.sOpts = append(svrConf.sOpts,app.SAdvertiseAddress(value.AdAddr))
            }
            if len(value.Count) > 0 {
                svrConf.sOpts = append(svrConf.sOpts,app.SMaxThreadCount(value.Count))
            }else{
                svrConf.sOpts = append(svrConf.sOpts,app.SMaxThreadCount("1024"))
            }

            if len(value.TemplatePath) >0 {
                svrConf.templatePath += "/" + value.TemplatePath
            }else{
                svrConf.templatePath += "/" + value.Name
            }
            svrConf.sOpts = append(svrConf.sOpts,app.SServerPort(config.Port))
            svrConf.sOpts = append(svrConf.sOpts,app.SAdvertisePort(config.AdPort))
            notFound = false
            break
        }
    }
    svrConf.sql = sqlResolving(config)
    svrConf.mem = memcacheResolving(config)
    if notFound {
        return nil
    }
    return svrConf
}

func sqlResolving(config *config.Configer) config.Sql{
    sql := config.Config.Sql
    if len(sql.Driver) == 0{
        sql.Driver = "mysql"
    }
    if len(sql.Port) == 0{
        sql.Port = "13306"
    }
    if len(sql.Pwd) == 0 {
        sql.Pwd = "123456"
    }
    if len(sql.Addr) == 0 {
        sql.Addr = "127.0.0.1"
    }
    if len(sql.Db) == 0 {
        sql.Db = "higo"
    }
    if len(sql.File) == 0 {
        sql.File = "db/default.sql"
    }else{
        sql.File = config.Config.RootPath + sql.File
    }
    if len(sql.Net) ==0 {
        sql.Net = "tcp"
    }
    if len(sql.User) == 0 {
        sql.User = "root"
    }
    if sql.MaxIdleConn <= 0 {
        sql.MaxIdleConn = 100
    }
    if sql.MaxOpenConn <= 0{
        sql.MaxOpenConn = 100
    }
    return sql
}
func memcacheResolving(config *config.Configer) config.Memcache{
    mem :=config.Config.Memcache
    if mem.MaxIdleConn <= 0 {
        mem.MaxIdleConn = 10
    }
    if len(mem.Server) == 0 {
        mem.Server = append(mem.Server,"127.0.0.1:11211")
    }
    return mem
}

func LogConfig(logger  zap.Config, name string) *zap.Config{
    if logger.MaxAge <= 0 {
        logger.MaxAge = 1
    }
    if logger.MaxSize <= 0 {
        logger.MaxSize =1
    }
    if logger.MaxBackups <= 0 {
        logger.MaxBackups = 3
    }
    if strings.Compare(logger.Type, zap.TYPE_FILE) == 0{
        logger.Filename = name
    }else{
        logger.Filename = ""
    }
    return &logger
}