package main
import (
    admin_member "github.com/LongMarch7/higo-admin/controller/admin/member"
    admin_portal "github.com/LongMarch7/higo-admin/controller/admin/portal"
    admin_setting "github.com/LongMarch7/higo-admin/controller/admin/setting"
    "github.com/LongMarch7/higo/config"
    "github.com/LongMarch7/higo/db"
    "github.com/LongMarch7/higo/app"
    "github.com/LongMarch7/higo/middleware"
    "github.com/LongMarch7/higo/service/web"
    admin "github.com/LongMarch7/higo-admin/controller/admin"
    admin_user "github.com/LongMarch7/higo-admin/controller/admin/user"
    "github.com/LongMarch7/higo-admin/controller/portal"
    "github.com/LongMarch7/higo/util/global"
    "github.com/LongMarch7/higo/view"
    "google.golang.org/grpc/grpclog"
    "github.com/LongMarch7/higo/auth"
)

type svrConfig struct{
    sOpts        []app.SOption
    mOpts        []middleware.MOption
    templatePath string
    sql          config.Sql
    mem          config.Memcache
}

func InitAdminServer(){
    admin.Init()
    admin_user.Init()
    admin_portal.Init()
    admin_member.Init()
    admin_setting.Init()
}

func InitPortalServer(){
    portal.Init()
}
func Server(config *config.Configer) {
    grpclog.SetLoggerV2(LogConfig(config.Config.Logger, config.Name + config.Port + ".log").NewLogger())

    svrConf :=svrResolving(config)
    if svrConf == nil{
        grpclog.Error("SvrResolving error")
        return
    }

    switch global.SeverName {
    case "AdminServer":
        InitAdminServer()
    case "PortalServer":
        InitPortalServer()
    default:
        grpclog.Error("Not found server by name")
        return
    }
    db.NewDb(db.DefaultNAME, db.Dialect(svrConf.sql.Driver),
        db.Args(svrConf.sql.User + ":" + svrConf.sql.Pwd + "@" + svrConf.sql.Net + "(" + svrConf.sql.Addr + ":" + svrConf.sql.Port + ")/"+svrConf.sql.Db),
        db.MaxOpenConns(svrConf.sql.MaxOpenConn),
        db.MaxIdleConns(svrConf.sql.MaxIdleConn),
        db.Show(svrConf.sql.Show),
        db.Level(svrConf.sql.Level))
    db.MemcacheInit(svrConf.mem.MaxIdleConn, svrConf.mem.Server)
    auth.NewCasbin()
    view.NewView(view.Dir(svrConf.templatePath), view.DelimsLeft("{["), view.DelimsRight("]}"))
    server := app.NewServer(svrConf.sOpts...)
    webServer := &web.GrpcServer{}
    webService := web.NewService()
    manager := middleware.NewMiddleware()
    HtmlOpts := append(svrConf.mOpts, middleware.MethodName("HTML"),middleware.Endpoint(web.MakeHtmlCallServerEndpoint(webService)))
    webServer.HtmlCallHandler = manager.AddMiddleware(HtmlOpts...).NewServer()
    ApiOpts := append(svrConf.mOpts, middleware.MethodName("API"), middleware.Endpoint(web.MakeApiCallServerEndpoint(webService)))
    webServer.ApiCallHandler = manager.AddMiddleware(ApiOpts...).NewServer()
    server.RegisterServiceServer(web.MakeRegisteFunc(webServer))
    server.Run()
}

